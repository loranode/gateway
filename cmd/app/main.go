package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/merzzzl/proto-rest-api/swagger"

	"github.com/loranode/gateway/api/rest"
	"github.com/loranode/gateway/internal/config"
	"github.com/loranode/gateway/internal/controller"
	"github.com/loranode/gateway/internal/repositories/radio"
	"github.com/loranode/gateway/internal/repositories/sqlite"
	"github.com/loranode/gateway/internal/repositories/webhook"
	"github.com/loranode/gateway/internal/services/events"
	"github.com/loranode/gateway/internal/services/registry"
	"github.com/loranode/gateway/internal/services/transport"
	"github.com/loranode/gateway/internal/worker"
)

func main() {
	const (
		shutdownTimeout = 5 * time.Second
		httpReadTimeout = 10 * time.Second
	)

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	cfg, err := config.Load()
	if err != nil {
		slog.Error("load config", "err", err)
		os.Exit(1)
	}

	db, err := sqlite.New(cfg.DBPath)
	if err != nil {
		slog.Error("open database", "err", err)
		os.Exit(1)
	}

	if err := db.Migrate(); err != nil {
		slog.Error("migrate database", "err", err)
		os.Exit(1)
	}

	reg := registry.New(db)
	ev := events.New(db, webhook.New())
	tr := transport.New(radio.New(cfg.NodeAddr))
	wrk := worker.New(tr, reg, ev)
	ctl := controller.New(reg, ev, wrk.Send)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go wrk.Run(ctx)

	router := runtime.NewRouter()
	rest.RegisterMeshServiceHandler(router, ctl)
	rest.RegisterCallbackServiceHandler(router, ctl)

	mux := router.Mux()
	mux.Handle("/swagger/", swagger.Handler(rest.GetGatewaySwagger()))
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)

		_, _ = w.Write([]byte("ok"))
	})

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           mux,
		ReadHeaderTimeout: httpReadTimeout,
	}

	go func() {
		<-ctx.Done()

		sctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		_ = srv.Shutdown(sctx) //nolint:contextcheck // shutdown must outlive the cancelled parent ctx
	}()

	slog.Info("starting gateway", "node_addr", cfg.NodeAddr, "http_addr", cfg.HTTPAddr)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("rest server stopped", "err", err)
		os.Exit(1)
	}
}

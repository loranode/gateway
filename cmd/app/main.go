package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/merzzzl/proto-rest-api/swagger"
	"google.golang.org/grpc"

	"github.com/loranode/gateway/api"
	"github.com/loranode/gateway/internal/config"
	"github.com/loranode/gateway/internal/controller"
	"github.com/loranode/gateway/internal/repositories/meshtastic"
	"github.com/loranode/gateway/internal/repositories/sqlite"
	"github.com/loranode/gateway/internal/services/events"
	"github.com/loranode/gateway/internal/services/mesh"
	"github.com/loranode/gateway/internal/services/registry"
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
	ev := events.New()
	meshSvc := mesh.New(meshtastic.New(cfg.NodeAddr))
	wrk := worker.New(meshSvc, reg, ev)
	ctl := controller.New(reg, ev, meshSvc)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go wrk.Run(ctx)

	// REST (HTTP/1.1) via proto-rest-api.
	router := runtime.NewRouter()
	api.RegisterMeshServiceHandler(router, ctl)

	mux := router.Mux()
	mux.Handle("/swagger/", swagger.Handler(api.GetGatewaySwagger()))
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)

		_, _ = w.Write([]byte("ok"))
	})

	// gRPC (HTTP/2) for MeshService and the event stream.
	grpcSrv := grpc.NewServer()
	api.RegisterMeshServiceServer(grpcSrv, ctl)
	api.RegisterEventServiceServer(grpcSrv, ctl)

	// One port for both: route gRPC (HTTP/2) by content-type, everything else to
	// REST. Protocols enables unencrypted HTTP/2 (h2c) so gRPC works over cleartext.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
			grpcSrv.ServeHTTP(w, r)

			return
		}

		mux.ServeHTTP(w, r)
	})

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           handler,
		ReadHeaderTimeout: httpReadTimeout,
		Protocols:         new(http.Protocols),
	}
	srv.Protocols.SetHTTP1(true)
	srv.Protocols.SetUnencryptedHTTP2(true)

	go func() {
		<-ctx.Done()

		grpcSrv.GracefulStop()

		sctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		_ = srv.Shutdown(sctx) //nolint:contextcheck // shutdown must outlive the cancelled parent ctx
	}()

	slog.Info("starting gateway", "node_addr", cfg.NodeAddr, "addr", cfg.HTTPAddr)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("server stopped", "err", err)
		os.Exit(1)
	}
}

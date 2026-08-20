FROM golang:1.26 AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -o /app ./cmd/app

FROM alpine:3.21
RUN apk add --no-cache ca-certificates
COPY --from=build /app /app

EXPOSE 8080
ENTRYPOINT ["/app"]

# ========= Build Stage ============
FROM golang:1.25.5-alpine AS builder 

WORKDIR /app 

COPY go.mod go.sum ./

RUN go mod download 

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \ 
    go build -o product-service ./cmd/grpc

# ========= Runtime Stage ============
FROM gcr.io/distroless/base-debian12

WORKDIR /app 

COPY --from=builder /app/.env /app
COPY --from=builder /app/product-service /app/product-service

EXPOSE 50051

USER nonroot:nonroot 

ENTRYPOINT ["/app/product-service"]
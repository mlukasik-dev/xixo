# STEP 1: Build the gateway
FROM golang:1.15-alpine AS builder

# Set necessary environmet variables needed for the image
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# Copy local dependencies
COPY protobuf /build/protobuf
COPY api /build/api
# Copy go.mod and go.sum then download dependencies using go mod
COPY api/gateway/go.mod api/gateway/go.sum /build/api/gateway/
WORKDIR /build/api/gateway
RUN go mod download

# Build the application
RUN go build -o main .

# STEP 2: Build a tiny image
FROM scratch

COPY --from=builder /build/api/gateway/main /dist/api_gateway

# Command to run when starting the container
CMD ["/dist/api_gateway"]
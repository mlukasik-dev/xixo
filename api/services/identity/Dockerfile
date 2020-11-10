# STEP 1: Build the service
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
COPY api/services/identity/go.mod api/services/identity/go.sum /build/api/services/identity/
WORKDIR /build/api/services/identity
RUN go mod download

# Build the application
RUN go build -o main .

# STEP 2: Build a tiny image
FROM scratch

COPY --from=builder /build/api/services/identity/main /dist/identity_service

# Command to run when starting the container
CMD ["/dist/identity_service"]
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
COPY api/services/account/go.mod api/services/account/go.sum /build/api/services/account/
WORKDIR /build/api/services/account
RUN go mod download

# Build the application
RUN go build -o main .

# STEP 2: Build a tiny image
FROM scratch

COPY --from=builder /build/api/services/account/main /dist/account_service

# Command to run when starting the container
CMD ["/dist/account_service"]
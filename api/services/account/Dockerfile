# STEP 1: Build service
FROM golang:1.15-alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# Copy go source files and download dependency using go mod
COPY protobuf /build/protobuf
COPY api /build/api
# Move to directory with go.mod and go.sum
WORKDIR /build/api
RUN go mod download

# Build the application
RUN go build -o account_service ./services/account

# STEP 2: Build a tiny image
FROM scratch

COPY --from=builder /build/api/account_service /dist/account_service

# Command to run when starting the container
CMD ["/dist/account_service"]
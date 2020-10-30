FROM golang:1.15-alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy go source files and download dependency using go mod
COPY protobuf ./protobuf
COPY api ./api
# Move to directory with go.mod and go.sum
WORKDIR /build/api
RUN go mod download

# Build the application
RUN go build -o main ./gateway

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/api/main .

# Command to run when starting the container
CMD ["/dist/main"]
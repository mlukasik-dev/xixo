package grpcerror

import (
	"github.com/vektah/gqlparser/gqlerror"
	"google.golang.org/grpc/status"
)

// GetError .
func GetError(err error) error {
	statusCode, ok := status.FromError(err)
	if !ok {
		return err
	}
	return gqlerror.Errorf(statusCode.Message())
}

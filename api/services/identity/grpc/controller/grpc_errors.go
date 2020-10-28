package controller

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrNoClaims .
var ErrNoClaims = status.Errorf(codes.PermissionDenied, "cannot get claims")

// ErrAccountIDsDoNotMatch .
var ErrAccountIDsDoNotMatch = status.Errorf(codes.PermissionDenied, "claimed accoundID and requested accoundID do not match")

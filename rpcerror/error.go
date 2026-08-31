package rpcerror

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewInternal 创建服务内部错误
func NewInternal(msg string) error {
	return status.Error(codes.Internal, msg)
}

// NewInvalidArgument 创建参数错误
func NewInvalidArgument(msg string) error {
	return status.Error(codes.InvalidArgument, msg)
}

// NewNotFound 创建资源不存在错误
func NewNotFound(msg string) error {
	return status.Error(codes.NotFound, msg)
}

// NewAlreadyExists 创建资源已存在错误
func NewAlreadyExists(msg string) error {
	return status.Error(codes.AlreadyExists, msg)
}

// NewUnauthenticated 创建未认证错误
func NewUnauthenticated(msg string) error {
	return status.Error(codes.Unauthenticated, msg)
}

// NewPermissionDenied 创建无权限错误
func NewPermissionDenied(msg string) error {
	return status.Error(codes.PermissionDenied, msg)
}

// NewResourceExhausted 创建资源不足错误
func NewResourceExhausted(msg string) error {
	return status.Error(codes.ResourceExhausted, msg)
}

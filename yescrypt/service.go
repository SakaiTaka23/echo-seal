package yescrypt

import (
	"context"
	yescryptv1 "echo-seal/gen/yescrypt/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	yescryptv1.UnimplementedYescryptServiceServer
	core Yescrypt
}

func NewService(yescrypt Yescrypt) *Service {
	return &Service{
		core: yescrypt,
	}
}

func (s *Service) Hash(_ context.Context, req *yescryptv1.HashRequest) (*yescryptv1.HashResponse, error) {
	password := req.GetPassword()

	hash, err := s.core.Hash(password)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &yescryptv1.HashResponse{
		Hash: hash,
	}, nil
}

func (s *Service) Verify(_ context.Context, req *yescryptv1.VerifyRequest) (*yescryptv1.VerifyResponse, error) {
	hash := req.GetHash()
	password := req.GetPassword()

	result, err := s.core.Verify(password, hash)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &yescryptv1.VerifyResponse{
		Result: result,
	}, nil
}

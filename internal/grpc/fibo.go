package grpc

import (
	"context"

	"test/internal/usecase"
	fibo "test/proto"
)

type FiboService struct {
	fibo.UnimplementedFiboServiceServer
	usecase usecase.IUsecase
}

func NewFiboService(usecase usecase.IUsecase) *FiboService {
	return &FiboService{
		usecase: usecase,
	}
}

func (f *FiboService) GetFibo(ctx context.Context, in *fibo.FiboRequest) (*fibo.FigoResponse, error) {
	response, err := f.usecase.Fib(ctx, in.X, in.Y)
	if err != nil {
		return nil, err
	}

	return &fibo.FigoResponse{
		Response: response,
	}, nil
}

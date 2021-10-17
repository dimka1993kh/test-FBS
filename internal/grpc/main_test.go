package grpc_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"test/internal/grpc"
	"test/mocks"
)

func GetService(t *testing.T) (usecase *grpc.FiboService, rep *mocks.MockIUsecase) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHedler := mocks.NewMockIUsecase(ctrl)

	service := grpc.NewFiboService(mockHedler)

	return service, mockHedler
}

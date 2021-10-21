package http_test

import (
	"test/mocks"
	"testing"

	"github.com/golang/mock/gomock"

	us "test/internal/http"
)

func GetService(t *testing.T) (usecase *us.Hendler, rep *mocks.MockIUsecase) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHedler := mocks.NewMockIUsecase(ctrl)

	service := us.NewHandler(mockHedler)

	return service, mockHedler
}

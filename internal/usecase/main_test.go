package usecase_test

import (
	"test/mocks"
	"testing"

	"github.com/golang/mock/gomock"

	us "test/internal/usecase"
)

func GetService(t *testing.T) (usecase *us.Service, rep *mocks.MockRedisInterface) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRedisInterface(ctrl)

	service := us.New(&us.Config{Repository: mockRepo})

	return service, mockRepo
}

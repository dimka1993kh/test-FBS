package usecase

import (
	"context"
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"

	"test/internal/repository"
)

const (
	firstNumber  = 1
	secondNumber = 2
	baseTen      = 10
	bitSize      = 64
)

//go:generate mockgen -package=mocks -source=./../../internal/usecase/usecase.go -destination=./../../mocks/Usecase.go
type IUsecase interface {
	Fib(ctx context.Context, x, y string) ([]uint64, error)
}

type Config struct {
	Repository repository.RedisInterface
}

type Service struct {
	repositpry repository.RedisInterface
}

func (s *Service) Fib(ctx context.Context, x, y string) ([]uint64, error) {
	newX, err := strconv.Atoi(x)
	if err != nil {
		return nil, errors.New("ошибка Х: Х должно быть числом")
	}

	newY, err := strconv.Atoi(y)
	if err != nil {
		return nil, errors.New("ошибка Y: Y должно быть числом")
	}

	if newX <= 0 || newY <= 0 {
		return nil, errors.New("порядковый номер должен быть больше 0")
	}

	if newX == 1 && newY == 2 {
		return []uint64{0, 1}, nil
	}

	if newX >= newY {
		return nil, errors.New("error: x > y")
	}

	res := make([]uint64, 0, newY-newX)

	for i := newX; i <= newY; i++ {
		res = append(res, s.fiboNumber(ctx, i))
	}

	log.Logger.Info().Msgf("fibo response: %d", res)

	return res, nil
}

func (s *Service) fiboNumber(ctx context.Context, serialNumber int) uint64 {
	var redisNumber bool

	res := []uint64{0, 1}
	serialNumber--

	if serialNumber == 0 {
		return 0
	}

	if serialNumber == 1 {
		return 1
	}

	val, err := s.repositpry.HGet(ctx, strconv.Itoa(serialNumber))
	if err == nil {
		redisNumber = true
	}

	if redisNumber {
		n, err := strconv.ParseInt(val, baseTen, bitSize)
		if err != nil {
			log.Logger.Fatal().Err(err)
		}

		log.Logger.Info().Msgf("Элемент %d последовательности Фибоначчи получен из кэша", serialNumber+1)

		return uint64(n)
	}

	for i := 0; i <= serialNumber-secondNumber; i++ {
		lastSerialNumber := len(res) - firstNumber
		penultimateSerialNumber := len(res) - secondNumber
		number := res[lastSerialNumber] + res[penultimateSerialNumber]
		res = append(res, number)
	}

	err = s.repositpry.HSet(ctx, strconv.Itoa(serialNumber), res[len(res)-1])
	if err != nil {
		log.Logger.Fatal().Err(err)
	}

	log.Logger.Info().Msgf("Элемент %d последовательности Фибоначчи рассчитан", serialNumber+1)

	return res[len(res)-1]
}

func New(cfg *Config) *Service {
	return &Service{
		repositpry: cfg.Repository,
	}
}

// Todo: покрыть тестами
// Todo: написать ридми

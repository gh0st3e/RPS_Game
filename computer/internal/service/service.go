package service

import (
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type Service struct {
	logger *logrus.Logger
}

func NewService(logger *logrus.Logger) *Service {
	return &Service{logger: logger}
}

func (s *Service) GenerateValue() string {
	var value string
	rand.Seed(time.Now().UnixNano())
	iValue := rand.Intn(3) + 1
	switch iValue {
	case 1:
		value = "r"
	case 2:
		value = "s"
	default:
		value = "p"
	}
	s.logger.Printf("[COMPUTER SERVICE] Computer Value: %s", value)
	return value
}

package service

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type Service struct {
	logger *logrus.Logger
}

func NewService(logger *logrus.Logger) *Service {
	return &Service{logger: logger}
}

func (s *Service) CheckValue(value string) error {
	s.logger.Printf("[USER SERVICE] UserValue: %s", value)
	if value != "s" && value != "p" && value != "r" {
		s.logger.Printf("incorrect input:%s", value)
		return fmt.Errorf("you can pick only r(rock),s(scissors),p(paper)")
	}
	s.logger.Printf("[USER SERVICE] User Choosed Correct Value %s!", value)
	return nil
}

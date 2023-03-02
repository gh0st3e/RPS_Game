package service

import "github.com/sirupsen/logrus"

const (
	rock     = "r"
	scissors = "s"
	paper    = "p"

	userWin     = "User Win"
	computerWin = "Computer Win"
	drawn       = "Drawn"
)

type Service struct {
	logger *logrus.Logger
}

func NewService(logger *logrus.Logger) *Service {
	return &Service{logger: logger}
}

func (s *Service) WhoIsWinner(userValue, computerValue string) string {
	s.logger.Printf("[RESULT SERVICE] WhoIsWinner started; UserValue=%s; ComputerValue=%s;", userValue, computerValue)
	if userValue == computerValue {
		return drawn
	}
	switch userValue {
	case rock:
		if computerValue == paper {
			return computerWin
		} else {
			return userWin
		}
	case paper:
		if computerValue == scissors {
			return computerWin
		} else {
			return userWin
		}
	case scissors:
		if computerValue == rock {
			return computerWin
		} else {
			return userWin
		}
	}
	return drawn
}

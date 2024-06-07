package request

import (
	"errors"
	"fmt"
	"strconv"
)

type FizzBuzzRequest struct {
	FromStr string `json:"from"`
	ToStr   string `json:"to"`
	From    int
	To      int
}

func (s *FizzBuzzRequest) Validate() error {
	from, err := strconv.Atoi(s.FromStr)
	if err != nil {
		return err
	}
	s.From = from
	to, err := strconv.Atoi(s.ToStr)
	if err != nil {
		return err
	}
	s.To = to
	if s.From > s.To {
		return errors.New(fmt.Sprintf("from(%d) must be greater than to(%d)", s.From, s.To))
	}
	if s.To-s.From > 100 {
		return errors.New(fmt.Sprintf("range of from(%d) and to(%d) cannot be greater than 100", s.From, s.To))
	}

	return nil
}

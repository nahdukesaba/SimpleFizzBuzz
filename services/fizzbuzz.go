package services

import (
	"context"
	"fmt"
	"strings"

	"SimpleFizzBuzz/services/request"
)

type ApisService struct {
}

type Apis interface {
	HandleFizzBuzz(ctx context.Context, form *request.FizzBuzzRequest) (string, error)
}

func NewApisService() Apis {
	return &ApisService{}
}

func (as *ApisService) HandleFizzBuzz(ctx context.Context, form *request.FizzBuzzRequest) (string, error) {
	var result []string
	if err := form.Validate(); err != nil {
		return "", err
	}

	for i := form.From; i <= form.To; i++ {
		result = append(result, SingleFizzBuzz(i))
	}

	return strings.Join(result, " "), nil
}

func SingleFizzBuzz(ins int) string {
	if ins%3 == 0 && ins%5 == 0 {
		return "FizzBuzz"
	} else if ins%3 == 0 {
		return "Fizz"
	} else if ins%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf("%d", ins)
}

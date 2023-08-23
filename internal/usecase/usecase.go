package usecase

import (
	"fizzbuzz/internal/value_objects"
	"strconv"
)

type rule struct{}

func NewRule() rule {
	return rule{}
}

func (r rule) Convert(number int) (value_objects.Tag, error) {
	var result value_objects.Tag

	if number%15 == 0 {
		result = "FizzBuzz"
	} else if number%3 == 0 {
		result = "Fizz"
	} else if number%5 == 0 {
		result = "Buzz"
	} else {
		result = value_objects.Tag(strconv.Itoa(number))
	}

	return result, nil
}

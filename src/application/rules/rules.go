package rules

import (
	"fizzbuzz/src/domain/value_objects"
)

type strategy interface {
	IsTruthy(int) bool
}

type tagRule struct {
	tag      value_objects.Tag
	strategy strategy
}

func NewRule(tag value_objects.Tag, strategy strategy) tagRule {
	return tagRule{tag: tag, strategy: strategy}
}

func (r tagRule) Tag() value_objects.Tag {
	return r.tag
}

func (r tagRule) IsSuccess(number int) bool {
	return r.strategy.IsTruthy(number)
}

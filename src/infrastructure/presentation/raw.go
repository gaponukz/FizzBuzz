package presentation

import (
	"fizzbuzz/src/domain/value_objects"
)

type rawPresenter struct{}

func NewRawPresenter() rawPresenter {
	return rawPresenter{}
}

func (p rawPresenter) Present(tag value_objects.Tag) []byte {
	return []byte(tag)
}

package controller

import (
	"fizzbuzz/src/domain/value_objects"
	"net/http"
	"strconv"
)

type converter interface {
	Find(int, value_objects.Tag) value_objects.Tag
}

type presenter interface {
	Present(value_objects.Tag) []byte
}

type controller struct {
	service   converter
	presenter presenter
}

func NewController(service converter, presenter presenter) controller {
	return controller{service: service, presenter: presenter}
}

func (c controller) Convert(responseWriter http.ResponseWriter, request *http.Request) {
	body := request.URL.Query().Get("number")
	number, err := strconv.Atoi(body)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	tag := c.service.Find(number, value_objects.Tag(body))

	responseWriter.Write(c.presenter.Present(tag))
}

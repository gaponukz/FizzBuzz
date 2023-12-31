package presentation

import (
	"encoding/json"
	"fizzbuzz/src/domain/value_objects"
)

type jsonPresenter struct{}

func NewJSONPresenter() jsonPresenter {
	return jsonPresenter{}
}

func (p jsonPresenter) Present(tag value_objects.Tag) []byte {
	result := struct {
		Result string `json:"result"`
	}{Result: string(tag)}

	b, err := json.Marshal(result)
	if err != nil {
		return []byte("")
	}

	return b
}

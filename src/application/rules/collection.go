package rules

import "fizzbuzz/src/domain/value_objects"

type rule interface {
	Tag() value_objects.Tag
	IsSuccess(int) bool
}

type rulesCollection struct {
	collection []rule
}

func NewRulesCollection(collection ...rule) rulesCollection {
	return rulesCollection{collection: collection}
}

func (r rulesCollection) Find(num int, defaultValue value_objects.Tag) value_objects.Tag {
	for _, rule := range r.collection {
		if rule.IsSuccess(num) {
			return rule.Tag()
		}
	}

	return defaultValue
}

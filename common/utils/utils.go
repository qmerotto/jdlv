package utils

import (
	"fmt"
	"reflect"
)

type Array[T any] []T

func (a Array[T]) Map(key string) (res []any, err error) {

	for i := 0; i < len(a); i++ {
		t := reflect.TypeOf(a[0])
		if t.Kind() != reflect.Struct {
			return nil, fmt.Errorf("invalid type")
		}

		numFields := t.NumField()
		for j := 0; j < numFields; j++ {
			fieldName := reflect.TypeOf(a[i]).Field(j).Name
			if fieldName == key {
				res = append(res, reflect.ValueOf(a[i]).String())
			}
		}
	}

	return
}

func Ter(c bool, t interface{}, f interface{}) interface{} {
	if c {
		return t
	} else {
		return f
	}
}

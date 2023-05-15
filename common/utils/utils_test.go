package utils

import (
	"fmt"
	"testing"
)

type RandStruct struct {
	ID string
}

func TestMap(t *testing.T) {
	v := Array[RandStruct]{{ID: "1"}, {ID: "2"}, {ID: "3"}}
	res, _ := v.Map("ID")
	fmt.Print(res)

	tt := Array[interface{}]{
		0, "fd", RandStruct{},
	}
	res, _ = tt.Map("")
}

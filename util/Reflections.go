package util

import (
	"reflect"
)

func TypeofObject(variable interface{}) string {
	t := reflect.TypeOf(variable).String()
	return t
}

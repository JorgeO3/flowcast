package utils

import (
	"reflect"
)

// IsStructEmpty checks if a struct is empty.
func IsStructEmpty(s interface{}) bool {
	return reflect.DeepEqual(s, reflect.Zero(reflect.TypeOf(s)).Interface())
}

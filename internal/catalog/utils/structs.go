package utils

import "reflect"

func IsStructEmpty(s interface{}) bool {
	// Verifica si el valor es de tipo struct
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Recorre los campos de la struct
	for i := 0; i < val.NumField(); i++ {
		if !reflect.DeepEqual(val.Field(i).Interface(), reflect.Zero(val.Field(i).Type()).Interface()) {
			return false
		}
	}

	return true
}

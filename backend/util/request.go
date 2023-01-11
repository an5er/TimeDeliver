package util

import (
	"reflect"
)

func CheckNull(u interface{}) bool {

	v := reflect.ValueOf(u)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			if v.Field(i).Type().Kind() == reflect.Struct {
				structField := v.Field(i).Type()
				for j := 0; j < structField.NumField(); j++ {
					if v.Field(i).Field(j).Interface() == "" {
						return false
					}
				}
				continue
			}
			if v.Field(i).Interface() == nil {
				return false
			}
		}
	}
	return true
}

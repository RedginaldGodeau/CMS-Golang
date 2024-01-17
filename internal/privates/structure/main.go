package structure

import "reflect"

func StructToArray(s interface{}) []any {
	val := reflect.ValueOf(s)
	var result []any
	for i := 0; i < val.NumField(); i++ {
		result = append(result, val.Field(i).Interface())
	}

	return result
}

package utils

import "reflect"

func ConvertToDTO(origin interface{}, target interface{}) (interface{}, error) {
	elem := reflect.ValueOf(target).Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Type().Field(i)
		sourceField := reflect.ValueOf(origin).Elem().FieldByName(field.Name)
		if !sourceField.IsValid() {
			continue
		}

		destField := elem.Field(i)

		if !destField.IsValid() {
			continue
		}

		if !destField.CanSet() {
			continue
		}

		destField.Set(sourceField)
	}
	return target, nil
}

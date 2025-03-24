package hw09structvalidator

import (
	"fmt"
	"reflect"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	panic("implement me")
}

func Validate(v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("ожидалась структура, получено: %s", val.Kind())
	}

	var errors ValidationErrors

	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Type().Field(i)
		fieldValue := val.Field(i)

		validateTag := field.Tag.Get("validate")

		if validateTag == "" {
			continue
		}

		fmt.Printf("Fild: %s, Tag: %s, Value: %v\n", field.Name, validateTag, fieldValue.Interface())

	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

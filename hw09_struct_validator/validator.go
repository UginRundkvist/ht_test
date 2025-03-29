package hw09structvalidator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var errMsgs []string
	for _, err := range v {
		errMsgs = append(errMsgs, fmt.Sprintf("%s: %s", err.Field, err.Err))
	}
	return strings.Join(errMsgs, ", ")
}

func Validate(v interface{}) error {
	var errors ValidationErrors
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct && val.Kind() != reflect.Ptr {
		return fmt.Errorf("ожиалась структура или указатель %s", val.Kind())
	}

	if val.Kind() == reflect.Ptr && val.IsNil() {
		return nil
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		validateTag := field.Tag.Get("validate")
		if validateTag == "" {
			continue
		}

		fieldValue := val.Field(i)
		fieldName := field.Name

		rules := strings.Split(validateTag, "|")
		for _, rule := range rules {
			parts := strings.SplitN(rule, ":", 2)
			validatorName := parts[0]
			var validatorArgs string
			if len(parts) > 1 {
				validatorArgs = parts[1]
			}

			switch fieldValue.Kind() {
			case reflect.Int:
				num := fieldValue.Int()
				if err := validateInt(num, validatorName, validatorArgs); err != nil {
					errors = append(errors, ValidationError{Field: fieldName, Err: err})
				}
			case reflect.String:
				str := fieldValue.String()
				if err := validateString(str, validatorName, validatorArgs); err != nil {
					errors = append(errors, ValidationError{Field: fieldName, Err: err})
				}
			case reflect.Slice:
				err := validateSlice(fieldValue, fieldName, validatorArgs, errors, validatorName)
				if err != nil {
					errors = append(errors, ValidationError{Field: fieldName, Err: err})
				}
			default:
				fmt.Printf("Неверный тип поля: %s\n", fieldValue.Kind())
			}
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func validateSlice(fieldValue reflect.Value, fieldName string, validatorArgs string, errors ValidationErrors, validatorName string) error {
	for j := 0; j < fieldValue.Len(); j++ {
		elementValue := fieldValue.Index(j)
		switch elementValue.Kind() {
		case reflect.Int:
			num := elementValue.Int()
			if err := validateInt(num, validatorName, validatorArgs); err != nil {
				errors = append(errors, ValidationError{Field: fieldName, Err: err})
			}
		case reflect.String:
			str := elementValue.String()
			if err := validateString(str, validatorName, validatorArgs); err != nil {
				errors = append(errors, ValidationError{Field: fieldName, Err: err})
			}
		default:
			fmt.Printf("Неверный элемент слайса: %s\n", elementValue.Kind())
		}
	}
	return nil
}

func validateInt(num int64, validatorName, validatorArgs string) error {
	switch validatorName {
	case "min":
		minValue, err := strconv.ParseInt(validatorArgs, 10, 64)
		if err != nil {
			return fmt.Errorf("неверное минимальное значение %w", err)
		}
		if num < minValue {
			return fmt.Errorf("должно быть больше %d", minValue)
		}
	case "max":
		maxValue, err := strconv.ParseInt(validatorArgs, 10, 64)
		if err != nil {
			return fmt.Errorf("неверное максимальное значение: %w", err)
		}
		if num > maxValue {
			return fmt.Errorf("должно быть меньше %d", maxValue)
		}
	case "in":
		allowedValues := strings.Split(validatorArgs, ",")
		found := false
		for _, val := range allowedValues {
			intValue, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return fmt.Errorf("неверное значение для 'in' %w", err)
			}
			if num == intValue {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("должно быть в: %s", validatorArgs)
		}
	default:
		return fmt.Errorf("неизвесный валидатор %s", validatorName)
	}
	return nil
}

func validateString(str string, validatorName, validatorArgs string) error {
	switch validatorName {
	case "len":
		expectedLen, err := strconv.Atoi(validatorArgs)
		if err != nil {
			return fmt.Errorf("неверное строковое значение: %w", err)
		}
		if len(str) != expectedLen {
			return fmt.Errorf("длинна должна быть %d", expectedLen)
		}
	case "regexp":
		re, err := regexp.Compile(validatorArgs)
		if err != nil {
			return fmt.Errorf("неверный regexp: %w", err)
		}
		if !re.MatchString(str) {
			return fmt.Errorf("должно соответствовать регулярному выражению '%s'", validatorArgs)
		}
	case "in":
		allowedValues := strings.Split(validatorArgs, ",")
		found := false
		for _, val := range allowedValues {
			if str == val {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("должно быть одно из: %s", validatorArgs)
		}
	default:
		return fmt.Errorf("неизвесный валидатор: %s", validatorName)
	}
	return nil
}

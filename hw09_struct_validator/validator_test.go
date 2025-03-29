package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

type UserRole string

type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	// }.

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "a1b2c3d4-e5f6-7890-1234-567890abcdef",
				Age:    25,
				Email:  "test@example.com",
				Role:   "admin",
				Phones: []string{"12345678901", "98765432109"},
			},
			expectedErr: nil,
		},
		{
			in: App{
				Version: "0.0.1.0",
			},
			expectedErr: ValidationErrors{{Field: "Version", Err: fmt.Errorf("длина должна быть %d", 5)}},
		},
		{
			in: Response{
				Code: 2000,
				Body: "OK",
			},
			expectedErr: ValidationErrors{{Field: "Code", Err: fmt.Errorf("должно быть в: %d,%d,%d", 200, 404, 500)}},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.in)
			if !compareErrors(err, tt.expectedErr) {
				t.Errorf("ошибка валидации: %v, ожидаемая ошибка: %v", err, tt.expectedErr)
			}
		})
	}
}

func compareErrors(actual, expected error) bool {
	if (actual != nil) != (expected != nil) {
		return false
	}

	if expected == nil {
		return true
	}

	var expectedVErr ValidationErrors
	if errors.As(expected, &expectedVErr) {
		var actualVErr ValidationErrors
		if !errors.As(actual, &actualVErr) {
			return false
		}

		if len(actualVErr) != len(expectedVErr) {
			return false
		}

		for i := range expectedVErr {
			if actualVErr[i].Err.Error() != expectedVErr[i].Err.Error() || actualVErr[i].Field != expectedVErr[i].Field {
				return false
			}
		}
		return true
	}

	return errors.Is(actual, expected)
}

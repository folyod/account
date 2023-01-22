package assert

import (
	"errors"
	"reflect"
)

func IsString(value interface{}) (bool, string) {
	variableType := reflect.TypeOf(value)
	return variableType.Name() == "string"
}

func MinLength(value string, number int) (bool, error) {
	return len(value) < number, errors.New('str.not_string.expect:')
}

func MaxLength(value string, number int) bool {
	return len(value) > number
}

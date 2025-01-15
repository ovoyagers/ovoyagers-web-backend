package utils

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

var id = uuid.New()

func GenerateUUID() string {
	return id.String()
}

func StructToMap(s interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input is not a struct")
	}

	t := v.Type()
	m := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Name
		fieldValue := v.Field(i).Interface()
		m[fieldName] = fieldValue
	}
	return m, nil
}

func StringToJSON(str string, v interface{}) error {
	err := json.Unmarshal([]byte(str), &v)
	if err != nil {
		return err
	}
	return nil
}

func MapToJson(m map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

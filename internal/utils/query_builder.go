package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type QueryBuilder struct {
	Values map[string]interface{}
	Errors []string
}

func (q *QueryBuilder) AddRequiredValue(key string, value interface{}) {
	if reflect.ValueOf(value).IsZero() {
		errStr := fmt.Sprintf("%s cannot be empty", key)
		q.addError(errStr)
		return
	}

	q.AddValue(key, value)
}

func (q *QueryBuilder) addError(err string) {
	q.Errors = append(q.Errors, err)
}

func (q *QueryBuilder) HasError() bool {
	return len(q.Errors) > 0
}

func (q *QueryBuilder) GetError() error {
	errors := q.Errors
	q.Errors = []string{}
	return fmt.Errorf("QueryBuilder error: %s", strings.Join(errors, ", "))
}

func (q *QueryBuilder) AddValue(key string, value interface{}) {
	if reflect.ValueOf(value).IsZero() {
		return
	}

	switch value.(type) {
	case string:
		// no need to transform
	case int:
		// no need to transform
	case int64:
		// no need to transform
	default:
		value, _ = marshalString(value)
	}

	q.Values[key] = value
}

func marshalString(d interface{}) (string, error) {
	byteMetrics, err := json.Marshal(d)
	if err != nil {
		return "", fmt.Errorf("marshal err: %w", err)
	}
	result := string(byteMetrics)
	return result, nil
}

func (q *QueryBuilder) GetKeysAndValues() ([]string, []interface{}) {
	var keys []string
	var values []interface{}

	for k, v := range q.Values {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

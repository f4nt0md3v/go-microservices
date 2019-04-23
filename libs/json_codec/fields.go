package json_codec

import (
	"errors"
	"fmt"
)

// Get and check isExist field which has ANY type.
func GetAny(field string, fields map[string]interface{}) (interface{}, error) {
	fieldPure, okPureField := fields[field]
	if !okPureField {
		return nil, errors.New(fmt.Sprintf("Cannot get ANY field '%s'", field))
	}
	return fieldPure, nil
}

// Get and check isExist field which has STRING type.
func GetString(field string, fields map[string]interface{}) (string, error) {
	fieldPure, okPureField := fields[field]
	if !okPureField {
		return "", errors.New(fmt.Sprintf("Cannot get STRING field '%s'", field))
	}
	fieldValue, okFieldValue := fieldPure.(string)
	if !okFieldValue {
		return "", errors.New(fmt.Sprintf("Cannot convert field '%s' to STRING", field))
	}
	return fieldValue, nil
}

// Get and check isExist field which has BOOL type.
func GetBool(field string, fields map[string]interface{}) (bool, error) {
	fieldPure, okPureField := fields[field]
	if !okPureField {
		return false, errors.New(fmt.Sprintf("Cannot get BOOL field '%s'", field))
	}
	fieldValue, okFieldValue := fieldPure.(bool)
	if !okFieldValue {
		return false, errors.New(fmt.Sprintf("Cannot convert field '%s' to BOOL", field))
	}
	return fieldValue, nil
}

// Get and check isExist field which has INT64 type.
func GetInt64(field string, fields map[string]interface{}) (int64, error) {
	fieldPure, okPureField := fields[field]
	if !okPureField {
		return 0, errors.New(fmt.Sprintf("Cannot get INT64 field '%s'", field))
	}
	fieldValue, okFieldValue := fieldPure.(float64)
	if !okFieldValue {
		return 0, errors.New(fmt.Sprintf("Cannot convert field '%s' to INT6", field))
	}
	return int64(fieldValue), nil
}

// Get and check isExist field which has ARRAY STRING type.
func GetArrayString(field string, fields map[string]interface{}) ([]string, error) {
	fieldPure, okPureField := fields[field]
	if !okPureField {
		return nil, errors.New(fmt.Sprintf("Cannot get ARRAY STRING field '%s'", field))
	}
	fieldValueArray, okFieldValueArray := fieldPure.([]interface{})
	if !okFieldValueArray {
		return nil, errors.New(fmt.Sprintf("Cannot convert field '%s' to ARRAY STRING", field))
	}
	fieldValueArrayString := []string{}
	for _, v := range fieldValueArray {
		temp, okTemp := v.(string)
		if !okTemp {
			return nil, errors.New(fmt.Sprintf("Cannot convert field '%s' to ARRAY STRING", field))
		}
		fieldValueArrayString = append(fieldValueArrayString, temp)
	}

	return fieldValueArrayString, nil
}

// Get and check isExist field which has JSON type.
func GetJson(field string, fields map[string]interface{}) (map[string]interface{}, error) {
	fieldPure, okPureField := fields[field]
	if !okPureField {
		return nil, errors.New(fmt.Sprintf("Cannot get JSON field '%s'", field))
	}
	fieldValue, okFieldValue := fieldPure.(map[string]interface{})
	if !okFieldValue {
		return nil, errors.New(fmt.Sprintf("Cannot convert field '%s' to JSON", field))
	}
	return fieldValue, nil
}

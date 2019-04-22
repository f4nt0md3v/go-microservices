// Package implement logic encode/decode JSON/JSON ARRAY.
package json_codec

import (
	"encoding/json"
	"go-microservices/libs/errors_handler"
)

// Decode json string to map[string]interface{}.
func JsonParse(source string) (map[string]interface{}, errors_handler.Error) {
	blob := []byte(source)
	var param map[string]interface{}
	err := json.Unmarshal(blob, &param)
	if err != nil {
		return nil, errors_handler.GetError(11, map[string]interface{}{
			"json": source,
		})
	}

	return param, nil
}

// Decode json string to []map[string]interface{}.
func JsonArrayParse(source string) ([]map[string]interface{}, errors_handler.Error) {
	blob := []byte(source)
	var param []map[string]interface{}
	err := json.Unmarshal(blob, &param)
	if err != nil {
		return nil, errors_handler.GetError(12, map[string]interface{}{
			"json_array": source,
		})
	}

	return param, nil
}

// Encode map[string]interface{} to json string.
func JsonEncode(source map[string]interface{}) (string, errors_handler.Error) {
	result, err := json.Marshal(source)
	if err != nil {
		return "", errors_handler.GetError(13, map[string]interface{}{
			"json": source,
		})
	}

	return string(result), nil
}

// Encode []map[string]interface{} to json array string.
func JsonArrayEncode(source []map[string]interface{}) (string, errors_handler.Error) {
	result, err := json.Marshal(source)
	if err != nil {
		return "", errors_handler.GetError(14, map[string]interface{}{
			"json_array": source,
		})
	}

	return string(result), nil
}

package conv

import "errors"

func ToStringMap(value interface{}) (map[string]interface{}, error) {
	switch value.(type) {
	case map[string]interface{}:
		return value.(map[string]interface{}), nil
	default:
		return nil, errors.New("unsupported type conversion")
	}
}

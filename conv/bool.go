package conv

import "errors"

func ToBool(value interface{}) (bool, error) {
	switch value.(type) {
	case string:
		if value.(string) == "true" {
			return true, nil
		}
		if value.(string) == "false" {
			return false, nil
		}
		return false, errors.New("invalid string to bool conversion")
	case bool:
		return value.(bool), nil
	default:
		return false, errors.New("unsupported conversion type")
	}
}

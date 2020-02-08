package conv

import "errors"

// ToStringSlice Convert value to []string
func ToStringSlice(value interface{}) ([]string, error) {
	switch value.(type) {
	case []string:
		return value.([]string), nil
	case []interface{}:
		v := value.([]interface{})
		result := []string{}
		for _, item := range v {
			result = append(result, ToString(item))
		}
		return result, nil
	default:
		return nil, errors.New("unsupported type conversion")
	}
}

// ToUint32Slice Convert value to []uint32
func ToUint32Slice(value interface{}) ([]uint32, error) {
	switch value.(type) {
	case []uint32:
		return value.([]uint32), nil
	case []interface{}:
		v := value.([]interface{})
		result := []uint32{}
		for _, item := range v {
			num, err := ToUint32(item)
			if err != nil {
				return nil, err
			}
			result = append(result, num)
		}
		return result, nil
	default:
		return nil, errors.New("unsupported type conversion")
	}
}

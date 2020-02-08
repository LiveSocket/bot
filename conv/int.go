package conv

import (
	"errors"
	"strconv"
)

// ToInt Convert value to int
func ToInt(value interface{}) (int, error) {
	switch value.(type) {
	case string:
		return strconv.Atoi(value.(string))
	case int:
		return value.(int), nil
	case int8:
		return int(value.(int8)), nil
	case int16:
		return int(value.(int16)), nil
	case int32:
		return int(value.(int32)), nil
	case int64:
		return int(value.(int64)), nil
	case uint:
		return int(value.(uint)), nil
	case uint8:
		return int(value.(uint8)), nil
	case uint16:
		return int(value.(uint16)), nil
	case uint32:
		return int(value.(uint32)), nil
	case uint64:
		return int(value.(uint64)), nil
	case float32:
		return int(value.(float32)), nil
	case float64:
		return int(value.(float64)), nil
	default:
		return 0, errors.New("unsupported conversion type")
	}
}

// ToInt8 Converts value to int8
func ToInt8(value interface{}) (int8, error) {
	switch value.(type) {
	case string:
		r, err := strconv.ParseInt(value.(string), 10, 8)
		return int8(r), err
	case int:
		return int8(value.(int)), nil
	case int8:
		return value.(int8), nil
	case int16:
		return int8(value.(int16)), nil
	case int32:
		return int8(value.(int32)), nil
	case int64:
		return int8(value.(int64)), nil
	case uint:
		return int8(value.(uint)), nil
	case uint8:
		return int8(value.(uint8)), nil
	case uint16:
		return int8(value.(uint16)), nil
	case uint32:
		return int8(value.(uint32)), nil
	case uint64:
		return int8(value.(uint64)), nil
	case float32:
		return int8(value.(float32)), nil
	case float64:
		return int8(value.(float64)), nil
	default:
		return 0, errors.New("unsupported conversion type")
	}
}

// ToInt16 Converts value to int16
func ToInt16(value interface{}) (int16, error) {
	switch value.(type) {
	case string:
		r, err := strconv.ParseInt(value.(string), 10, 16)
		return int16(r), err
	case int:
		return int16(value.(int)), nil
	case int8:
		return int16(value.(int8)), nil
	case int16:
		return value.(int16), nil
	case int32:
		return int16(value.(int32)), nil
	case int64:
		return int16(value.(int64)), nil
	case uint:
		return int16(value.(uint)), nil
	case uint8:
		return int16(value.(uint8)), nil
	case uint16:
		return int16(value.(uint16)), nil
	case uint32:
		return int16(value.(uint32)), nil
	case uint64:
		return int16(value.(uint64)), nil
	case float32:
		return int16(value.(float32)), nil
	case float64:
		return int16(value.(float64)), nil
	default:
		return 0, errors.New("unsupported conversion type")
	}
}

// ToInt32 Converts value to int32
func ToInt32(value interface{}) (int32, error) {
	switch value.(type) {
	case string:
		r, err := strconv.ParseInt(value.(string), 10, 32)
		return int32(r), err
	case int:
		return int32(value.(int)), nil
	case int8:
		return int32(value.(int8)), nil
	case int16:
		return int32(value.(int16)), nil
	case int32:
		return value.(int32), nil
	case int64:
		return int32(value.(int64)), nil
	case uint:
		return int32(value.(uint)), nil
	case uint8:
		return int32(value.(uint8)), nil
	case uint16:
		return int32(value.(uint16)), nil
	case uint32:
		return int32(value.(uint32)), nil
	case uint64:
		return int32(value.(uint64)), nil
	case float32:
		return int32(value.(float32)), nil
	case float64:
		return int32(value.(float64)), nil
	default:
		return 0, errors.New("unsupported conversion type")
	}
}

// ToInt64 Converts value to int64
func ToInt64(value interface{}) (int64, error) {
	switch value.(type) {
	case string:
		r, err := strconv.ParseInt(value.(string), 10, 64)
		return int64(r), err
	case int:
		return int64(value.(int)), nil
	case int8:
		return int64(value.(int8)), nil
	case int16:
		return int64(value.(int16)), nil
	case int32:
		return int64(value.(int32)), nil
	case int64:
		return value.(int64), nil
	case uint:
		return int64(value.(uint)), nil
	case uint8:
		return int64(value.(uint8)), nil
	case uint16:
		return int64(value.(uint16)), nil
	case uint32:
		return int64(value.(uint32)), nil
	case uint64:
		return int64(value.(uint64)), nil
	case float32:
		return int64(value.(float32)), nil
	case float64:
		return int64(value.(float64)), nil
	default:
		return 0, errors.New("unsupported conversion type")
	}
}

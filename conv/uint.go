package conv

import (
	"errors"
	"strconv"
)

// ToUint Converts value to uint
func ToUint(value interface{}) (uint, error) {
	switch value.(type) {
	case string:
		r, err := strconv.ParseUint(value.(string), 10, 64)
		return uint(r), err
	case int:
		return uint(value.(int)), nil
	case int8:
		return uint(value.(int8)), nil
	case int16:
		return uint(value.(int16)), nil
	case int32:
		return uint(value.(int32)), nil
	case int64:
		return uint(value.(int64)), nil
	case uint:
		return value.(uint), nil
	case uint8:
		return uint(value.(uint8)), nil
	case uint16:
		return uint(value.(uint16)), nil
	case uint32:
		return uint(value.(uint32)), nil
	case uint64:
		return uint(value.(uint64)), nil
	case float32:
		return uint(value.(float32)), nil
	case float64:
		return uint(value.(float64)), nil
	default:
		return 0, errors.New("unsupported conversion type")
	}
}

// ToUint8 Converts value to uint8
func ToUint8(value interface{}) (uint8, error) {
	switch value.(type) {
	case string:
		r, err := strconv.ParseUint(value.(string), 10, 8)
		return uint8(r), err
	case int:
		return uint8(value.(int)), nil
	case int8:
		return uint8(value.(int8)), nil
	case int16:
		return uint8(value.(int16)), nil
	case int32:
		return uint8(value.(int32)), nil
	case int64:
		return uint8(value.(int64)), nil
	case uint:
		return uint8(value.(uint)), nil
	case uint8:
		return value.(uint8), nil
	case uint16:
		return uint8(value.(uint16)), nil
	case uint32:
		return uint8(value.(uint32)), nil
	case uint64:
		return uint8(value.(uint64)), nil
	case float32:
		return uint8(value.(float32)), nil
	case float64:
		return uint8(value.(float64)), nil
	default:
		return 0, errors.New("unsupported conversion type")
	}
}

// ToUint16 Converts value to uint16
func ToUint16(value interface{}) (uint16, error) {
	switch value.(type) {
	case string:
		r, err := strconv.ParseUint(value.(string), 10, 16)
		return uint16(r), err
	case int:
		return uint16(value.(int)), nil
	case int8:
		return uint16(value.(int8)), nil
	case int16:
		return uint16(value.(int16)), nil
	case int32:
		return uint16(value.(int32)), nil
	case int64:
		return uint16(value.(int64)), nil
	case uint:
		return uint16(value.(uint)), nil
	case uint8:
		return uint16(value.(uint8)), nil
	case uint16:
		return value.(uint16), nil
	case uint32:
		return uint16(value.(uint32)), nil
	case uint64:
		return uint16(value.(uint64)), nil
	case float32:
		return uint16(value.(float32)), nil
	case float64:
		return uint16(value.(float64)), nil
	default:
		return 0, errors.New("unsupported conversion type")
	}
}

// ToUint32 Converts value to uint32
func ToUint32(value interface{}) (uint32, error) {
	switch value.(type) {
	case string:
		r, err := strconv.ParseUint(value.(string), 10, 32)
		return uint32(r), err
	case int:
		return uint32(value.(int)), nil
	case int8:
		return uint32(value.(int8)), nil
	case int16:
		return uint32(value.(int16)), nil
	case int32:
		return uint32(value.(int32)), nil
	case int64:
		return uint32(value.(int64)), nil
	case uint:
		return uint32(value.(uint)), nil
	case uint8:
		return uint32(value.(uint8)), nil
	case uint16:
		return uint32(value.(uint16)), nil
	case uint32:
		return value.(uint32), nil
	case uint64:
		return uint32(value.(uint64)), nil
	case float32:
		return uint32(value.(float32)), nil
	case float64:
		return uint32(value.(float64)), nil
	default:
		return 0, errors.New("unsupported conversion type")
	}
}

// ToUint64 Converts value to uint64
func ToUint64(value interface{}) (uint64, error) {
	switch value.(type) {
	case string:
		r, err := strconv.ParseUint(value.(string), 10, 64)
		return uint64(r), err
	case int:
		return uint64(value.(int)), nil
	case int8:
		return uint64(value.(int8)), nil
	case int16:
		return uint64(value.(int16)), nil
	case int32:
		return uint64(value.(int32)), nil
	case int64:
		return uint64(value.(int64)), nil
	case uint:
		return uint64(value.(uint)), nil
	case uint8:
		return uint64(value.(uint8)), nil
	case uint16:
		return uint64(value.(uint16)), nil
	case uint32:
		return uint64(value.(uint32)), nil
	case uint64:
		return value.(uint64), nil
	case float32:
		return uint64(value.(float32)), nil
	case float64:
		return uint64(value.(float64)), nil
	default:
		return 0, errors.New("unsupported conversion type")
	}
}

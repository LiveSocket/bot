package conv

import (
	"strconv"
)

// ToString Converts the value to a string
func ToString(value interface{}) string {
	switch value.(type) {
	case string:
		return value.(string)
	case int:
		return strconv.FormatInt(int64(value.(int)), 10)
	case int8:
		return strconv.FormatInt(int64(value.(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(value.(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(value.(int32)), 10)
	case int64:
		return strconv.FormatInt(value.(int64), 10)
	case uint:
		return strconv.FormatUint(uint64(value.(uint)), 10)
	case uint8:
		return strconv.FormatUint(uint64(value.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(value.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(value.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(value.(uint64), 10)
	case float32:
		return strconv.FormatFloat(float64(value.(float32)), 'f', 10, 32)
	case float64:
		return strconv.FormatFloat(value.(float64), 'f', 10, 64)
	default:
		return ""
	}
}

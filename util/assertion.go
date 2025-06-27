package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// Type assert api for String()
type apiString interface {
	String() string
}

// Type assert api for Error().
type apiError interface {
	Error() string
}

func AnyToInt(t1 interface{}) int {
	switch t1.(type) {
	case uint:
		return int(t1.(uint))
	case int8:
		return int(t1.(int8))
	case uint8:
		return int(t1.(uint8))
	case int16:
		return int(t1.(int16))
	case uint16:
		return int(t1.(uint16))
	case int32:
		return int(t1.(int32))
	case uint32:
		return int(t1.(uint32))
	case int64:
		return int(t1.(int64))
	case uint64:
		return int(t1.(uint64))
	case float32:
		return int(t1.(float32))
	case float64:
		return int(t1.(float64))
	case string:
		t2, _ := strconv.Atoi(t1.(string))
		return t2
	default:
		return t1.(int)
	}
}

func AnyToUint(t1 interface{}) uint {
	switch t1.(type) {
	case int8:
		return uint(t1.(int8))
	case uint8:
		return uint(t1.(uint8))
	case int16:
		return uint(t1.(int16))
	case uint16:
		return uint(t1.(uint16))
	case int32:
		return uint(t1.(int32))
	case uint32:
		return uint(t1.(uint32))
	case int64:
		return uint(t1.(int64))
	case uint64:
		return uint(t1.(uint64))
	case float32:
		return uint(t1.(float32))
	case float64:
		return uint(t1.(float64))
	case string:
		t, _ := strconv.ParseUint(t1.(string), 10, 64)
		return uint(t)
	default:
		return t1.(uint)
	}
}

func AnyToString(i interface{}) string {
	if i == nil {
		return ""
	}
	switch value := i.(type) {
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	case time.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *time.Time:
		if value == nil {
			return ""
		}
		return value.String()
	default:
		// Empty checks.
		if value == nil {
			return ""
		}
		if f, ok := value.(apiString); ok {
			// If the variable implements the String() interface,
			// then use that interface to perform the conversion
			return f.String()
		} else if f, ok := value.(apiError); ok {
			// If the variable implements the ErrorResp() interface,
			// then use that interface to perform the conversion
			return f.Error()
		} else {
			// Reflect checks.
			rv := reflect.ValueOf(value)
			kind := rv.Kind()
			switch kind {
			case reflect.Chan,
				reflect.Map,
				reflect.Slice,
				reflect.Func,
				reflect.Ptr,
				reflect.Interface,
				reflect.UnsafePointer:
				if rv.IsNil() {
					return ""
				}
			}
			if kind == reflect.Ptr {
				return AnyToString(rv.Elem().Interface())
			}
			// Finally use json.Marshal to convert.
			if jsonContent, err := json.Marshal(value); err != nil {
				return fmt.Sprint(value)
			} else {
				return string(jsonContent)
			}
		}
	}
}

func Decimal(num float64) float64 {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return num
}

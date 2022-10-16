package convertTo

import "strconv"

func GetUint32Value(value interface{}) uint32 {
	var rs uint32
	switch value.(type) {
	case string:
		v, _ := strconv.Atoi(value.(string))
		rs = uint32(v)
	case int:
		rs = uint32(value.(int))
	case int8:
		rs = uint32(value.(int8))
	case int16:
		rs = uint32(value.(int16))
	case int32:
		rs = uint32(value.(int32))
	case int64:
		rs = uint32(value.(int64))
	case uint:
		rs = uint32(value.(uint))
	case uint8:
		rs = uint32(value.(uint8))
	case uint16:
		rs = uint32(value.(uint16))
	case uint32:
		rs = value.(uint32)
	case uint64:
		rs = uint32(value.(uint64))
	case float64:
		rs = uint32(value.(float64))
	case float32:
		rs = uint32(value.(float32))
	}
	return rs
}

func GetUint64Value(value interface{}) uint64 {
	var rs uint64
	switch value.(type) {
	case string:
		v, _ := strconv.ParseInt(value.(string), 10, 64)
		rs = uint64(v)
	case int:
		rs = uint64(value.(int))
	case int8:
		rs = uint64(value.(int8))
	case int16:
		rs = uint64(value.(int16))
	case int32:
		rs = uint64(value.(int32))
	case int64:
		rs = uint64(value.(int64))
	case uint:
		rs = uint64(value.(uint))
	case uint8:
		rs = uint64(value.(uint8))
	case uint16:
		rs = uint64(value.(uint16))
	case uint32:
		rs = uint64(value.(uint32))
	case uint64:
		rs = value.(uint64)
	case float64:
		rs = uint64(value.(float64))
	case float32:
		rs = uint64(value.(float32))
	}
	return rs
}

func GetUint8Value(value interface{}) uint8 {
	var rs uint8
	switch value.(type) {
	case string:
		v, _ := strconv.Atoi(value.(string))
		rs = uint8(v)
	case int:
		rs = uint8(value.(int))
	case int8:
		rs = uint8(value.(int8))
	case int16:
		rs = uint8(value.(int16))
	case int32:
		rs = uint8(value.(int32))
	case int64:
		rs = uint8(value.(int64))
	case uint:
		rs = uint8(value.(uint))
	case uint8:
		rs = value.(uint8)
	case uint16:
		rs = uint8(value.(uint16))
	case uint32:
		rs = uint8(value.(uint32))
	case uint64:
		rs = value.(uint8)
	case float64:
		rs = uint8(value.(float64))
	case float32:
		rs = uint8(value.(float32))
	}
	return rs
}

func GetInt32Value(value interface{}) int32 {
	var rs int32
	switch value.(type) {
	case string:
		v, _ := strconv.Atoi(value.(string))
		rs = int32(v)
	case int:
		rs = int32(value.(int))
	case int8:
		rs = int32(value.(int8))
	case int16:
		rs = int32(value.(int16))
	case int32:
		rs = value.(int32)
	case int64:
		rs = int32(value.(int64))
	case uint:
		rs = int32(value.(uint))
	case uint8:
		rs = int32(value.(uint8))
	case uint16:
		rs = int32(value.(uint16))
	case uint32:
		rs = int32(value.(uint32))
	case uint64:
		rs = int32(value.(uint64))
	case float64:
		rs = int32(value.(float64))
	case float32:
		rs = int32(value.(float32))
	}
	return rs
}

func GetUIntValue(value interface{}) uint {
	var rs uint
	switch value.(type) {
	case string:
		v, _ := strconv.Atoi(value.(string))
		rs = uint(v)
	case int:
		rs = uint(value.(int))
	case int8:
		rs = uint(value.(int8))
	case int16:
		rs = uint(value.(int16))
	case int32:
		rs = uint(value.(int32))
	case int64:
		rs = uint(value.(int64))
	case uint:
		rs = value.(uint)
	case uint8:
		rs = uint(value.(uint8))
	case uint16:
		rs = uint(value.(uint16))
	case uint32:
		rs = uint(value.(uint32))
	case uint64:
		rs = uint(value.(uint64))
	case float64:
		rs = uint(value.(float64))
	case float32:
		rs = uint(value.(float32))
	}
	return rs
}

func GetIntValue(value interface{}) int {
	rs := 0
	switch value.(type) {
	case string:
		rs, _ = strconv.Atoi(value.(string))
	case int:
		rs = value.(int)
	case int8:
		rs = int(value.(int8))
	case int16:
		rs = int(value.(int16))
	case int32:
		rs = int(value.(int32))
	case int64:
		rs = int(value.(int64))
	case uint:
		rs = int(value.(uint))
	case uint8:
		rs = int(value.(uint8))
	case uint16:
		rs = int(value.(uint16))
	case uint32:
		rs = int(value.(uint32))
	case uint64:
		rs = int(value.(uint64))
	case float64:
		rs = int(value.(float64))
	case float32:
		rs = int(value.(float32))
	}
	return rs
}

func GetFloat64Value(value interface{}) float64 {
	var rs float64
	switch value.(type) {
	case string:
		rs, _ = strconv.ParseFloat(value.(string), 64)
	case int:
		rs = float64(value.(int))
	case int8:
		rs = float64(value.(int8))
	case int16:
		rs = float64(value.(int16))
	case int32:
		rs = float64(value.(int32))
	case int64:
		rs = float64(value.(int64))
	case uint:
		rs = float64(value.(uint))
	case uint8:
		rs = float64(value.(uint8))
	case uint16:
		rs = float64(value.(uint16))
	case uint32:
		rs = float64(value.(uint32))
	case uint64:
		rs = float64(value.(uint64))
	case float64:
		rs = value.(float64)
	case float32:
		rs = float64(value.(float32))
	}
	return rs
}

func GetInt64Value(value interface{}) int64 {
	var rs int64
	switch value.(type) {
	case string:
		rs, _ = strconv.ParseInt(value.(string), 10, 64)
	case int:
		rs = int64(value.(int))
	case int8:
		rs = int64(value.(int8))
	case int16:
		rs = int64(value.(int16))
	case int32:
		rs = int64(value.(int32))
	case int64:
		rs = value.(int64)
	case uint:
		rs = int64(value.(uint))
	case uint8:
		rs = int64(value.(uint8))
	case uint16:
		rs = int64(value.(uint16))
	case uint32:
		rs = int64(value.(uint32))
	case uint64:
		rs = int64(value.(uint64))
	case float64:
		rs = int64(value.(float64))
	case float32:
		rs = int64(value.(float32))
	}
	return rs
}

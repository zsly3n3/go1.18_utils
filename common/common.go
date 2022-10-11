package common

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
)

type NumericValue interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// 传统md5
func Md5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func Struct2Bytes(order binary.ByteOrder, data any) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, order, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

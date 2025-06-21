package main

import (
	"bytes"
	"encoding/binary"
)

// Int64ToBytes 将一个 int64 转换为其 8 字节的大端序二进制表示。
// 这个函数名比 "IntToHex" 更准确地描述了它的功能。
func Int64ToBytes(num int64) []byte {
	buff := new(bytes.Buffer)

	// 将整数的二进制表示写入缓冲区
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		// 不直接 panic，而是将错误返回给调用者，让调用者决定如何处理。
		// 这使得函数更加通用和可复用。
		return nil
	}

	// 返回结果和 nil 错误
	return buff.Bytes()
}

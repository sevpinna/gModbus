package Comm

import (
	"bytes"
	"encoding/binary"
)

type Register [2]byte
type Registers []Register

// BytesCombine 字节切片合并函数
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
func Uint16ToByte(value uint16) (data []byte) {
	data = make([]byte, 2)
	binary.BigEndian.PutUint16(data, value)
	return
}
func (Rs Registers) Len() (Len uint16) {
	var r []Register = Rs
	return uint16(len(r))
}
func (R Register) ToByte(order string) (data []byte) {
	data = make([]byte, 2)
	switch order {
	case "AB":
		data[0], data[1] = R[0], R[1]
	case "BA":
		data[1], data[0] = R[0], R[1]
	}
	return
}
func (Rs Registers) ToByte(order string) (data []byte) {
	var r []Register = Rs
	data = make([]byte, len(r)*2)
	for _, item := range r {
		data = BytesCombine(item.ToByte(""))
	}
	return
}
func DataChange(Input []byte, Order string) (Data []byte) {
	switch Order {
	case "AB":
	case "BA":
	case "ABCD":
	case "CDAB":
	case "BADC":
	case "DCBA":
	}
	return
}

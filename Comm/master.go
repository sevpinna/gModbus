package Comm

import (
	"fmt"
	"math"
)

// BuildReadCoilStatus  构建读取线圈状态报文
func BuildReadCoilStatus(Id uint8, RegisterAddress, Length uint16) (Message []byte) {
	Message = BytesCombine([]byte{Id}, []byte{0x01}, Uint16ToByte(RegisterAddress), Uint16ToByte(Length))
	return
}

// BuildReadInputStatus 构建读取输入状态报文
func BuildReadInputStatus(Id uint8, RegisterAddress, Length uint16) (Message []byte) {
	Message = BytesCombine([]byte{Id}, []byte{0x02}, Uint16ToByte(RegisterAddress), Uint16ToByte(Length))
	return
}

// BuildReadHoldingRegister 构建读取保持寄存器报文
func BuildReadHoldingRegister(Id uint8, RegisterAddress, Length uint16) (Message []byte) {
	Message = BytesCombine([]byte{Id}, []byte{0x03}, Uint16ToByte(RegisterAddress), Uint16ToByte(Length))
	return
}

// BuildReadInputRegister 构建读取输入寄存器报文
func BuildReadInputRegister(Id uint8, RegisterAddress, Length uint16) (Message []byte) {
	Message = BytesCombine([]byte{Id}, []byte{0x04}, Uint16ToByte(RegisterAddress), Uint16ToByte(Length))
	return
}

// BuildWriteSingleCoilStatus 构建写单个线圈报文
func BuildWriteSingleCoilStatus(Id uint8, RegisterAddress uint16, Data bool) (Message []byte) {
	var data = make([]byte, 1)
	if Data {
		data[0] = 1
	} else {
		data[0] = 0
	}
	Message = BytesCombine([]byte{Id}, []byte{0x05}, Uint16ToByte(RegisterAddress), data)
	return
}

// BuildWriteSingleHoldingRegister 构建写单个保持寄存器报文
func BuildWriteSingleHoldingRegister(Id uint8, RegisterAddress uint16, Data []byte) (Message []byte) {
	Message = BytesCombine([]byte{Id}, []byte{0x06}, Uint16ToByte(RegisterAddress), []byte{1, 2}, Data)
	return
}

// BuildWriteMultipleCoilStatus 构建写多个线圈报文
func BuildWriteMultipleCoilStatus(Id uint8, RegisterAddress uint16, Data []bool) (Message []byte) {
	lenF := uint16(math.Ceil(float64(len(Data)) / 8))
	fmt.Println(lenF)
	var data = make([]byte, lenF)
	fmt.Println(data)
	for i := 0; i < int(lenF); i++ {
		for j := 0; j < len(Data)-i*8; j++ {
			if Data[i*8+j] {
				data[i] = data[i] + uint8(1<<j)
			}
		}
	}
	Message = BytesCombine([]byte{Id}, []byte{0x0f}, Uint16ToByte(RegisterAddress), Uint16ToByte(uint16(len(Data))), []byte{uint8(lenF)}, data)
	return
}

// BuildWriteMultipleHoldingRegister 构建写多个保持寄存器报文
func BuildWriteMultipleHoldingRegister(Id uint8, RegisterAddress uint16, Data []byte) (Message []byte) {
	Message = BytesCombine([]byte{Id}, []byte{0x10}, Uint16ToByte(RegisterAddress), Uint16ToByte(uint16(len(Data)/2)), []byte{uint8(len(Data))}, Data)
	return
}

package Comm

import "math"

// BuildReadCoilStatus  构建读取线圈状态报文
func BuildReadCoilStatus(Id uint8, RegisterAddress, Length uint16) (Message []byte) {
	Message = BytesCombine([]byte{Id}, []byte[]{0x01}, Uint16ToByte(RegisterAddress), Uint16ToByte(Length))
	return
}

// BuildReadInputStatus 构建读取输入状态报文
func BuildReadInputStatus(Id uint8, RegisterAddress, Length uint16) (Message []byte) {
	Message = BytesCombine([]byte{Id}, []byte[]{0x02}, Uint16ToByte(RegisterAddress), Uint16ToByte(Length))
	return
}

// BuildReadHoldingRegister 构建读取保持寄存器报文
func BuildReadHoldingRegister(Id uint8, RegisterAddress, Length uint16) (Message []byte) {
	Message = BytesCombine([]byte{Id}, []byte[]{0x03}, Uint16ToByte(RegisterAddress), Uint16ToByte(Length))
	return nil
}

// BuildReadInputRegister 构建读取输入寄存器报文
func BuildReadInputRegister(Id uint8, RegisterAddress, Length uint16) (Message []byte) {
	Message = BytesCombine([]byte{Id}, []byte[]{0x04}, Uint16ToByte(RegisterAddress), Uint16ToByte(Length))
	return nil
}

// BuildWriteSingleCoilStatus 构建写单个线圈报文
func BuildWriteSingleCoilStatus(Id uint8, RegisterAddress uint16, Data bool) (Message []byte) {
	Message = BytesCombine([]byte{Id}, []byte[]{0x05}, Uint16ToByte(RegisterAddress))
	return nil
}

// BuildWriteSingleHoldingRegister 构建写单个保持寄存器报文
func BuildWriteSingleHoldingRegister(Id uint8, RegisterAddress, Length uint16, Data Register) (Message []byte) {
	Message = BytesCombine([]byte{Id}, []byte[]{0x10}, Uint16ToByte(RegisterAddress), []byte{1, 2})
	return nil
}

// BuildWriteMultipleCoilStatus 构建写多个线圈报文
func BuildWriteMultipleCoilStatus(Id uint8, RegisterAddress uint16, Data []bool) (Message []byte) {
	lenf := math.Ceil(float64(len(Data) / 8))
	Message = BytesCombine([]byte{Id}, []byte[]{0x0f}, Uint16ToByte(RegisterAddress), Uint16ToByte(uint16(len(Data))), Uint16ToByte(uint16(lenf)))
	return nil
}

// BuildWriteMultipleHoldingRegister 构建写多个保持寄存器报文
func BuildWriteMultipleHoldingRegister(Id uint8, RegisterAddress uint16, Data Registers) (Message []byte) {

	Message = BytesCombine([]byte{Id}, []byte[]{0x10}, Uint16ToByte(RegisterAddress), Uint16ToByte(Data.Len()), Uint16ToByte(Data.Len()*2))
	return nil
}

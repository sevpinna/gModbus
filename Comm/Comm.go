package Comm

type Option struct {
	Ip            string
	AutoReconnect bool
	Serial        string
	Baud          int
	Parity        byte
	StopBits      byte
	Timeout       int
}
type ModbusMaster interface {
	Open()
	// ReadCoilStatus 读取线圈状态
	ReadCoilStatus(Id uint8, RegisterAddress, Length uint16) (Data []bool)
	// ReadInputStatus 读取输入状态
	ReadInputStatus(Id uint8, RegisterAddress, Length uint16) (Data []bool)
	// ReadHoldingRegister 读取保持寄存器
	ReadHoldingRegister(Id uint8, RegisterAddress, Length uint16) (Data []byte)
	// ReadInputRegister 读取输入寄存器
	ReadInputRegister(Id uint8, RegisterAddress, Length uint16) (Data []byte)
	// WriteSingleCoilStatus 写单个线圈
	WriteSingleCoilStatus(Id uint8, RegisterAddress uint16, Data bool) (err error)
	// WriteMultipleCoilStatus 写多个线圈
	WriteMultipleCoilStatus(Id uint8, RegisterAddress uint16, Data []bool) (err error)
	// WriteSingleHoldingRegister 写单个保持寄存器
	WriteSingleHoldingRegister(Id uint8, RegisterAddress uint16, Data []byte) (err error)
	// WriteMultipleHoldingRegister 写多个保持寄存器
	WriteMultipleHoldingRegister(Id uint8, RegisterAddress uint16, Data []byte) (err error)
}
type ModbusSlave interface {
}

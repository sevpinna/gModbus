package ModbusRtu

import (
	"gitee.com/sevpinna/gModbus/Comm"
	"gitee.com/sevpinna/gModbus/Comm/master"
	"github.com/tarm/serial"
	"log"
	"time"
)

type ModbusRtu struct {
	s           *serial.Port
	Name        string
	Baud        int
	Parity      serial.Parity
	StopBits    serial.StopBits
	ReadTimeout time.Duration
}

func (m *ModbusRtu) Open() {
	c := &serial.Config{Name: m.Name, Baud: m.Baud, Parity: m.Parity, StopBits: m.StopBits, ReadTimeout: m.ReadTimeout}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	m.s = s
}

// ReadCoilStatus 读取线圈状态
func (m *ModbusRtu) ReadCoilStatus(Id uint8, RegisterAddress, Length uint16) (Data []bool, err error) {
	msg := master.BuildReadCoilStatus(Id, RegisterAddress, Length)
	buff := Comm.BytesCombine(msg, Comm.CRC16CheckSum(msg))

	_, err = m.s.Write(buff)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 256)
	_, err = m.s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	Data = make([]bool, Length)
	for i := 0; i < int(Length); i++ {
		if (buf[i/8] & uint8(1<<i)) == uint8(1<<i) {
			Data[i] = true
		}
	}
	return
}

// ReadInputStatus 读取输入状态
func (m *ModbusRtu) ReadInputStatus(Id uint8, RegisterAddress, Length uint16) (Data []bool, err error) {
	return nil, err
}

// ReadHoldingRegister 读取保持寄存器
func (m *ModbusRtu) ReadHoldingRegister(Id uint8, RegisterAddress, Length uint16) (Data []byte, err error) {
	return nil, err
}

// ReadInputRegister 读取输入寄存器
func (m *ModbusRtu) ReadInputRegister(Id uint8, RegisterAddress, Length uint16) (Data []byte, err error) {
	return nil, err
}

// WriteSingleCoilStatus 写单个线圈
func (m *ModbusRtu) WriteSingleCoilStatus(Id uint8, RegisterAddress uint16, Data bool) (err error) {
	return err
}

// WriteMultipleCoilStatus 写多个线圈
func (m *ModbusRtu) WriteMultipleCoilStatus(Id uint8, RegisterAddress uint16, Data []bool) (err error) {
	return err
}

// WriteSingleHoldingRegister 写单个保持寄存器
func (m *ModbusRtu) WriteSingleHoldingRegister(Id uint8, RegisterAddress uint16, Data []byte) (err error) {
	return err
}

// WriteMultipleHoldingRegister 写多个保持寄存器
func (m *ModbusRtu) WriteMultipleHoldingRegister(Id uint8, RegisterAddress uint16, Data []byte) (err error) {
	return err
}

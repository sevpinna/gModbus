package ModbusTcp

import (
	"errors"
	"gitee.com/sevpinna/gModbus/Comm"
	"gitee.com/sevpinna/gModbus/Comm/master"
	"math"
	"net"
	"time"
)

type ModbusTcp struct {
	conn net.Conn
	//IP地址加端口
	Ip string
	//是否启用自动重连
	AutoReConnect bool
	//超时
	Timeout int
}

// Open 创建TCP链接
func (m *ModbusTcp) Open() {
	for {
		conn, err := net.Dial("tcp", m.Ip)
		if err != nil {
			continue
		}
		m.conn = conn
		if m.AutoReConnect {
			go m.autoReConnect()
		}
		break
	}
}

// 自动重连函数
func (m *ModbusTcp) autoReConnect() {
	for {
		if m.conn == nil {
			for {
				conn, err := net.Dial("tcp", m.Ip)
				if err != nil {
					time.Sleep(100 * time.Millisecond)
					continue
				}
				m.conn = conn
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// 发送并接收网络数据
func (m *ModbusTcp) sendData(sendBuff []byte) (revBuff []byte, err error) {
	var n int
	err = m.conn.SetDeadline(time.Now().Add(time.Duration(m.Timeout) * time.Millisecond))
	_, err = m.conn.Write(sendBuff)
	if err != nil {
		m.conn = nil
	}

	revBuff = make([]byte, 256)
	err = m.conn.SetDeadline(time.Now().Add(time.Duration(m.Timeout) * time.Millisecond))
	n, err = m.conn.Read(revBuff)
	if err != nil {
		m.conn = nil
	}
	return revBuff[:n], err
}

// ReadCoilStatus 读取线圈状态
func (m *ModbusTcp) ReadCoilStatus(Id uint8, RegisterAddress, Length uint16) (Data []bool, err error) {
	msg := master.BuildReadCoilStatus(Id, RegisterAddress, Length)
	buff := Comm.BytesCombine([]byte{0x00, 0x00, 0x00, 0x00}, Comm.Uint16ToByte(uint16(len(msg))), msg)

	readBuff, errA := m.sendData(buff)
	f := readBuff[9 : 9+readBuff[8]]

	Data = make([]bool, Length)

	for i := 0; i < int(Length); i++ {
		j := int(math.Floor(float64(i) / 8))
		if (f[i/8] & uint8(1<<(i-8*j))) == uint8(1<<(i-8*j)) {
			Data[i] = true
		}

	}
	return Data, errA
}

// ReadInputStatus 读取输入状态
func (m *ModbusTcp) ReadInputStatus(Id uint8, RegisterAddress, Length uint16) (Data []bool, err error) {
	msg := master.BuildReadCoilStatus(Id, RegisterAddress, Length)
	buff := Comm.BytesCombine([]byte{0x00, 0x00, 0x00, 0x00}, Comm.Uint16ToByte(uint16(len(msg))), msg)
	readBuff, errA := m.sendData(buff)
	f := readBuff[9 : 9+readBuff[8]]

	Data = make([]bool, Length)

	for i := 0; i < int(Length); i++ {
		j := int(math.Floor(float64(i) / 8))
		if (f[i/8] & uint8(1<<(i-8*j))) == uint8(1<<(i-8*j)) {
			Data[i] = true
		}

	}
	return Data, errA
}

// ReadHoldingRegister 读取保持寄存器
func (m *ModbusTcp) ReadHoldingRegister(Id uint8, RegisterAddress, Length uint16) (Data []byte, err error) {
	msg := master.BuildReadHoldingRegister(Id, RegisterAddress, Length)
	buff := Comm.BytesCombine([]byte{0x00, 0x00, 0x00, 0x00}, Comm.Uint16ToByte(uint16(len(msg))), msg)

	readBuff, errA := m.sendData(buff)
	f := readBuff[9 : 9+readBuff[8]]
	return f, errA
}

// ReadInputRegister 读取输入寄存器
func (m *ModbusTcp) ReadInputRegister(Id uint8, RegisterAddress, Length uint16) (Data []byte, err error) {
	msg := master.BuildReadInputRegister(Id, RegisterAddress, Length)
	buff := Comm.BytesCombine([]byte{0x00, 0x00, 0x00, 0x00}, Comm.Uint16ToByte(uint16(len(msg))), msg)

	readBuff, errA := m.sendData(buff)
	f := readBuff[9 : 9+readBuff[8]]
	return f, errA
}

// WriteSingleCoilStatus 写单个线圈
func (m *ModbusTcp) WriteSingleCoilStatus(Id uint8, RegisterAddress uint16, Data bool) (err error) {
	msg := master.BuildWriteSingleCoilStatus(Id, RegisterAddress, Data)
	buff := Comm.BytesCombine([]byte{0x00, 0x00, 0x00, 0x00}, Comm.Uint16ToByte(uint16(len(msg))), msg)
	readBuff, errA := m.sendData(buff)
	if len(readBuff) != len(buff) {
		return errors.New("error")
	}
	return errA
}

// WriteMultipleCoilStatus 写多个线圈
func (m *ModbusTcp) WriteMultipleCoilStatus(Id uint8, RegisterAddress uint16, Data []bool) (err error) {
	msg := master.BuildWriteMultipleCoilStatus(Id, RegisterAddress, Data)
	buff := Comm.BytesCombine([]byte{0x00, 0x00, 0x00, 0x00}, Comm.Uint16ToByte(uint16(len(msg))), msg)
	readBuff, errA := m.sendData(buff)
	if len(readBuff) != len(buff) {
		return errors.New("error")
	}
	return errA
}

// WriteSingleHoldingRegister 写单个保持寄存器
func (m *ModbusTcp) WriteSingleHoldingRegister(Id uint8, RegisterAddress uint16, Data []byte) (err error) {
	msg := master.BuildWriteSingleHoldingRegister(Id, RegisterAddress, Data)
	buff := Comm.BytesCombine([]byte{0x00, 0x00, 0x00, 0x00}, Comm.Uint16ToByte(uint16(len(msg))), msg)

	readBuff, errA := m.sendData(buff)

	if len(readBuff) != len(buff) {
		return errors.New("error")
	}
	return errA
}

// WriteMultipleHoldingRegister 写多个保持寄存器
func (m *ModbusTcp) WriteMultipleHoldingRegister(Id uint8, RegisterAddress uint16, Data []byte) (err error) {
	msg := master.BuildWriteMultipleHoldingRegister(Id, RegisterAddress, Data)
	buff := Comm.BytesCombine([]byte{0x00, 0x00, 0x00, 0x00}, Comm.Uint16ToByte(uint16(len(msg))), msg)
	readBuff, errA := m.sendData(buff)
	if len(readBuff) == len(buff) {
		return errors.New("error")
	}
	return errA
}

package ModbusTcp

import (
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
}

// Connect 创建TCP链接
func (m *ModbusTcp) Connect() {
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
					time.Sleep(100 * time.Microsecond)
					continue
				}
				m.conn = conn
			}
		}
		time.Sleep(100 * time.Microsecond)
	}
}

// ReadCoilStatus 读取线圈状态
func (m *ModbusTcp) ReadCoilStatus(Id uint8, RegisterAddress, Length uint16) (Data []bool, err error) {
	msg := master.BuildReadCoilStatus(Id, RegisterAddress, Length)
	buff := Comm.BytesCombine([]byte{0x00, 0x00, 0x00, 0x00}, Comm.Uint16ToByte(uint16(len(msg))), msg)

	m.conn.SetDeadline(time.Now().Add(1000 * time.Microsecond))
	_, err = m.conn.Write(buff)
	if err != nil {
		return nil, err
	}
	var readBuff = make([]byte, 256)
	m.conn.SetDeadline(time.Now().Add(1000 * time.Microsecond))
	_, readErr := m.conn.Read(readBuff)
	if readErr != nil {
		return nil, readErr
	}
	f := readBuff[9 : 9+readBuff[8]]

	Data = make([]bool, Length)

	for i := 0; i < int(Length); i++ {
		j := int(math.Floor(float64(i) / 8))
		if (f[i/8] & uint8(1<<(i-8*j))) == uint8(1<<(i-8*j)) {
			Data[i] = true
		}

	}
	return Data, nil
}

// ReadInputStatus 读取输入状态
func (m *ModbusTcp) ReadInputStatus(Id uint8, RegisterAddress, Length uint16) (Data []bool, err error) {
	return nil, err
}

// ReadHoldingRegister 读取保持寄存器
func (m *ModbusTcp) ReadHoldingRegister(Id uint8, RegisterAddress, Length uint16) (Data []byte, err error) {
	return nil, err
}

// ReadInputRegister 读取输入寄存器
func (m *ModbusTcp) ReadInputRegister(Id uint8, RegisterAddress, Length uint16) (Data []byte, err error) {
	return nil, err
}

// WriteSingleCoilStatus 写单个线圈
func (m *ModbusTcp) WriteSingleCoilStatus(Id uint8, RegisterAddress uint16, Data bool) (err error) {
	return err
}

// WriteMultipleCoilStatus 写多个线圈
func (m *ModbusTcp) WriteMultipleCoilStatus(Id uint8, RegisterAddress uint16, Data []bool) (err error) {
	return err
}

// WriteSingleHoldingRegister 写单个保持寄存器
func (m *ModbusTcp) WriteSingleHoldingRegister(Id uint8, RegisterAddress, Length uint16, Data []byte) (err error) {
	return err
}

// WriteMultipleHoldingRegister 写多个保持寄存器
func (m *ModbusTcp) WriteMultipleHoldingRegister(Id uint8, RegisterAddress, Length uint16, Data []byte) (err error) {
	return err
}

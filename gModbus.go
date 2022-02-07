package gModbus

import (
	"gitee.com/sevpinna/gModbus/Comm"
	"gitee.com/sevpinna/gModbus/ModbusRtu"
	"gitee.com/sevpinna/gModbus/ModbusTcp"
	"github.com/tarm/serial"
	"time"
)

type Option struct {
	Ip            string
	AutoReconnect bool
	Serial        string
	Baud          int
	Parity        byte
	StopBits      byte
	Timeout       int
}

func CreateModbusTcpClient(ip string, autoReconnect bool, Timeout int) (modbusTcp *ModbusTcp.ModbusTcp) {
	return &ModbusTcp.ModbusTcp{Ip: ip, AutoReConnect: autoReconnect, Timeout: Timeout}
}
func CreateModbusRtuClient(name string, baud int, parity byte, stopBits byte, readTimeout int) (modbusRtu *ModbusRtu.ModbusRtu) {
	return &ModbusRtu.ModbusRtu{Name: name, Baud: baud, Parity: serial.Parity(parity), StopBits: serial.StopBits(stopBits),
		ReadTimeout: time.Duration(readTimeout) * time.Microsecond}
}
func CreateModbusClient(Type string, option Option) (modbus Comm.ModbusMaster) {
	switch Type {
	case "tcp":
		modbus = CreateModbusTcpClient(option.Ip, option.AutoReconnect, option.Timeout) //&ModbusTcp.ModbusTcp{Ip: option.Ip, AutoReConnect: option.AutoReconnect}
	case "rtu":
		modbus = CreateModbusRtuClient(option.Serial, option.Baud, option.Parity, option.StopBits, option.Timeout)
		//&ModbusRtu.ModbusRtu{Name: option.Serial, Baud: option.Baud, Parity: serial.Parity(option.Parity), StopBits: serial.StopBits(option.StopBits),
		//ReadTimeout: time.Duration(option.ReadTimeout) * time.Microsecond}
	}
	return modbus
}

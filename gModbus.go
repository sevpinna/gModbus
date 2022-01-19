package gModbus

import (
	"gitee.com/sevpinna/gModbus/ModbusRtu"
	"gitee.com/sevpinna/gModbus/ModbusTcp"
	"github.com/tarm/serial"
	"time"
)

func CreateModbusTcpClient(ip string, autoReconnect bool) (modbusTcp *ModbusTcp.ModbusTcp) {
	return &ModbusTcp.ModbusTcp{Ip: ip, AutoReConnect: autoReconnect}
}
func CreateModbusRtuClient(name string, baud int, parity byte, stopBits byte, readTimeout int) (modbusRtu *ModbusRtu.ModbusRtu) {
	return &ModbusRtu.ModbusRtu{Name: name, Baud: baud, Parity: serial.Parity(parity), StopBits: serial.StopBits(stopBits),
		ReadTimeout: time.Duration(readTimeout) * time.Microsecond}
}

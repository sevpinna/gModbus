package gModbus

import (
	"fmt"
	"gitee.com/sevpinna/gModbus/Comm"
	"testing"
)

func TestQ(t *testing.T) {
	g := CreateModbusClient("tcp", Comm.Option{Ip: "127.0.0.1:502", AutoReconnect: true, Timeout: 1000})
	//g := CreateModbusTcpClient("10.254.2.53:502", true,1000)
	g.Open()
	//h:=g.ReadCoilStatus(1,0,1)
	//gg, err := g.WriteSingleCoilStatus(1, 0, false)
	err := g.WriteMultipleHoldingRegister(1, 0, []byte{0x00, 0x05, 0xff, 0xff})
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println("OK")
	////f := Comm.BuildReadCoilStatus(1, 0, 9)
	//f2 := master.BuildWriteMultipleHoldingRegister(1, 0, []byte{0, 111})
	////f2 := Comm.BuildWriteMultipleCoilStatus(1, 0, []bool{false, true, false, true, true, false, true, false, true, true})
	//fmt.Printf("%+v\n", f2)
	//f3 := Comm.BytesCombine([]byte{0x00, 0x00, 0x00, 0x00}, Comm.Uint16ToByte(uint16(len(f2))), f2)
	//conn, _ := net.Dial("tcp", "10.254.2.53:502")
	//conn.Write(f3)
	//g := make([]byte, 30)
	//conn.Read(g)
	//fmt.Printf("%+v\n", g)
}

package gModbus

import (
	"fmt"
	"gitee.com/sevpinna/gModbus/Comm"
	"net"
	"testing"
)

func TestQ(t *testing.T) {
	//f := Comm.BuildReadCoilStatus(1, 0, 9)
	f2 := Comm.BuildWriteMultipleHoldingRegister(1, 0, []byte{0, 111})
	//f2 := Comm.BuildWriteMultipleCoilStatus(1, 0, []bool{false, true, false, true, true, false, true, false, true, true})
	fmt.Printf("%+v\n", f2)
	f3 := Comm.BytesCombine([]byte{0x00, 0x00, 0x00, 0x00}, Comm.Uint16ToByte(uint16(len(f2))), f2)
	conn, _ := net.Dial("tcp", "10.254.2.53:502")
	conn.Write(f3)
	g := make([]byte, 30)
	conn.Read(g)
	fmt.Printf("%+v\n", g)
}
package at

import (
	"fmt"
	"unsafe"
)

type BaudRate int

const (
	BaudRate115200 BaudRate = 115200
	BaudRate57600  BaudRate = 57600
	BaudRate38400  BaudRate = 38400
	BaudRate19200  BaudRate = 19200
	BaudRate14400  BaudRate = 14400
	BaudRate9600   BaudRate = 9600
	BaudRate4800   BaudRate = 4800
	BaudRate2400   BaudRate = 2400
	BaudRate1200   BaudRate = 1200
)

type DataBits int

const (
	// DataBits7 7位数据位 暂不支持
	DataBits7 DataBits = iota
	// DataBits8 8位数据位
	DataBits8
)

type StopBits int

const (
	// StopBits1 1位停止位
	StopBits1 StopBits = iota
	// StopBits2 2位停止位
	StopBits2
)

type Parity int

const (
	// ParityNone 无校验
	ParityNone Parity = iota
	// ParityOdd 奇校验
	ParityOdd
	// ParityEven 偶校验
	ParityEven
)

// SetSerialCommand 串口参数设置
type SetSerialCommand interface {
	Commander
	SetBaudRate(baudRate BaudRate) SetSerialCommand
	SetDataBits(dataBits DataBits) SetSerialCommand
	SetStopBits(stopBits StopBits) SetSerialCommand
	SetParity(parity Parity) SetSerialCommand
}

func NewSetSerialCommand() SetSerialCommand {
	return &setSerial{baseCommand: &baseCommand{Value: "@DTU:0000:UARTCFG=%d,%d,%d,%d"}}
}

type setSerial struct {
	*baseCommand
	baudRate BaudRate
	dataBits DataBits
	stopBits StopBits
	parity   Parity
}

//func (s *setSerial) SetType(commandType CommandType) error {
//	switch commandType {
//	case CommandTypeSet:
//		s.Value += "UARTCFG=%d,%d,%d,%d"
//	case CommandTypeGet:
//		s.Value += "UARTCFG?"
//	}
//	return nil
//}

func (s *setSerial) SetBaudRate(baudRate BaudRate) SetSerialCommand {
	s.baudRate = baudRate
	return s
}

func (s *setSerial) SetDataBits(dataBits DataBits) SetSerialCommand {
	s.dataBits = dataBits
	return s
}

func (s *setSerial) SetStopBits(stopBits StopBits) SetSerialCommand {
	s.stopBits = stopBits
	return s
}

func (s *setSerial) SetParity(parity Parity) SetSerialCommand {
	s.parity = parity
	return s
}

func (s *setSerial) String() string {
	return fmt.Sprintf(s.Value, s.baudRate, s.dataBits, s.parity, s.stopBits)
}

func (s *setSerial) Bytes() []byte {
	s1 := s.String()
	return unsafe.Slice(unsafe.StringData(s1), len(s1))
}

var ReadSerialCommand = &readSerial{baseCommand: &baseCommand{Value: "@DTU:0000:UARTCFG?"}}

type readSerial struct {
	*baseCommand
}

func (r *readSerial) String() string {
	return r.Value
}

func (r *readSerial) Bytes() []byte {
	s := r.String()
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

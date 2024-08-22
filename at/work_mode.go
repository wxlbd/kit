package at

import (
	"fmt"
	"unsafe"
)

// ReadWorkModeCommand 读取DTU工作模式
var ReadWorkModeCommand = &readworkMode{baseCommand: &baseCommand{Value: "@DTU:0000:DTUMODE?"}}

// SetWorkModeCommand DTU通道工作模式配置
type SetWorkModeCommand interface {
	Commander
	SetMode(mode DTUWorkMode) SetWorkModeCommand
	SetChannelNumber(channelNumber int) SetWorkModeCommand
}

func NewSetWorkModeCommand() SetWorkModeCommand {
	return &setWorkMode{baseCommand: &baseCommand{Value: "@DTU:0000:DTUMODE=%d,%d"}, channelNumber: 1}
}

type DTUWorkMode int

const (
	DTUWorkModeNone DTUWorkMode = iota
	DTUWorkModeTCPUDP
	DTUWorkModeMQTT
	_
	_
	DTUWorkModeHTTP
	DTUWorkModeAliyun
)

type setWorkMode struct {
	*baseCommand
	mode          DTUWorkMode
	channelNumber int
}

// SetMode 设置工作模式 取值范围 0~6，
//
// 0 不启用该通道；
//
// 1 TCP/UDP 透传；
//
// 2 MQTT 透传；
//
// 3 塔石透传云连接；
//
// 4 塔石 IOT 云连接；
//
// 5 HTTP 透传模式；
//
// 6 阿里云直连；
func (w *setWorkMode) SetMode(mode DTUWorkMode) SetWorkModeCommand {
	w.mode = mode
	return w
}

func (w *setWorkMode) SetChannelNumber(channelNumber int) SetWorkModeCommand {
	w.channelNumber = channelNumber
	return w
}

func (w *setWorkMode) String() string {
	return fmt.Sprintf(w.Value, w.mode, w.channelNumber)
}

func (w *setWorkMode) Bytes() []byte {
	s := w.String()
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

type readworkMode struct {
	*baseCommand
}

package at

import (
	"bytes"
	"unsafe"
)

type CommandType int

const (
	// CommandTypeSet 远程设置
	CommandTypeSet CommandType = iota
	// CommandTypeGet 远程获取
	CommandTypeGet
)

// Commander 定义了AT指令执行器的接口
type Commander interface {
	// String 转换为命令字符串
	String() string
	Bytes() []byte
	// ParseResponse 解析响应
	ParseResponse([]byte) ([]byte, error)
}

type baseCommand struct {
	Value string
}

func (i *baseCommand) String() string {
	return i.Value
}

func (i *baseCommand) Bytes() []byte {
	s := i.String()
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func (i *baseCommand) ParseResponse(bs []byte) ([]byte, error) {
	ok := []byte("OK")
	if !bytes.Contains(bs, ok) {
		return nil, ErrOperationFailed
	}
	return ok, nil
}

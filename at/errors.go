package at

import "fmt"

var (
	// ErrUnsupportedCommandType 该指令不支持此命令类型
	ErrUnsupportedCommandType = fmt.Errorf("该指令不支持此命令类型")
	// ErrInvalidWorkMode 无效的工作模式
	ErrInvalidWorkMode = fmt.Errorf("无效的工作模式")
	// ErrOperationFailed 操作失败
	ErrOperationFailed = fmt.Errorf("操作失败")
)

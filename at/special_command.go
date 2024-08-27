package at

// SaveConfigCommand  保存参数
var (
	SaveConfigCommand           = &saveConfig{baseCommand: &baseCommand{Value: "@DTU:0000:AT&W"}}
	_                 Commander = (*saveConfig)(nil)
)

// saveConfig 保存配置
type saveConfig struct {
	*baseCommand
}

// PowerOffCommand 重启设备
var PowerOffCommand = &powerOff{baseCommand: &baseCommand{Value: "@DTU:0000:POWEROFF"}}

type powerOff struct {
	*baseCommand
}

func (s *powerOff) SetType(t CommandType) error {
	return nil
}

var EnterConfigurationMode = &enterConfigurationMode{baseCommand: &baseCommand{Value: "@DTU:0000:+++"}}

type enterConfigurationMode struct {
	*baseCommand
}

// ReadSignal 读取信号
var ReadSignal = &readSignal{baseCommand: &baseCommand{Value: "@DTU:0000:CSQ"}}

// 读取信号值
type readSignal struct {
	*baseCommand
}

func (s *readSignal) String() string {
	return s.Value
}

// ReadICCIDCommand  获取卡号
var ReadICCIDCommand = &iccid{baseCommand: &baseCommand{Value: "@DTU:0000:ICCID"}}

type iccid struct {
	*baseCommand
}

// ReadIMEICommand  获取IMEI
var ReadIMEICommand = &readIMEI{baseCommand: &baseCommand{Value: "@DTU:0000:GSN?"}}

type readIMEI struct {
	*baseCommand
}

func (s *readIMEI) SetType(t CommandType) error {
	return nil
}

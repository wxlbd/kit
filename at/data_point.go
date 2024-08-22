package at

import (
	"fmt"
	"unsafe"
)

// CollectType 自定义轮询使能
//
// 0 关闭数据采集功能
//
// 1 开启字串采集，透明上报
//
// 2 寄存器采集，JSON 上报
type CollectType int

const (
	CollectTypeOff  CollectType = iota
	CollectTypeStr              // 字串采集
	CollectTypeJSON             // 寄存器采集，JSON 上报
)

// DataFormat 数据输入格式
type DataFormat int

const (
	// DataFormatAscii ASCII 格式，设置为 0 表示之后输入的轮询指令均为以 ASCII 形式轮 即输入什么字串就轮询什么字串
	DataFormatAscii DataFormat = iota
	// DataFormatHex HEX 格式，设置为 1 表示之后输入的轮询指令需要满足 HEX 格式，轮 询时会自动转成 16 进制对应的 ASCII 字串
	DataFormatHex
)

type SetPollCommand interface {
	Commander
	SetCollectType(collectType CollectType) SetPollCommand
	SetPollInterval(pollInterval int) SetPollCommand
	SetDataFormat(dataFormat DataFormat) SetPollCommand
}

type setPollCommand struct {
	*baseCommand
	collectType  CollectType
	pollInterval int
	dataFormat   DataFormat
}

// SetCollectType 自定义轮询使能，取值范围 0-1

func (p *setPollCommand) SetCollectType(collectType CollectType) SetPollCommand {
	p.collectType = collectType
	return p
}

// SetPollInterval 设置轮询间隔，单位秒,取值范围1-65535
func (p *setPollCommand) SetPollInterval(pollInterval int) SetPollCommand {
	p.pollInterval = pollInterval
	return p
}

// SetDataFormat 轮询数据输入格式，取值范围 0-1
//
// 0 ASCII 模式
//
// 1 HEX 格式
func (p *setPollCommand) SetDataFormat(dataFormat DataFormat) SetPollCommand {
	p.dataFormat = dataFormat
	return p
}

func (p *setPollCommand) String() string {
	return fmt.Sprintf(p.Value, p.collectType, p.pollInterval, p.dataFormat)
}

func (p *setPollCommand) Bytes() []byte {
	s := p.String()
	fmt.Println(s)
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func NewSetPollCommand() SetPollCommand {
	return &setPollCommand{
		baseCommand: &baseCommand{Value: "@DTU:0000:POLL=%d,%d,%d"},
	}
}

// ReadPollCommand 读取轮询配置
var ReadPollCommand = &readPoll{baseCommand: &baseCommand{Value: "@DTU:0000:POLL?"}}

type readPoll struct {
	*baseCommand
}

// NewSetPollStrCommand 轮询字串设置
func NewSetPollStrCommand() SetPollStrCommand {
	return &setPollStr{baseCommand: &baseCommand{Value: `@DTU:0000:POLLSTR=%d,%d,%d,"%s"`}}
}

type SetPollStrCommand interface {
	Commander
	SetIndex(index int) SetPollStrCommand
	SetEnable(enable int) SetPollStrCommand
	SetCRCEnable(crcEnable int) SetPollStrCommand
	SetStr(str string) SetPollStrCommand
}

type setPollStr struct {
	*baseCommand
	index     int
	enable    int
	crcEnable int
	str       string
}

// SetIndex  轮询字串号，取值范围 1-20
func (p *setPollStr) SetIndex(index int) SetPollStrCommand {
	p.index = index
	return p
}

// SetEnable 字串轮询使能，取值范围 0-1
//
// 0 禁用该条轮询
//
// 1 启用该条轮询
func (p *setPollStr) SetEnable(enable int) SetPollStrCommand {
	p.enable = enable
	return p
}

// SetCRCEnable CRC 校验开关，取值范围 0-1
//
// 0 关闭 CRC 校验
//
// 1 开启 CRC 校验
func (p *setPollStr) SetCRCEnable(crcEnable int) SetPollStrCommand {
	p.crcEnable = crcEnable
	return p
}

// SetStr 轮询字串内容
func (p *setPollStr) SetStr(str string) SetPollStrCommand {
	p.str = str
	return p
}

func (p *setPollStr) String() string {
	return fmt.Sprintf(p.Value, p.index, p.enable, p.crcEnable, p.str)
}

func (p *setPollStr) Bytes() []byte {
	s := p.String()
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// ReadPollStrCommand 读取轮询字串配置
var ReadPollStrCommand = &readPollStr{baseCommand: &baseCommand{Value: "@DTU:0000:POLLSTR?"}}

type readPollStr struct {
	*baseCommand
	index int
}

// PollJsonCfgCommand 轮询JSON使能配置
type PollJsonCfgCommand interface {
	Commander
	SetIndex(index int) PollJsonCfgCommand
	SetEnable(enable int) PollJsonCfgCommand
	SetKey(key string) PollJsonCfgCommand
	// SetSource 数据来源，取值范围 0-1
	//
	//0 Modbus RTU，可以根据 Modbus 协议进行数据解析
	//
	//1 自定义
	SetSource(source int) PollJsonCfgCommand
}

func NewPollJsonCfgCommand() PollJsonCfgCommand {
	return &pollJsonEnableConfig{baseCommand: &baseCommand{Value: `@DTU:0000:JSONCFG=%d,%d,"%s",%d,%d`}}
}

type pollJsonEnableConfig struct {
	*baseCommand
	index  int
	enable int
	key    string
	source int
}

func (p *pollJsonEnableConfig) SetIndex(index int) PollJsonCfgCommand {
	p.index = index
	return p
}

func (p *pollJsonEnableConfig) SetEnable(enable int) PollJsonCfgCommand {
	p.enable = enable
	return p
}

func (p *pollJsonEnableConfig) SetKey(key string) PollJsonCfgCommand {
	p.key = key
	return p
}

func (p *pollJsonEnableConfig) SetSource(source int) PollJsonCfgCommand {
	p.source = source
	return p
}

func (p *pollJsonEnableConfig) String() string {
	return fmt.Sprintf(p.Value, p.index, p.enable, p.key, p.source, 0)
}

func (p *pollJsonEnableConfig) Bytes() []byte {
	s := p.String()
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// PollJsonRegCfgCommand 轮询 JSON 寄存器配置
type PollJsonRegCfgCommand interface {
	Commander
	SetIndex(index int) PollJsonRegCfgCommand
	// SetSlave 从站地址，取值范围 1-255
	SetSlave(slave int) PollJsonRegCfgCommand
	// SetFuncCode 功能码，取值范围 1-4 1,2 开关量（暂无） 3,4 模拟量
	SetFuncCode(funcCode int) PollJsonRegCfgCommand
	// SetRegAddr 寄存器地址，取值范围 0-65535
	SetRegAddr(regAddr int) PollJsonRegCfgCommand
	// SetDataType 数据类型，取值范围 0-6
	//
	// 0 16 位 有符号数
	//
	// 1 16 位 无符号数
	//
	// 2 16 位 按位读取(读取数据的第一位数据是 0 还是 1)
	//
	// 3 32 位 有符号数
	//
	// 4 32 位 无符号数
	//
	// 5 32 位 按位读取
	//
	// 6 32 位 按位读取
	SetDataType(dataType RegResponseDataFormat) PollJsonRegCfgCommand
	// SetByteOrder 设置字节顺序，取值范围 0-4
	// 当 F 为 0、1、2 时(即采集数据为 16 位时)代表含义如下 0 AB 1 BA 其他数值无意义代表为 BA
	// 当 F 为 3、4、5(即采集数据为 32 位时)代表含义如下 0 ABCD 1 CDAB 2 BADC 3 DCBA
	SetByteOrder(byteOrder RegResponseDataByteOrder) PollJsonRegCfgCommand
	// SetReserve 设置保留小数位数 保留小数点位数，取值范围-9~9，正数小数点左移，负数小数点右移.
	SetReserve(reserve int) PollJsonRegCfgCommand
	// SetUnit 设置单位 0-10 个字节，如果没有填空””
	SetUnit(unit string) PollJsonRegCfgCommand
	// SetQuote 设置是否包含引号 0 不包含引号 1 包含引号
	SetQuote(quote int) PollJsonRegCfgCommand
	// SetInterval 串口轮询间隔，取值范围>500ms，轮询完该条指令后经过多久轮询下一条(单位毫秒)
	SetInterval(interval int) PollJsonRegCfgCommand
}

type pollJsonRegConfig struct {
	*baseCommand
	index     int
	slave     int
	funcCode  int
	regAddr   int
	dataType  RegResponseDataFormat
	byteOrder RegResponseDataByteOrder
	reserve   int
	unit      string
	quote     int
	interval  int
}

func (p *pollJsonRegConfig) SetIndex(index int) PollJsonRegCfgCommand {
	p.index = index
	return p
}

func (p *pollJsonRegConfig) SetSlave(slave int) PollJsonRegCfgCommand {
	p.slave = slave
	return p
}

func (p *pollJsonRegConfig) SetFuncCode(funcCode int) PollJsonRegCfgCommand {
	p.funcCode = funcCode
	return p
}

func (p *pollJsonRegConfig) SetRegAddr(regAddr int) PollJsonRegCfgCommand {
	p.regAddr = regAddr
	return p
}

func (p *pollJsonRegConfig) SetDataType(dataType RegResponseDataFormat) PollJsonRegCfgCommand {
	p.dataType = dataType
	return p
}

func (p *pollJsonRegConfig) SetByteOrder(byteOrder RegResponseDataByteOrder) PollJsonRegCfgCommand {
	p.byteOrder = byteOrder
	return p
}

func (p *pollJsonRegConfig) SetReserve(reserve int) PollJsonRegCfgCommand {
	p.reserve = reserve
	return p
}

func (p *pollJsonRegConfig) SetUnit(unit string) PollJsonRegCfgCommand {
	p.unit = unit
	return p
}

func (p *pollJsonRegConfig) SetQuote(quote int) PollJsonRegCfgCommand {
	p.quote = quote
	return p
}

func (p *pollJsonRegConfig) SetInterval(interval int) PollJsonRegCfgCommand {
	p.interval = interval
	return p
}

func (p *pollJsonRegConfig) String() string {
	return fmt.Sprintf(p.Value, p.index, p.slave, p.funcCode, p.regAddr, p.dataType, p.byteOrder, p.reserve, p.unit, p.quote, p.interval)
}

func (p *pollJsonRegConfig) Bytes() []byte {
	s := p.String()
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func NewPollJsonRegConfig() PollJsonRegCfgCommand {
	return &pollJsonRegConfig{
		baseCommand: &baseCommand{Value: "@DTU:0000:REGCFG=%d,%d,%d,%d,0,%d,0,%d,%d,”%s”,%d,%d"},
		interval:    500,
	}
}

// RegResponseDataFormat 寄存器配置响应数据格式
type RegResponseDataFormat int

const (
	// RegResponseDataFormatSigned16 16位有符号数
	RegResponseDataFormatSigned16 RegResponseDataFormat = iota
	// RegResponseDataFormatUnsigned16 16位无符号数
	RegResponseDataFormatUnsigned16
	// RegResponseDataFormatBit16 16 位 按位读取(读取数据的第一位数据是 0 还是 1)
	RegResponseDataFormatBit16
	// RegResponseDataFormatSigned32 32位有符号数
	RegResponseDataFormatSigned32
	// RegResponseDataFormatUnsigned32 32位无符号数
	RegResponseDataFormatUnsigned32
	// RegResponseDataFormatFloat32 32位浮点数
	RegResponseDataFormatFloat32
)

type RegResponseDataByteOrder int

const (
	// RegResponseDataByteOrderBigEndian 大端
	RegResponseDataByteOrderBigEndian RegResponseDataByteOrder = iota
	// RegResponseDataByteOrderLittleEndian 小端
	RegResponseDataByteOrderLittleEndian
	RegResponseDataByteOrderBADC
	RegResponseDataByteOrderDCBA
)

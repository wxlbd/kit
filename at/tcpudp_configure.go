package at

import (
	"fmt"
	"unsafe"
)

const (
	TransferTCP = "TCP"
	TransferUDP = "UDP"
)

type TransferProto string

const (
	TransferProtoTCP TransferProto = "TCP"
	TransferProtoUDP TransferProto = "UDP"
)

type SetTCPUDPServerCfgCommand interface {
	SetTransferProto(transferMode TransferProto) SetTCPUDPServerCfgCommand
	SetChannelNumber(channelNumber int) SetTCPUDPServerCfgCommand
	SetAddress(address string) SetTCPUDPServerCfgCommand
	SetPort(port int) SetTCPUDPServerCfgCommand
}

func NewSetTCPUDPServerCfgCommand() SetTCPUDPServerCfgCommand {
	return &setTcpudpServer{baseCommand: &baseCommand{Value: `@DTU:0000:DSCADDR=%d,"%s","%s",%d`}, channelNumber: 1}
}

type setTcpudpServer struct {
	*baseCommand
	protocol      TransferProto
	channelNumber int
	address       string
	port          int
}

// SetTransferProto 设置传输模式 TCP/UDP
func (t *setTcpudpServer) SetTransferProto(transferMode TransferProto) SetTCPUDPServerCfgCommand {
	t.protocol = transferMode
	return t
}

func (t *setTcpudpServer) SetChannelNumber(channelNumber int) SetTCPUDPServerCfgCommand {
	t.channelNumber = channelNumber
	return t
}

// SetAddress 设置服务器地址
func (t *setTcpudpServer) SetAddress(address string) SetTCPUDPServerCfgCommand {
	t.address = address
	return t
}

// SetPort 设置服务器端口
func (t *setTcpudpServer) SetPort(port int) SetTCPUDPServerCfgCommand {
	t.port = port
	return t
}

func (t *setTcpudpServer) String() string {
	return fmt.Sprintf(t.Value, t.channelNumber, t.protocol, t.address, t.port)
}

func (t *setTcpudpServer) Bytes() []byte {
	s := t.String()
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

type RegisterMode int

const (
	RegisterModeDisable RegisterMode = iota
	// RegisterModeOnlyConnect 仅连接时上传
	RegisterModeOnlyConnect
	// RegisterModeData 和数据一起上传
	RegisterModeData
	// RegisterModeDataAndConnect 包括 1,2
	RegisterModeDataAndConnect
)

type RegisterContentType int

const (
	RegisterContentTypeCustom RegisterContentType = iota
	// RegisterContentTypeIMEI 15 位模块对应的唯一识别码
	RegisterContentTypeIMEI
	// RegisterContentTypeICCID 20 位 SIM 卡对应编码
	RegisterContentTypeICCID
)

type SetRegisterPacketCfgCommand interface {
	// SetRegisterMode 设置注册包模式
	//
	// 0 不启用注册包
	//
	// 1 仅连接时上传
	//
	// 2 和数据一起上传，在数据前
	//
	// 3 包括 1,2
	SetRegisterMode(registerMode RegisterMode) SetRegisterPacketCfgCommand
	// SetRegisterContentType 注册包内容，取值范围 0-2
	//
	// 0 自定义注册包
	//
	// 1 IMEI(15 位模块对应的唯一识别码)(选取该位后，格式为 ASCII 字符)
	//
	// 2 ICCID(20 位 SIM 卡对应编码)(选取该位后，格式为 ASCII 字符)
	SetRegisterContentType(registerContentType RegisterContentType) SetRegisterPacketCfgCommand
	// SetDataFormat 数据输入格式，取值范围 0-1
	//
	// 0 ASCII 格式
	//
	// 1 HEX 格式
	SetDataFormat(dataFormat DataFormat) SetRegisterPacketCfgCommand
	SetDataContent(dataContent string) SetRegisterPacketCfgCommand
	SetChannelNumber(channelNumber int) SetRegisterPacketCfgCommand
}

func NewSetRegisterPacketCfgCommand() SetRegisterPacketCfgCommand {
	return &setRegisterPacket{baseCommand: &baseCommand{Value: `"@DTU:0000:DTUID=%d,%d,%d,"%s",%d"`}, channelNumber: 1}
}

type setRegisterPacket struct {
	*baseCommand
	registerMode        RegisterMode
	registerContentType RegisterContentType
	dataFormat          DataFormat
	dataContent         string
	channelNumber       int
}

func (d *setRegisterPacket) SetRegisterMode(registerMode RegisterMode) SetRegisterPacketCfgCommand {
	d.registerMode = registerMode
	return d
}

func (d *setRegisterPacket) SetRegisterContentType(registerContentType RegisterContentType) SetRegisterPacketCfgCommand {
	d.registerContentType = registerContentType
	return d
}

func (d *setRegisterPacket) SetDataFormat(dataFormat DataFormat) SetRegisterPacketCfgCommand {
	d.dataFormat = dataFormat
	return d
}

// SetDataContent 数据内容，最大长度为 512Byte
func (d *setRegisterPacket) SetDataContent(dataContent string) SetRegisterPacketCfgCommand {
	d.dataContent = dataContent
	return d
}

func (d *setRegisterPacket) SetChannelNumber(channelNumber int) SetRegisterPacketCfgCommand {
	d.channelNumber = channelNumber
	return d
}

func (d *setRegisterPacket) String() string {
	return fmt.Sprintf(d.Value, d.registerMode, d.registerContentType, d.dataFormat, d.dataContent, d.channelNumber)
}

func (d *setRegisterPacket) Bytes() []byte {
	s := d.String()
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// SetHeartbeatCfgCommand 心跳配置
type SetHeartbeatCfgCommand interface {
	SetHeartbeatInterval(heartbeatInterval int) SetHeartbeatCfgCommand
	SetDataFormat(dataFormat DataFormat) SetHeartbeatCfgCommand
	SetDataContent(dataContent string) SetHeartbeatCfgCommand
	SetChannelNumber(channelNumber int) SetHeartbeatCfgCommand
}

func NewSetHeartbeatCfgCommand() SetHeartbeatCfgCommand {
	return &heartbeatConfigure{baseCommand: &baseCommand{Value: `@DTU:0000:KEEPALIVE=%d,%d,"%s",%d`}, channelNumber: 1}
}

type heartbeatConfigure struct {
	*baseCommand
	heartbeatInterval int
	dataFormat        DataFormat
	dataContent       string
	channelNumber     int
}

// SetHeartbeatInterval 心跳时间间隔，取值范围 0-65535
func (h *heartbeatConfigure) SetHeartbeatInterval(heartbeatInterval int) SetHeartbeatCfgCommand {
	h.heartbeatInterval = heartbeatInterval
	return h
}

// SetDataFormat 数据输入格式，取值范围 0-1
//
// 0 ASCII 模式
//
// 1 HEX 模式
func (h *heartbeatConfigure) SetDataFormat(dataFormat DataFormat) SetHeartbeatCfgCommand {
	h.dataFormat = dataFormat
	return h
}

// SetDataContent 数据内容，最大长度 256 Byte
func (h *heartbeatConfigure) SetDataContent(dataContent string) SetHeartbeatCfgCommand {
	h.dataContent = dataContent
	return h
}

func (h *heartbeatConfigure) SetChannelNumber(channelNumber int) SetHeartbeatCfgCommand {
	h.channelNumber = channelNumber
	return h
}

func (h *heartbeatConfigure) String() string {
	return fmt.Sprintf(h.Value, h.heartbeatInterval, h.dataFormat, h.dataContent, h.channelNumber)
}

func (h *heartbeatConfigure) Bytes() []byte {
	s1 := h.String()
	return unsafe.Slice(unsafe.StringData(s1), len(s1))
}

// SetHeartbeatAvoidanceCfgCommand 设置心跳避让配置
type SetHeartbeatAvoidanceCfgCommand interface {
	SetEnable(enable int) SetHeartbeatAvoidanceCfgCommand
	SetChannelNumber(channelNumber int) SetHeartbeatAvoidanceCfgCommand
}

func NewSetHeartbeatAvoidanceCfgCommand() SetHeartbeatAvoidanceCfgCommand {
	return &setHeartbeatAvoidanceConfig{baseCommand: &baseCommand{Value: `@DTU:0000:KEEPALIVEAVOID=%d,%d`}, enable: 1, channelNumber: 1}
}

type setHeartbeatAvoidanceConfig struct {
	*baseCommand
	enable        int
	channelNumber int
}

// SetEnable 业务心跳避让开关，0 避让， 1 不避让
func (h *setHeartbeatAvoidanceConfig) SetEnable(enable int) SetHeartbeatAvoidanceCfgCommand {
	h.enable = enable
	return h
}

func (h *setHeartbeatAvoidanceConfig) SetChannelNumber(channelNumber int) SetHeartbeatAvoidanceCfgCommand {
	h.channelNumber = channelNumber
	return h
}

func (h *setHeartbeatAvoidanceConfig) String() string {
	return fmt.Sprintf(h.Value, h.enable, h.channelNumber)
}

func (h *setHeartbeatAvoidanceConfig) Bytes() []byte {
	s1 := h.String()
	return unsafe.Slice(unsafe.StringData(s1), len(s1))
}

type readHeartbeatAvoidanceConfig struct {
	*baseCommand
}

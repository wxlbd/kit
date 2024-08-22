package at

import (
	"fmt"
	"unsafe"
)

// SetMQTTServerCfgCommand  MQTT服务器配置
type SetMQTTServerCfgCommand interface {
	Commander
	SetAddress(address string) SetMQTTServerCfgCommand
	SetPort(port int) SetMQTTServerCfgCommand
	SetChannelNumber(channelNumber int) SetMQTTServerCfgCommand
}

func NewSetMQTTServerCfgCommand() SetMQTTServerCfgCommand {
	return &setMqttServerConfig{baseCommand: &baseCommand{Value: `@DTU:0000:IPPORT="%s",%d,%d`}, channelNumber: 1}
}

type setMqttServerConfig struct {
	*baseCommand
	address       string
	port          int
	channelNumber int
}

func (m *setMqttServerConfig) SetAddress(address string) SetMQTTServerCfgCommand {
	m.address = address
	return m
}

func (m *setMqttServerConfig) SetPort(port int) SetMQTTServerCfgCommand {
	m.port = port
	return m
}

func (m *setMqttServerConfig) SetChannelNumber(channelNumber int) SetMQTTServerCfgCommand {
	m.channelNumber = channelNumber
	return m
}

func (m *setMqttServerConfig) String() string {
	return fmt.Sprintf(m.Value, m.address, m.port, m.channelNumber)
}

func (m *setMqttServerConfig) Bytes() []byte {
	s := m.String()
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// SetMQTTClientIDCfgCommand MQTT客户端ID配置
type SetMQTTClientIDCfgCommand interface {
	Commander
	SetClientID(clientID string) SetMQTTClientIDCfgCommand
	SetChannelNumber(channelNumber int) SetMQTTClientIDCfgCommand
}

func NewSetMQTTClientIDCfgCommand() SetMQTTClientIDCfgCommand {
	return &setMqttClientIDConfig{baseCommand: &baseCommand{Value: `@DTU:0000:CLIENTID="%s",%d`}, channelNumber: 1}
}

type setMqttClientIDConfig struct {
	*baseCommand
	clientID      string
	channelNumber int
}

func (m *setMqttClientIDConfig) SetClientID(clientID string) SetMQTTClientIDCfgCommand {
	m.clientID = clientID
	return m
}

func (m *setMqttClientIDConfig) SetChannelNumber(channelNumber int) SetMQTTClientIDCfgCommand {
	m.channelNumber = channelNumber
	return m
}

func (m *setMqttClientIDConfig) String() string {
	return fmt.Sprintf(m.Value, m.clientID, m.channelNumber)
}

func (m *setMqttClientIDConfig) Bytes() []byte {
	s := m.String()
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// SetMQTTUserPwdCfgCommand MQTT用户名密码配置
type SetMQTTUserPwdCfgCommand interface {
	Commander
	SetUsername(username string) SetMQTTUserPwdCfgCommand
	SetPassword(password string) SetMQTTUserPwdCfgCommand
	SetChannelNumber(channelNumber int) SetMQTTUserPwdCfgCommand
}

func NewSetMQTTUserPwdCfgCommand() SetMQTTUserPwdCfgCommand {
	return &setMqttUserPwdConfig{baseCommand: &baseCommand{Value: `@DTU:0000:USERPWD="%s","%s",%d`}, channelNumber: 1}
}

type setMqttUserPwdConfig struct {
	*baseCommand
	username      string
	password      string
	channelNumber int
}

// SetUsername 设置MQTT用户名
func (m *setMqttUserPwdConfig) SetUsername(username string) SetMQTTUserPwdCfgCommand {
	m.username = username
	return m
}

// SetPassword 设置MQTT密码
func (m *setMqttUserPwdConfig) SetPassword(password string) SetMQTTUserPwdCfgCommand {
	m.password = password
	return m
}

func (m *setMqttUserPwdConfig) SetChannelNumber(channelNumber int) SetMQTTUserPwdCfgCommand {
	m.channelNumber = channelNumber
	return m
}

func (m *setMqttUserPwdConfig) String() string {
	return fmt.Sprintf(m.Value, m.username, m.password, m.channelNumber)
}

func (m *setMqttUserPwdConfig) Bytes() []byte {
	s := m.String()
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// SetMQTTAutoSubCfgCommand MQTT自动订阅配置
type SetMQTTAutoSubCfgCommand interface {
	Commander
	SetEnable(enable int) SetMQTTAutoSubCfgCommand
	SetTopic(topic string) SetMQTTAutoSubCfgCommand
	SetQos(qos int) SetMQTTAutoSubCfgCommand
	SetChannelNumber(channelNumber int) SetMQTTAutoSubCfgCommand
}

func NewSetMQTTAutoSubCfgCommand() SetMQTTAutoSubCfgCommand {
	return &setMqttAutoSubConfig{baseCommand: &baseCommand{Value: `@DTU:0000:AUTOSUB=%d,"%s",%d,%d`}, channelNumber: 1}
}

type setMqttAutoSubConfig struct {
	*baseCommand
	enable        int
	topic         string
	qos           int
	channelNumber int
}

func (m *setMqttAutoSubConfig) SetEnable(enable int) SetMQTTAutoSubCfgCommand {
	m.enable = enable
	return m
}

func (m *setMqttAutoSubConfig) SetTopic(topic string) SetMQTTAutoSubCfgCommand {
	m.topic = topic
	return m
}

func (m *setMqttAutoSubConfig) SetQos(qos int) SetMQTTAutoSubCfgCommand {
	m.qos = qos
	return m
}

func (m *setMqttAutoSubConfig) SetChannelNumber(channelNumber int) SetMQTTAutoSubCfgCommand {
	m.channelNumber = channelNumber
	return m
}

func (m *setMqttAutoSubConfig) String() string {
	return fmt.Sprintf(m.Value, m.enable, m.topic, m.qos, m.channelNumber)
}

func (m *setMqttAutoSubConfig) Bytes() []byte {
	s := m.String()
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// SetMQTTAutoPubCfgCommand MQTT自动发布配置
type SetMQTTAutoPubCfgCommand interface {
	Commander
	SetEnable(enable int) SetMQTTAutoPubCfgCommand
	SetTopic(topic string) SetMQTTAutoPubCfgCommand
	SetQos(qos int) SetMQTTAutoPubCfgCommand
	SetSessionKeep(sessionKeep int) SetMQTTAutoPubCfgCommand
	SetChannelNumber(channelNumber int) SetMQTTAutoPubCfgCommand
}

func NewSetMQTTAutoPubCfgCommand() SetMQTTAutoPubCfgCommand {
	return &setMqttAutoPubConfigure{baseCommand: &baseCommand{Value: `@DTU:0000:AUTOPUB=%d,"%s",%d,%d,%d`}, channelNumber: 1}
}

type setMqttAutoPubConfigure struct {
	*baseCommand
	enable        int
	topic         string
	qos           int
	sessionKeep   int
	channelNumber int
}

// SetEnable MQTT推送使能，取值范围 0-1
//
// 0 不开启自动推送
//
// 1 开启自动推送
func (m *setMqttAutoPubConfigure) SetEnable(enable int) SetMQTTAutoPubCfgCommand {
	m.enable = enable
	return m
}

// SetTopic MQTT自动发布主题
func (m *setMqttAutoPubConfigure) SetTopic(topic string) SetMQTTAutoPubCfgCommand {
	m.topic = topic
	return m
}

// SetQos MQTT自动发布QoS
func (m *setMqttAutoPubConfigure) SetQos(qos int) SetMQTTAutoPubCfgCommand {
	m.qos = qos
	return m
}

// SetSessionKeep MQTT自动发布会话保持
func (m *setMqttAutoPubConfigure) SetSessionKeep(sessionKeep int) SetMQTTAutoPubCfgCommand {
	m.sessionKeep = sessionKeep
	return m
}

func (m *setMqttAutoPubConfigure) SetChannelNumber(channelNumber int) SetMQTTAutoPubCfgCommand {
	m.channelNumber = channelNumber
	return m
}

func (m *setMqttAutoPubConfigure) String() string {
	return fmt.Sprintf(m.Value, m.enable, m.topic, m.qos, m.sessionKeep, m.channelNumber)
}

func (m *setMqttAutoPubConfigure) Bytes() []byte {
	s := m.String()
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

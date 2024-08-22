package at

import (
	"testing"
)

func TestNewSetMQTTServerCfgCommand(t *testing.T) {
	cmd := NewSetMQTTServerCfgCommand().SetAddress("127.0.0.1").SetPort(1883).SetChannelNumber(1)
	if cmd.String() != "@DTU:0000:IPPORT=\"127.0.0.1\",1883,1" {
		t.Fatal(cmd.String())
	}
}

func TestNewSetMQTTClientIDCfgCommand(t *testing.T) {
	cmd := NewSetMQTTClientIDCfgCommand().SetClientID("test").SetChannelNumber(1)
	if cmd.String() != "@DTU:0000:CLIENTID=\"test\",1" {
		t.Fatal(cmd.String())
	}
}

func TestNewSetMQTTUserPwdCfgCommand(t *testing.T) {
	cmd := NewSetMQTTUserPwdCfgCommand().SetUsername("test").SetPassword("test").SetChannelNumber(1)
	if cmd.String() != `@DTU:0000:USERPWD="test","test",1` {
		t.Fatal(cmd.String())
	}
}

func TestNewSetMQTTAutoSubCfgCommand(t *testing.T) {
	cmd := NewSetMQTTAutoSubCfgCommand().SetEnable(1).SetQos(1).SetTopic("test")
	if cmd.String() != `@DTU:0000:AUTOSUB=1,"test",1,1` {
		t.Fatal(cmd.String())
	}
}

func TestNewSetMQTTAutoPubCfgCommand(t *testing.T) {
	cmd := NewSetMQTTAutoPubCfgCommand().SetSessionKeep(1).SetChannelNumber(4).SetQos(2).SetTopic("test").SetEnable(0)
	if cmd.String() != `@DTU:0000:AUTOPUB=0,"test",2,1,4` {
		t.Fatal(cmd.String())
	}
}

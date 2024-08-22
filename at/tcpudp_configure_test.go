package at

import (
	"testing"
)

func TestNewSetTCPUDPServerCfgCommand(t *testing.T) {
	cmd := NewSetTCPUDPServerCfgCommand().SetAddress("127.0.0.1").SetPort(1883).SetChannelNumber(1).SetTransferProto(TransferProtoTCP)
	if cmd.String() != `@DTU:0000:DSCADDR=1,"TCP","127.0.0.1",1883` {
		t.Fatal(cmd.String())
	}
}

func TestNewSetRegisterPacketCfgCommand(t *testing.T) {
	cmd := NewSetRegisterPacketCfgCommand().SetRegisterMode(RegisterModeData).SetDataFormat(DataFormatHex).SetDataContent("register").SetRegisterContentType(RegisterContentTypeIMEI).SetChannelNumber(1)
	if cmd.String() != `@DTU:0000:DTUID=2,1,1,"register",1` {
		t.Fatal(cmd.String())
	}
}

func TestNewSetHeartbeatCfgCommand(t *testing.T) {
	cmd := NewSetHeartbeatCfgCommand().SetHeartbeatInterval(60).SetDataFormat(DataFormatHex).SetDataContent("heartbeat").SetChannelNumber(1)
	if cmd.String() != `@DTU:0000:KEEPALIVE=60,1,"heartbeat",1` {
		t.Fatal(cmd.String())
	}
}

func TestNewSetHeartbeatAvoidanceCfgCommand(t *testing.T) {
	cmd := NewSetHeartbeatAvoidanceCfgCommand().SetEnable(1).SetChannelNumber(3)
	if cmd.String() != `@DTU:0000:KEEPALIVEAVOID=1,3` {
		t.Fatal(cmd.String())
	}
}

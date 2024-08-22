package at

import (
	"fmt"
	"testing"
)

func TestNewJsonDataPointConfig(t *testing.T) {
	c := NewPollJsonCfgCommand()
	fmt.Println(c.SetIndex(1).SetKey("test").SetEnable(1).SetSource(0).String())
}

func TestNewPollJsonRegConfig(t *testing.T) {
	c := NewPollJsonRegConfig()
	c.SetIndex(1).SetSlave(1).SetFuncCode(3).SetRegAddr(0).SetDataType(1).SetByteOrder(1).SetReserve(1).SetQuote(0).SetInterval(500).String()
	fmt.Println(c.String())
}

func TestNewPoll(t *testing.T) {
	c := NewSetPollCommand()
	c.SetCollectType(CollectTypeJSON).SetPollInterval(60).SetDataFormat(1)
	fmt.Println(c.String())
}

func TestNewSetPollStrCommand(t *testing.T) {
	cmd := NewSetPollStrCommand().SetStr("test").SetCRCEnable(1).SetIndex(20).SetEnable(0)
	if cmd.String() != `@DTU:0000:POLLSTR=20,0,1,"test"` {
		t.Fatal(cmd.String())
	}
}

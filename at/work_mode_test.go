package at

import (
	"testing"
)

func TestNewSetWorkModeCommand(t *testing.T) {
	cmd := NewSetWorkModeCommand().SetMode(DTUWorkModeMQTT).SetChannelNumber(1)
	if cmd.String() != `@DTU:0000:DTUMODE=2,1` {
		t.Fatal(cmd.String())
	}
}

package at

import (
	"testing"
)

func TestNewSetSerialCommand(t *testing.T) {
	c := NewSetSerialCommand().SetBaudRate(BaudRate9600).SetDataBits(DataBits8).SetStopBits(StopBits2).SetParity(ParityEven)
	if c.String() != "@DTU:0000:UARTCFG=9600,1,2,1" {
		t.Fatal(c.String())
	}
}

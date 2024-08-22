package at

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewJsonDataPointConfig(t *testing.T) {
	c := NewPollJsonConfigureCommand()
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
	tests := []struct {
		name string
		want SetPollStrCommand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSetPollStrCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSetPollStrCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pollJsonRegConfig_Bytes(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		slave       int
		funcCode    int
		regAddr     int
		dataType    int
		byteOrder   int
		reserve     int
		unit        string
		quote       int
		interval    int
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pollJsonRegConfig{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				slave:       tt.fields.slave,
				funcCode:    tt.fields.funcCode,
				regAddr:     tt.fields.regAddr,
				dataType:    tt.fields.dataType,
				byteOrder:   tt.fields.byteOrder,
				reserve:     tt.fields.reserve,
				unit:        tt.fields.unit,
				quote:       tt.fields.quote,
				interval:    tt.fields.interval,
			}
			if got := p.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pollJsonRegConfig_SetByteOrder(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		slave       int
		funcCode    int
		regAddr     int
		dataType    int
		byteOrder   int
		reserve     int
		unit        string
		quote       int
		interval    int
	}
	type args struct {
		byteOrder int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PollJsonRegCfgCommand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pollJsonRegConfig{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				slave:       tt.fields.slave,
				funcCode:    tt.fields.funcCode,
				regAddr:     tt.fields.regAddr,
				dataType:    tt.fields.dataType,
				byteOrder:   tt.fields.byteOrder,
				reserve:     tt.fields.reserve,
				unit:        tt.fields.unit,
				quote:       tt.fields.quote,
				interval:    tt.fields.interval,
			}
			if got := p.SetByteOrder(tt.args.byteOrder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetByteOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pollJsonRegConfig_SetDataType(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		slave       int
		funcCode    int
		regAddr     int
		dataType    int
		byteOrder   int
		reserve     int
		unit        string
		quote       int
		interval    int
	}
	type args struct {
		dataType int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PollJsonRegCfgCommand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pollJsonRegConfig{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				slave:       tt.fields.slave,
				funcCode:    tt.fields.funcCode,
				regAddr:     tt.fields.regAddr,
				dataType:    tt.fields.dataType,
				byteOrder:   tt.fields.byteOrder,
				reserve:     tt.fields.reserve,
				unit:        tt.fields.unit,
				quote:       tt.fields.quote,
				interval:    tt.fields.interval,
			}
			if got := p.SetDataType(tt.args.dataType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDataType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pollJsonRegConfig_SetFuncCode(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		slave       int
		funcCode    int
		regAddr     int
		dataType    int
		byteOrder   int
		reserve     int
		unit        string
		quote       int
		interval    int
	}
	type args struct {
		funcCode int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PollJsonRegCfgCommand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pollJsonRegConfig{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				slave:       tt.fields.slave,
				funcCode:    tt.fields.funcCode,
				regAddr:     tt.fields.regAddr,
				dataType:    tt.fields.dataType,
				byteOrder:   tt.fields.byteOrder,
				reserve:     tt.fields.reserve,
				unit:        tt.fields.unit,
				quote:       tt.fields.quote,
				interval:    tt.fields.interval,
			}
			if got := p.SetFuncCode(tt.args.funcCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFuncCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pollJsonRegConfig_SetIndex(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		slave       int
		funcCode    int
		regAddr     int
		dataType    int
		byteOrder   int
		reserve     int
		unit        string
		quote       int
		interval    int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PollJsonRegCfgCommand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pollJsonRegConfig{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				slave:       tt.fields.slave,
				funcCode:    tt.fields.funcCode,
				regAddr:     tt.fields.regAddr,
				dataType:    tt.fields.dataType,
				byteOrder:   tt.fields.byteOrder,
				reserve:     tt.fields.reserve,
				unit:        tt.fields.unit,
				quote:       tt.fields.quote,
				interval:    tt.fields.interval,
			}
			if got := p.SetIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pollJsonRegConfig_SetInterval(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		slave       int
		funcCode    int
		regAddr     int
		dataType    int
		byteOrder   int
		reserve     int
		unit        string
		quote       int
		interval    int
	}
	type args struct {
		interval int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PollJsonRegCfgCommand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pollJsonRegConfig{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				slave:       tt.fields.slave,
				funcCode:    tt.fields.funcCode,
				regAddr:     tt.fields.regAddr,
				dataType:    tt.fields.dataType,
				byteOrder:   tt.fields.byteOrder,
				reserve:     tt.fields.reserve,
				unit:        tt.fields.unit,
				quote:       tt.fields.quote,
				interval:    tt.fields.interval,
			}
			if got := p.SetInterval(tt.args.interval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pollJsonRegConfig_SetQuote(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		slave       int
		funcCode    int
		regAddr     int
		dataType    int
		byteOrder   int
		reserve     int
		unit        string
		quote       int
		interval    int
	}
	type args struct {
		quote int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PollJsonRegCfgCommand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pollJsonRegConfig{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				slave:       tt.fields.slave,
				funcCode:    tt.fields.funcCode,
				regAddr:     tt.fields.regAddr,
				dataType:    tt.fields.dataType,
				byteOrder:   tt.fields.byteOrder,
				reserve:     tt.fields.reserve,
				unit:        tt.fields.unit,
				quote:       tt.fields.quote,
				interval:    tt.fields.interval,
			}
			if got := p.SetQuote(tt.args.quote); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pollJsonRegConfig_SetRegAddr(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		slave       int
		funcCode    int
		regAddr     int
		dataType    int
		byteOrder   int
		reserve     int
		unit        string
		quote       int
		interval    int
	}
	type args struct {
		regAddr int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PollJsonRegCfgCommand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pollJsonRegConfig{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				slave:       tt.fields.slave,
				funcCode:    tt.fields.funcCode,
				regAddr:     tt.fields.regAddr,
				dataType:    tt.fields.dataType,
				byteOrder:   tt.fields.byteOrder,
				reserve:     tt.fields.reserve,
				unit:        tt.fields.unit,
				quote:       tt.fields.quote,
				interval:    tt.fields.interval,
			}
			if got := p.SetRegAddr(tt.args.regAddr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRegAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pollJsonRegConfig_SetReserve(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		slave       int
		funcCode    int
		regAddr     int
		dataType    int
		byteOrder   int
		reserve     int
		unit        string
		quote       int
		interval    int
	}
	type args struct {
		reserve int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PollJsonRegCfgCommand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pollJsonRegConfig{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				slave:       tt.fields.slave,
				funcCode:    tt.fields.funcCode,
				regAddr:     tt.fields.regAddr,
				dataType:    tt.fields.dataType,
				byteOrder:   tt.fields.byteOrder,
				reserve:     tt.fields.reserve,
				unit:        tt.fields.unit,
				quote:       tt.fields.quote,
				interval:    tt.fields.interval,
			}
			if got := p.SetReserve(tt.args.reserve); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetReserve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pollJsonRegConfig_SetSlave(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		slave       int
		funcCode    int
		regAddr     int
		dataType    int
		byteOrder   int
		reserve     int
		unit        string
		quote       int
		interval    int
	}
	type args struct {
		slave int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PollJsonRegCfgCommand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pollJsonRegConfig{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				slave:       tt.fields.slave,
				funcCode:    tt.fields.funcCode,
				regAddr:     tt.fields.regAddr,
				dataType:    tt.fields.dataType,
				byteOrder:   tt.fields.byteOrder,
				reserve:     tt.fields.reserve,
				unit:        tt.fields.unit,
				quote:       tt.fields.quote,
				interval:    tt.fields.interval,
			}
			if got := p.SetSlave(tt.args.slave); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSlave() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pollJsonRegConfig_SetUnit(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		slave       int
		funcCode    int
		regAddr     int
		dataType    int
		byteOrder   int
		reserve     int
		unit        string
		quote       int
		interval    int
	}
	type args struct {
		unit string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   PollJsonRegCfgCommand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pollJsonRegConfig{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				slave:       tt.fields.slave,
				funcCode:    tt.fields.funcCode,
				regAddr:     tt.fields.regAddr,
				dataType:    tt.fields.dataType,
				byteOrder:   tt.fields.byteOrder,
				reserve:     tt.fields.reserve,
				unit:        tt.fields.unit,
				quote:       tt.fields.quote,
				interval:    tt.fields.interval,
			}
			if got := p.SetUnit(tt.args.unit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUnit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pollJsonRegConfig_String(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		slave       int
		funcCode    int
		regAddr     int
		dataType    int
		byteOrder   int
		reserve     int
		unit        string
		quote       int
		interval    int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pollJsonRegConfig{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				slave:       tt.fields.slave,
				funcCode:    tt.fields.funcCode,
				regAddr:     tt.fields.regAddr,
				dataType:    tt.fields.dataType,
				byteOrder:   tt.fields.byteOrder,
				reserve:     tt.fields.reserve,
				unit:        tt.fields.unit,
				quote:       tt.fields.quote,
				interval:    tt.fields.interval,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setPollStr_Bytes(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		enable      int
		crcEnable   int
		str         string
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &setPollStr{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				enable:      tt.fields.enable,
				crcEnable:   tt.fields.crcEnable,
				str:         tt.fields.str,
			}
			if got := p.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setPollStr_SetCRCEnable(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		enable      int
		crcEnable   int
		str         string
	}
	type args struct {
		crcEnable int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   SetPollStrCommand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &setPollStr{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				enable:      tt.fields.enable,
				crcEnable:   tt.fields.crcEnable,
				str:         tt.fields.str,
			}
			if got := p.SetCRCEnable(tt.args.crcEnable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCRCEnable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setPollStr_SetEnable(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		enable      int
		crcEnable   int
		str         string
	}
	type args struct {
		enable int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   SetPollStrCommand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &setPollStr{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				enable:      tt.fields.enable,
				crcEnable:   tt.fields.crcEnable,
				str:         tt.fields.str,
			}
			if got := p.SetEnable(tt.args.enable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetEnable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setPollStr_SetIndex(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		enable      int
		crcEnable   int
		str         string
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   SetPollStrCommand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &setPollStr{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				enable:      tt.fields.enable,
				crcEnable:   tt.fields.crcEnable,
				str:         tt.fields.str,
			}
			if got := p.SetIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setPollStr_SetStr(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		enable      int
		crcEnable   int
		str         string
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   SetPollStrCommand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &setPollStr{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				enable:      tt.fields.enable,
				crcEnable:   tt.fields.crcEnable,
				str:         tt.fields.str,
			}
			if got := p.SetStr(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setPollStr_String(t *testing.T) {
	type fields struct {
		baseCommand *baseCommand
		index       int
		enable      int
		crcEnable   int
		str         string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &setPollStr{
				baseCommand: tt.fields.baseCommand,
				index:       tt.fields.index,
				enable:      tt.fields.enable,
				crcEnable:   tt.fields.crcEnable,
				str:         tt.fields.str,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

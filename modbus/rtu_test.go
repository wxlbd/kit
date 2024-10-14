package modbus

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/sigurn/crc16"
	"reflect"
	"testing"
)

func TestRTUPacker_Decode(t *testing.T) {
	type fields struct {
		SlaveId byte
	}
	type args struct {
		adu []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantPdu *ProtocolDataUnit
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &RTUPacker{
				SlaveId: tt.fields.SlaveId,
			}
			gotPdu, err := p.Decode(tt.args.adu)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPdu, tt.wantPdu) {
				t.Errorf("Decode() gotPdu = %v, want %v", gotPdu, tt.wantPdu)
			}
		})
	}
}

func TestRTUPacker_Encode(t *testing.T) {
	type fields struct {
		SlaveId byte
	}
	type args struct {
		pdu *ProtocolDataUnit
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantAdu []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &RTUPacker{
				SlaveId: tt.fields.SlaveId,
			}
			gotAdu, err := p.Encode(tt.args.pdu)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAdu, tt.wantAdu) {
				t.Errorf("Encode() gotAdu = %v, want %v", gotAdu, tt.wantAdu)
			}
		})
	}
}

func TestRTUPacker_Verify(t *testing.T) {
	type fields struct {
		SlaveId byte
	}
	type args struct {
		aduRequest  []byte
		aduResponse []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &RTUPacker{
				SlaveId: tt.fields.SlaveId,
			}
			if err := p.Verify(tt.args.aduRequest, tt.args.aduResponse); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCrc16(t *testing.T) {

	crc := crc16.Checksum([]byte{01, 06, 01, 0x45, 01, 03}, crc16.MakeTable(crc16.CRC16_MODBUS))
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, crc)
	fmt.Println(hex.EncodeToString(bs))
}

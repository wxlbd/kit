package modbus

import (
	"fmt"
	"github.com/sigurn/crc16"
)

const (
	rtuMinSize = 4
	rtuMaxSize = 256

	rtuExceptionSize = 5
)

type RTUPacker struct {
	SlaveId byte
}

func (p *RTUPacker) Encode(pdu *ProtocolDataUnit) (adu []byte, err error) {
	length := len(pdu.Data) + 4
	if length > rtuMaxSize {
		err = fmt.Errorf("modbus: length of data '%v' must not be bigger than '%v'", length, rtuMaxSize)
		return
	}

	adu = make([]byte, length)
	adu[0] = p.SlaveId
	adu[1] = pdu.FunctionCode

	copy(adu[2:], pdu.Data)

	crc := crc16.Checksum(adu, crc16.MakeTable(crc16.CRC16_MODBUS))
	adu[length-1] = byte(crc >> 8)
	adu[length-2] = byte(crc)
	return
}

func (p *RTUPacker) Decode(adu []byte) (pdu *ProtocolDataUnit, err error) {
	length := len(adu)

	crc := crc16.Checksum(adu[:length-2], crc16.MakeTable(crc16.CRC16_MODBUS))
	checksum := uint16(adu[length-1])<<8 | uint16(adu[length-2])
	if checksum != crc {
		err = fmt.Errorf("modbus: response crc '%v' does not match expected '%v'", checksum, crc)
	}

	pdu = &ProtocolDataUnit{
		FunctionCode: adu[1],
		Data:         adu[2 : length-2],
	}
	return
}

func (p *RTUPacker) Verify(aduRequest []byte, aduResponse []byte) (err error) {
	length := len(aduResponse)
	// Minimum size (including address, function and CRC)
	if length < rtuMinSize {
		err = fmt.Errorf("modbus: response length '%v' does not meet minimum '%v'", length, rtuMinSize)
		return
	}
	// Slave address must match
	if aduResponse[0] != aduRequest[0] {
		err = fmt.Errorf("modbus: response slave id '%v' does not match request '%v'", aduResponse[0], aduRequest[0])
		return
	}
	return
}

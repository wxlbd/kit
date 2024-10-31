package ymodem

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"strconv"

	"github.com/sigurn/crc16"
)

// 定义YModem协议中的控制字符
const (
	SOH  byte = 0x01 // YModem协议的起始帧标志 128字节数据包
	STX  byte = 0x02 // YModem-1K协议的起始帧标志 1024字节数据包
	EOT  byte = 0x04 // 传输结束标志
	ACK  byte = 0x06 // 确认接收字符
	NAK  byte = 0x15 // 未确认接收字符
	CAN  byte = 0x18 // 取消传输字符
	POLL byte = 0x43 // 轮询字符，用于请求发送
)

type Frame struct {
	start byte   // 帧开始字节 SOH 或 STX
	seq   byte   // 包序号
	seqCh byte   // 包序号的补码
	data  []byte // 数据
	crc   uint16 // CRC16校验
}

var (
	ErrDataTooLong = errors.New("data too long")
	ErrNotAcked    = errors.New("not acked")
	ErrNotPoll     = errors.New("can't receive POLL")
)

// NewFrame 构造数据帧
func NewFrame(seq byte, data []byte) (*Frame, error) {
	l := len(data)
	if l > 1024 {
		return nil, ErrDataTooLong
	}
	var padded []byte
	start := STX
	if l <= 128 {
		start = SOH
		padded = make([]byte, 128)
		copy(padded[:l], data)
	} else {
		padded = make([]byte, 1024)
		copy(padded[:l], data)
	}
	crc := crc16.Checksum(padded, crc16.MakeTable(crc16.CRC16_XMODEM))
	return &Frame{
		start: start,
		seq:   seq,
		seqCh: ^seq,
		data:  padded,
		crc:   crc,
	}, nil
}

// Bytes 将帧转换为字节数组
func (f *Frame) Bytes() []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(f.start)
	buf.WriteByte(f.seq)
	buf.WriteByte(f.seqCh)
	buf.Write(f.data)
	_ = binary.Write(buf, binary.BigEndian, f.crc)
	return buf.Bytes()
}

// BuildStartFrame 构造起始帧
func BuildStartFrame(filename string, filesize int) (*Frame, error) {
	bs := new(bytes.Buffer)
	bs.WriteString(filename)
	bs.WriteByte(0)
	bs.WriteString(strconv.Itoa(filesize))
	bs.WriteByte(0)
	frame, err := NewFrame(0, bs.Bytes())
	if err != nil {
		return nil, err
	}
	return frame, nil
}

// BuildEndFrame 构造结束帧
func BuildEndFrame() (*Frame, error) {
	return NewFrame(0, nil)
}

// BuildDataFrames 构造数据帧
func BuildDataFrames(fileData []byte) ([]*Frame, error) {
	reader := bytes.NewReader(fileData)
	var frames []*Frame
	var seq byte = 1
	for {
		data := make([]byte, 1024)
		_, err := reader.Read(data)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if err == io.EOF {
			break
		}
		frame, err := NewFrame(seq, data)
		if err != nil {
			return nil, err
		}
		frames = append(frames, frame)
		seq++
	}
	return frames, nil
}

// Send 发送文件
func Send(rw io.ReadWriter, filename string, fileData []byte) error {
	buf := make([]byte, 1)
	// 等待POLL
	if _, err := rw.Read(buf); err != nil {
		return err
	}
	if buf[0] == POLL {
		// 发送起始帧
		frame, err := BuildStartFrame(filename, len(fileData))
		if err != nil {
			return err
		}
		if _, err := rw.Write(frame.Bytes()); err != nil {
			return err
		}
		if _, err := rw.Read(buf); err != nil {
			return err
		}
		if buf[0] != ACK {
			return ErrNotAcked
		}
		if _, err := rw.Read(buf); err != nil {
			return err
		}
		if buf[0] != POLL {
			return ErrNotPoll
		}
		// 发送数据帧
		frames, err := BuildDataFrames(fileData)
		if err != nil {
			return err
		}
		for _, frame := range frames {
			if err := sendWithAck(rw, frame.Bytes()); err != nil {
				return err
			}
		}
		if _, err := rw.Write([]byte{EOT}); err != nil {
			return err
		}
		if _, err := rw.Read(buf); err != nil {
			return err
		}
		if buf[0] != NAK {
			return errors.New("not NAK")
		}
		// 第二次发送结束帧
		if err := sendWithAck(rw, []byte{EOT}); err != nil {
			return err
		}
		if _, err := rw.Read(buf); err != nil {
			return err
		}
		if buf[0] != POLL {
			return ErrNotPoll
		}
	}
	return nil
}

// 发送数据帧并等待ACK
func sendWithAck(rw io.ReadWriter, data []byte) error {
ReSend:
	_, err := rw.Write(data)
	if err != nil {
		return err
	}
	buf := make([]byte, 1)
	_, err = rw.Read(buf)
	if err != nil {
		return err
	}
	if buf[0] == NAK {
		goto ReSend
	}
	if buf[0] != ACK {
		return ErrNotAcked
	}
	return nil
}

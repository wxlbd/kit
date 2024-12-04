package at

import (
	"errors"
	"go.bug.st/serial"
	"io"
	"net"
	"strings"
	"time"
)

type Mode int

const (
	LocalMode  Mode = 0
	RemoteMode Mode = 1
)

type Device struct {
	conn io.ReadWriter
	mode Mode
}

// NewDevice 创建一个新的AT设备，支持本地模式（go.bug.st/serial）和远程模式（net.Conn）
func NewDevice(conn io.ReadWriter, mode Mode) *Device {
	return &Device{conn: conn, mode: mode}
}

// SendATCommand 发送AT命令,并等待响应,如果响应包含ERROR则返回错误
func (d *Device) SendATCommand(command []byte) (string, error) {
	command = append(command, '\r', '\n')
	_, err := d.conn.Write(command)
	if err != nil {
		return "", err
	}
	if d.mode == LocalMode {
		if err := d.conn.(serial.Port).SetReadTimeout(2 * time.Second); err != nil {
			return "", err
		}
	} else {
		if err := d.conn.(net.Conn).SetReadDeadline(time.Now().Add(2 * time.Second)); err != nil {
			return "", err
		}
	}
	buf := make([]byte, 1024)
	var response string
	for {
		n, err := d.conn.Read(buf)
		if err != nil {
			return "", err
		}
		response += string(buf[:n])
		if response[len(response)-2:] == "\r\n" {
			break
		}
	}
	if strings.Contains(response, "ERROR") {
		return "", errors.New(response)
	}
	return response[:len(response)-2], nil
}

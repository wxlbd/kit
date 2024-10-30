package ymodem

import (
	"fmt"
	"os"
	"testing"

	"go.bug.st/serial"
)

func TestSend(t *testing.T) {
	list, err := serial.GetPortsList()
	if err != nil {
		return
	}
	fmt.Println(list)
	port, err := serial.Open("/dev/tty.usbserial-14440", &serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
		Parity:   0,
		StopBits: 0,
	})
	if err != nil {
		t.Fatal(err)
		return
	}
	defer port.Close()
	fileInfo, err := os.Stat("/Users/wangxinlong/Downloads/N32_RTOS_WK_202410300913.bin")
	if err != nil {
		t.Fatal("open file error", err)
		return
	}
	data, err := os.ReadFile("/Users/wangxinlong/Downloads/N32_RTOS_WK_202410300913.bin")
	if err != nil {
		t.Fatal("open file error", err)
		return
	}
	// port.Write([]byte("01060130000149F9"))
	err = Send(port, fileInfo.Name(), data)
	if err != nil {
		t.Fatal(err)
		return
	}
}

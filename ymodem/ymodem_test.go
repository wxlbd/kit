package ymodem

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"os"
	"testing"
	"time"

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
	fileInfo, err := os.Stat("/Users/wxl/Downloads/N32_RTOS_WK_202410300913.bin")
	if err != nil {
		t.Fatal("open file error", err)
		return
	}
	data, err := os.ReadFile("/Users/wxl/Downloads/N32_RTOS_WK_202410300913.bin")
	if err != nil {
		t.Fatal("open file error", err)
		return
	}
	if _, err := port.Write([]byte{0x72, 0x62, 0x20, 0x2D, 0x45, 0x0D, 0x0A}); err != nil {
		t.Fatal(err)
		return
	}
	err = Send(port, fileInfo.Name(), data)
	if err != nil {
		t.Fatal(err)
		return
	}
}

func TestBar(t *testing.T) {
	bar := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}
}

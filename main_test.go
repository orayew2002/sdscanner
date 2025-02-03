package sdscanner

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestScanning(t *testing.T) {
	_ = NewSdScanner(NewSdConnection(), SdDisconnected())

	for {
		time.Sleep(1 * time.Second)
	}
}

func NewSdConnection() func(d Disk) {
	return func(d Disk) {
		log.Println(fmt.Sprintf("new connection: %s", d.Path))
	}
}

func SdDisconnected() func(d Disk) {
	return func(d Disk) {
		log.Println(fmt.Sprintf("connection disconnected: %s", d.Path))
	}
}

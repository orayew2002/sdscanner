package sdscanner

import (
	"reflect"
	"runtime"
	"sync"
	"time"
)

// NewSdScanner there  when new device detected
func NewSdScanner(onConnected, onDisconnected func(disk Disk)) *SdScanner {
	var scanner = SdScanner{
		Disks:        make([]string, 0),
		Mu:           new(sync.Mutex),
		OnConnect:    onConnected,
		OnDisconnect: onDisconnected,
	}

	go func() {
		timeTicker := time.NewTicker(1 * time.Second)

		for {
			<-timeTicker.C

			switch osType := runtime.GOOS; osType {
			case "darwin":
				scanner.SetList(checkMacDrives())
			case "linux":
				scanner.SetList(checkLinuxDrives())
			default:
				break
			}
		}
	}()

	return &scanner
}

func (d *SdScanner) SetList(list []string) {
	if reflect.DeepEqual(d.Disks, list) {
		return
	}

	if len(d.Disks) < len(list) {
		if len(d.Disks) == 0 {
			d.Disks = list
			return
		}

		d.OnConnect(Disk{Path: getDifference(d.Disks, list)})
		d.Disks = list
		return
	}

	d.OnDisconnect(Disk{Path: getDifference(list, d.Disks)})
	d.Disks = list
}

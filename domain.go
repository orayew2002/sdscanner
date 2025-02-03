package sdscanner

import "sync"

type SdScanner struct {
	Mu *sync.Mutex

	Disks []string

	OnConnect    func(d Disk)
	OnDisconnect func(d Disk)
}

type Disk struct {
	Path string
}

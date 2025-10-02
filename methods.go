package sdscanner

import (
	"bytes"
	"os/exec"
	"reflect"
	"slices"
	"strings"
	"sync"
	"time"
)

type SdScanner struct {
	Mu           *sync.Mutex
	Disks        []string
	OnConnect    func(d string)
	OnDisconnect func(d string)
}

func NewSdScanner(onConnected, onDisconnected func(d string)) *SdScanner {
	scanner := &SdScanner{
		Disks:        make([]string, 0),
		Mu:           new(sync.Mutex),
		OnConnect:    onConnected,
		OnDisconnect: onDisconnected,
	}

	return scanner
}

func (d *SdScanner) Run() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		d.run()
	}
}

func (d *SdScanner) run() {
	d.Mu.Lock()
	defer d.Mu.Unlock()

	list, err := listBlockDevices()
	if err != nil {
		panic(err.Error())
	}

	if reflect.DeepEqual(d.Disks, list) {
		return
	}

	newDisks := getDifference(d.Disks, list)
	for _, path := range newDisks {
		if d.OnConnect != nil {
			d.OnConnect(path)
		}
	}

	removedDisks := getDifference(list, d.Disks)
	for _, path := range removedDisks {
		if d.OnDisconnect != nil {
			d.OnDisconnect(path)
		}
	}

	d.Disks = list
}

func getDifference(a1, a2 []string) []string {
	var diff []string
	for _, a := range a2 {
		if !slices.Contains(a1, a) {
			diff = append(diff, a)
		}
	}

	return diff
}

func listBlockDevices() ([]string, error) {
	cmd := exec.Command("lsblk", "-ln", "-o", "NAME,TYPE")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var disks []string
	lines := bytes.Split(out, []byte("\n"))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		fields := strings.Fields(string(line))
		if len(fields) < 2 {
			continue
		}

		name := fields[0]
		typ := fields[1]

		if typ == "part" {
			disks = append(disks, "/dev/"+name)
		}
	}

	return disks, nil
}

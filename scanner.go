package sdscanner

import (
	"fmt"
	"log"
	"os"
	"slices"
)

func checkMacDrives() []string {
	return checkDrives([]string{"/Volumes/"})
}

func checkLinuxDrives() []string {
	return checkDrives([]string{"/mnt/", "/media/", "/dev/"})
}

func checkDrives(mountPaths []string) []string {
	var list []string

	for _, mountPath := range mountPaths {
		drives, err := os.ReadDir(mountPath)
		if err != nil {
			log.Printf(fmt.Sprintf("failed scan disks"))
			return list
		}

		for _, drive := range drives {
			list = append(list, mountPath+drive.Name())
		}
	}

	return list
}

func getDifference(a1, a2 []string) string {
	var r string

	for _, a := range a2 {
		if !slices.Contains(a1, a) {
			return a
		}
	}

	return r
}

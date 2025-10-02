package main

import (
	"fmt"

	"github.com/orayew2002/sdscanner"
)

func main() {
	sd := sdscanner.NewSdScanner(func(d string) {
		fmt.Println("mount", d)
	}, func(d string) {
		fmt.Println("unmount", d)
	})

	sd.Run()
}

package main

import (
	"fmt"

	"github.com/orayew2002/sdscanner"
)

func main() {
	fmt.Println("service started working")

	sd := sdscanner.NewSdScanner(func(d string) {
		fmt.Println("mount", d)
	}, func(d string) {
		fmt.Println("unmount", d)
	})

	sd.Run()
}

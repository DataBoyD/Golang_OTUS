package main

import (
	"fmt"

	"github.com/beevik/ntp"
)

func main() {
	fmt.Println("Hello, world!")
	time, err := ntp.Time("")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(time)
}

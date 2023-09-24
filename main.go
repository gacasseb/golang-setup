package main

import (
	"fmt"
)

func main() {
	for {
		time.Sleep(10 * time.Second)
		fmt.Println("Running")
	}
}
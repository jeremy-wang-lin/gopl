package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	for idx, value := range os.Args[1:] {
		fmt.Println(strconv.Itoa(idx) + ": " + value)
	}
}

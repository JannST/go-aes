package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("this is a test my dear friend")

	buffer := make([]byte, 4)

	for {
		n, err := reader.Read(buffer)
		fmt.Println(n, err, buffer[:n])
		if err != nil {
			break
		}
	}
	fmt.Println("at the end")
}

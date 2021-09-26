package main

import (
	"fmt"
	"time"
)

func main() {

	channel := time.After(6 * time.Second)

	count := 0
	result := false

	for {
		select {
		case <-channel:
			result = true

		default:
			time.Sleep(time.Second)
			fmt.Println(count)
			count += 1
		}

		if result {
			break
		}
	}
}

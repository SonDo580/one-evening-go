package main

import "fmt"

var count_allocated int = 0

func AllocateBuffer() *string {
	if count_allocated == 3 {
		return nil
	}
	count_allocated++
	return new(string)
}

func main() {
	var buffers []*string

	for {
		b := AllocateBuffer()
		if b == nil {
			break
		}

		buffers = append(buffers, b)
	}

	fmt.Println("Allocated", len(buffers), "buffers")
}

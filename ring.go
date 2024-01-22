package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "value" + strconv.Itoa(i)
		data = data.Next()
	}

	// data.Value = "value 1"

	// data = data.Next()
	// data.Value = "value 2"

	// data = data.Next()
	// data.Value = "value 3"

	// data = data.Next()
	// data.Value = "value 4"

	// data = data.Next()
	// data.Value = "value 5"

	data.Do(func(value any) {
		fmt.Println(value)
	})
}

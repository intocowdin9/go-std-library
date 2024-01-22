package main

import (
	"container/list"
	"fmt"
)

func main() {

	var data *list.List = list.New()

	data.PushBack("james")
	data.PushBack("bond")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)

	next := head.Next()
	fmt.Println(next)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

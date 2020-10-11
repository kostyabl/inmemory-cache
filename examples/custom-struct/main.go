package main

import (
	"fmt"
	"github.com/kostyabl/inmemory-cache/cache"
)

//CItem - custom item
type CItem struct {
	Key string
	Val1 int64
	Val2 string
}

func (c CItem) KeyValue() interface{} {
	return c.Key
}

func (c CItem) Item() interface{} {
	return c
}

func main() {
	ic := cache.New()
	item := CItem{
		Key:  "some",
		Val1: 12,
		Val2: "bbb",
	}

	ic.StoreItem(item, 0)

	iItem, err := ic.ItemByKey("some")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d, %s", iItem.(CItem).Val1, iItem.(CItem).Val2)
}

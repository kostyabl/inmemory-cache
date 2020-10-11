package main

import (
	"fmt"
	"github.com/kostyabl/inmemory-cache/cache"
)

func main() {
	ic := cache.New()
	item := cache.Item{
		Key: 1,
		Value: "some val",
	}
	ic.StoreItem(item, 0)

	iItem, err := ic.ItemByKey(int64(1))
	if err != nil {
		panic(err)
	}

	fmt.Println(iItem.(cache.Item).Value)
}

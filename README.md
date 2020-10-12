# Simple in-memory cache
Simple in-memory cache written in the Go programming language and support ttl and you own custom struct.

Methods
=======
```go
type IInMemory interface {
	StoreItem(item IItem, ttlSecond int64)
	ItemByKey(key interface{}) (item interface{}, err error)
	DeleteItemByKey(key interface{})
	CountItems() int64
}
```

Basic usage
=======
```go
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
```

For create custom item struct, you should implement two methods from IItem iterface:
```go
type IItem interface {
	KeyValue() interface{}
	Item() interface{}
}
```

Examples
=======
You can find examples in examples folder.


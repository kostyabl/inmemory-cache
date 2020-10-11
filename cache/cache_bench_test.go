package cache

import (
	"strconv"
	"testing"
)

func prepareBench() {
	var i int64
	for i = 1; i <= itemsCount; i++ {
		items = append(items, getItem(i, strconv.FormatInt(i, 10)))
	}
}

func BenchmarkInMemory_StoreItem(b *testing.B) {
	prepareBench()
	ic := New()
	for i := 0; i < b.N; i++ {
		ic.StoreItem(items[i], 0)
	}
}

func BenchmarkInMemory_StoreItem2(b *testing.B) {
	prepareBench()
	ic := New()
	for i := 0; i < b.N; i++ {
		ic.StoreItem(items[i], 1)
	}
}

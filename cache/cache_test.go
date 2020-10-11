package cache

import (
	"errors"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

const (
	itemsCount = 1000000
)

var (
	items []IItem
)

func prepare(ttl int64) *InMemory {
	items = nil
	var i int64
	for i = 1; i <= itemsCount; i++ {
		items = append(items, getItem(i, strconv.FormatInt(i, 10)))
	}
	ic := New()
	for _, val := range items {
		ic.StoreItem(val, ttl)
	}
	return ic
}

func getItem(key int64, value string) IItem {
	return Item{
		Key:   key,
		Value: value,
	}
}

//TestInMemory_StoreItem - store test
func TestInMemory_StoreItem(t *testing.T) {
	ic := prepare(0)
	iCount := ic.CountItems()
	if iCount != itemsCount {
		t.Errorf("got: %d, expected: %d", iCount, itemsCount)
	}
}

//TestInMemory_StoreItem2 - ttl test
func TestInMemory_StoreItem2(t *testing.T) {
	ic := prepare(1)
	time.Sleep(time.Second * 2)
	iCount := ic.CountItems()
	if iCount != 0 {
		t.Errorf("items have not been deleted, got: %d, expected: 0", iCount)
	}
}

//TestInMemory_ItemByKey - test get item by key
func TestInMemory_ItemByKey(t *testing.T) {
	ic := prepare(0)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	rndKey := r1.Int63n(itemsCount) + 1

	item, err := ic.ItemByKey(rndKey)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	iVal := item.(Item).Value
	rndS := strconv.FormatInt(rndKey, 10)
	if iVal != rndS {
		t.Errorf("got: %s, expected: %s", iVal, rndS)
	}
}

//TestInMemory_DeleteItemByKey - delete item by key
func TestInMemory_DeleteItemByKey(t *testing.T) {
	ic := prepare(0)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	rndKey := r1.Int63n(itemsCount) + 1

	ic.DeleteItemByKey(rndKey)

	_, err := ic.ItemByKey(rndKey)
	if err == nil {
		t.Fatalf("item by key: %d exists", rndKey)
	}
	if !errors.Is(err, ErrItemNotFound) {
		t.Fatalf("item by key: %d, got: %+v, expected: %+v", rndKey, err, ErrItemNotFound)
	}
}

//TestInMemory_CountItems - count items
func TestInMemory_CountItems(t *testing.T) {
	TestInMemory_StoreItem(t)
}

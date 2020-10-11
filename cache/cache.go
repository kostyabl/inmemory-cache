package cache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	//ErrItemNotFound - error if item not found
	ErrItemNotFound = errors.New("item by key not found")
)

//InMemory - in memory main struct
type InMemory struct {
	storage   *sync.Map
	expiredAt *sync.Map
	cnt       ICounter
}

//New - inmemory constructor
func New() (im *InMemory) {
	im = &InMemory{
		storage:   &sync.Map{},
		expiredAt: &sync.Map{},
		cnt:       newCounter(),
	}
	go im.deleteExpiredItems()
	return im
}

//StoreItem - store item with ttl
func (im *InMemory) StoreItem(item IItem, ttlSecond int64) {
	im.storage.Store(item.KeyValue(), item.Item())
	if ttlSecond != 0 {
		im.expiredAt.Store(item.KeyValue(), time.Now().Add(time.Second*time.Duration(ttlSecond)))
	}
	im.cnt.Inc()
}

//ItemByKey - get item by key
func (im *InMemory) ItemByKey(key interface{}) (item interface{}, err error) {
	item, ok := im.storage.Load(key)
	if !ok {
		return item, fmt.Errorf("%w, key: %+v", ErrItemNotFound, key)
	}
	return item, nil
}

//DeleteItemByKey - delete item by key
func (im *InMemory) DeleteItemByKey(key interface{}) {
	im.storage.Delete(key)
	im.expiredAt.Delete(key)
	im.cnt.Dec()
}

func (im *InMemory) CountItems() int64 {
	return im.cnt.Count()
}

//deleteExpiredItems - automatically delete expired items in
func (im *InMemory) deleteExpiredItems() {
	for {
		im.expiredAt.Range(func(key, value interface{}) bool {
			expiredTime := value.(time.Time)
			if time.Now().After(expiredTime) {
				im.DeleteItemByKey(key)
			}
			return true
		})
		time.Sleep(time.Millisecond * 100)
	}
}

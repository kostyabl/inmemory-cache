package cache

//IInMemory - inmemory cache interface
type IInMemory interface {
	StoreItem(item IItem, ttlSecond int64)
	ItemByKey(key interface{}) (item interface{}, err error)
	DeleteItemByKey(key interface{})
	CountItems() int64
}

//IItem - inmemory item interface (all items stored to inmemory cache must confirm this contract)
type IItem interface {
	KeyValue() interface{}
	Item() interface{}
}

package cache

//Item - simple item struct
type Item struct {
	Key   int64
	Value string
}

func (i Item) KeyValue() interface{} {
	return i.Key
}

func (i Item) Item() interface{} {
	return i
}

package cache

//ICounter - counter interface
type ICounter interface {
	Inc()
	Dec()
	Count() int64
}

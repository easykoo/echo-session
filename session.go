package session

type Session interface {
	Set(key, val interface{}) error
	Get(key interface{}) interface{}
	Del(key interface{}) error
	Flush() error
}

type Store interface {
}

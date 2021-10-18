package store

type Storage interface {
	GetKey(key string) (value []byte, err error)
	SetKey(key string, value []byte) (err error)
}

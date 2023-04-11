package main

// LocalBigCache 本地缓存接口
type LocalBigCache interface {
	Set(key string, entry []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
}

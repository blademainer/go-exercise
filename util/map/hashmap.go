package _map

type HashMap interface {
	Put(key Entity, value interface{})
	Get(key Entity) interface{}
	Remove(key Entity) interface{}
}

type Entity interface {
	HashCode() uint64
	Equals(entity Entity) bool
}

package types

type InitFunc[T any] func() (T, error)

type BackendFactory interface {
	Create() (interface{}, error)
}

type Backend[T any] interface {
	Do() (T, error)
}

var backendMap = make(map[string]interface{})

func RegisterBackend(name string, factory interface{}) {
	backendMap[name] = factory
}

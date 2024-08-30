package types

type API[T any] interface {
	Do() (T, error)
}

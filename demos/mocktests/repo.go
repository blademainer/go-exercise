package mock

type UserRepository interface {
	Get(id int) (*User, error)
}

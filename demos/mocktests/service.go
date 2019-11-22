package mock

type UserService interface {
	Hello(id int) (*User, error)
}

type User struct {
	id   int
	name string
	age  int
}

type aImpl struct {
	UserRepository
}

func (a *aImpl) Hello(id int) (*User, error) {
	s, e := a.UserRepository.Get(id)
	if e != nil {
		return nil, e
	}
	return s, e
}

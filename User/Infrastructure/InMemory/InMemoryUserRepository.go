package InMemoryUserRepository

import User "mymodule/User/Domain"

type InMemoryUserRepository struct {
	users []User.User
}

func New() InMemoryUserRepository {
	return InMemoryUserRepository{users: []User.User{}}
}

func (self InMemoryUserRepository) Save(user User.User) error {
	self.users = append(self.users, user)
	return nil
}

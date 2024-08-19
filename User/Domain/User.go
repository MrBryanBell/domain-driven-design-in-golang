/*
	✌️USER NAME
	-[x] Name cannot be empty
	-[x] Name should have a minimum 3 characters
	-[x] Name should have a maximum 50 characters

	✌️USER EMAIL
	-[x] Email cannot be empty
	-[x] Email should be gmail
*/

/*
	MAKE IT WORK,
	MAKE IT RIGHT,
	MAKE IT FAST
*/

package User

import (
	"mymodule/User/Domain/UserEmail"
	"mymodule/User/Domain/UserName"
)

type User struct {
	name  UserName.UserName
	email UserEmail.UserEmail
}

func Create(name, email string) (User, error) {

	_name, error := UserName.New(name)
	if error != nil {
		return User{}, error
	}

	_email, error := UserEmail.New(email)
	if error != nil {
		return User{}, error
	}

	return User{
		name:  _name,
		email: _email,
	}, nil

}

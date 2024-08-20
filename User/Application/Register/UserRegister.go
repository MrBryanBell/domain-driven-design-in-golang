package UserRegister

import User "mymodule/User/Domain"

type UserRegister struct {
	repository User.UserRepository
}

func Setup(repository User.UserRepository) UserRegister {
	return UserRegister{repository: repository}
}

func (self UserRegister) Run(name, email string) error {

	var user, error = User.Create(name, email)
	if error != nil {
		return error
	}

	error = self.repository.Save(user)
	if error != nil {
		return error
	}

	return nil
}

/*
	GREAT PROGRAMMERS START FROM A DRAFT OF HOW THE CLIENT
	IS GOING TO USE THEIR CODE

	var register = UserRegister.New(repository)
	var register = UserRegister.Setup(repository)

	register.register()
	register.run()
*/

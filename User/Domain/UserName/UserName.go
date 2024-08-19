package UserName

import "fmt"

const MIN_LENGTH = 3
const MAX_LENGTH = 50

type UserName struct {
	value string
}

func New(name string) (UserName, error) {
	if len(name) < MIN_LENGTH {
		return UserName{}, fmt.Errorf("name should have a minimum %d characters", MIN_LENGTH)
	}

	if len(name) > MAX_LENGTH {
		return UserName{}, fmt.Errorf("name should have a maximum %d characters", MAX_LENGTH)
	}

	return UserName{value: name}, nil
}

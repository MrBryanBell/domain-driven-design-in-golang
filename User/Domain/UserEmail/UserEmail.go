package UserEmail

import (
	"fmt"
	"strings"
)

const VALID_DOMAIN = "gmail.com"

type UserEmail struct {
	value string
}

func New(email string) (UserEmail, error) {
	if !strings.Contains(email, VALID_DOMAIN) {
		return UserEmail{}, fmt.Errorf("email should be gmail")
	}

	return UserEmail{value: email}, nil
}

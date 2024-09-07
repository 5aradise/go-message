package valid

import (
	"errors"
	"regexp"
)

var emailRe = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

func Email(email string) error {
	if !emailRe.MatchString(email) {
		return errors.New("invalide email")
	}
	return nil
}

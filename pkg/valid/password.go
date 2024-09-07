package valid

import (
	"errors"
	"fmt"
	"regexp"
)

var MinPasswordLen = 1
var MaxPasswordLen = 64

var passwordRe = regexp.MustCompile(`^[a-zA-Z0-9\p{P}\p{S}]+$`)

func Password(password string) error {
	l := len(password)
	if l < MinPasswordLen {
		return fmt.Errorf("password must have at least %d character", MinPasswordLen)
	}
	if l > MaxPasswordLen {
		return fmt.Errorf("password must have no more than %d characters", MaxNameLen)
	}
	if !passwordRe.MatchString(password) {
		return errors.New("password must contains only english letters numbers and special characters")
	}
	return nil
}

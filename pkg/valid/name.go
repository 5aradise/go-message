package valid

import (
	"errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

var MinNameLen = 3
var MaxNameLen = 15

var nameRe = regexp.MustCompile(`^[\p{L}\p{N}_-]+$`)

func Name(name string) error {
	l := utf8.RuneCountInString(name)
	if l < MinNameLen {
		return fmt.Errorf("name must have at least %d characters", MinNameLen)
	}
	if l > MaxNameLen {
		return fmt.Errorf("name must have no more than %d characters", MaxNameLen)
	}
	if !nameRe.MatchString(name) {
		return errors.New("name must contains only letters, numbers and '_', '-'")
	}
	return nil
}

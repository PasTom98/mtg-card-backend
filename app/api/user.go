package api

import (
	"errors"
	"regexp"
	"strings"
)

type User struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

const (
	minNameLength = 2
	maxNameLength = 50
)

var (
	ErrNameTooShort    = errors.New("name too short")
	ErrNameTooLong     = errors.New("name too long")
	ErrNameInvalidChar = errors.New("name contains invalid characters")
)

var validNameRegex = regexp.MustCompile("^[a-zA-Z0-9]+$")

func (u *User) validateName(name string) error {
	name = strings.TrimSpace(name)

	switch {
	case len(name) < minNameLength:
		return ErrNameTooShort
	case len(name) > maxNameLength:
		return ErrNameTooLong
	case !validNameRegex.MatchString(name):
		return ErrNameInvalidChar
	}
	return nil
}

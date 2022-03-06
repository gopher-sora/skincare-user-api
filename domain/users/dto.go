package users

import (
	"strings"

	"github.com/gopher-sora/skincare-user-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FristName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}

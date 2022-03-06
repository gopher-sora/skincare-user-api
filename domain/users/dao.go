package users

import (
	"fmt"

	"github.com/gopher-sora/skincare-user-api/database/mysql/users"
	"github.com/gopher-sora/skincare-user-api/utils/errors"
)

const (
	INSERT_USER = "INSERT INTO users VALUES(?,?,?,?);"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users.Client.Prepare(INSERT_USER)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()),
		)
	}

	return nil
}

func (user *User) Get() *errors.RestErr {
	if err := users.UserDB.Ping(); err != nil {
		panic(err)
	}

	return nil, nil
}

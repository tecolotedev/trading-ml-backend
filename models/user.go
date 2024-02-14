package models

import (
	"context"
	"fmt"

	"github.com/tecolotedev/trading-ml-backend/db"
	sqlc "github.com/tecolotedev/trading-ml-backend/sqlc/code"
	"github.com/tecolotedev/trading-ml-backend/utils"
)

var pack = "models"

type User struct {
	sqlc.User
}

func (u *User) CreateUser() (sqlc.CreateUserRow, error) {

	newUser := sqlc.CreateUserRow{}

	err := u.hashPassword()

	// validate email

	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return newUser, fmt.Errorf("Error creating user please try it later")
	}

	params := sqlc.CreateUserParams{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}

	newUser, err = db.Queries.CreateUser(context.Background(), params)

	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return newUser, err
	}

	return newUser, nil

}

func (u *User) hashPassword() error {

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	return nil
}

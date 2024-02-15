package models

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"

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

	// hash password
	hashedPassword, err := utils.HashPassword(u.Password)
	u.Password = hashedPassword
	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return newUser, fmt.Errorf("error creating user please try it later")
	}

	// TODO:  validate email

	params := sqlc.CreateUserParams{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}

	newUser, err = db.Queries.CreateUser(context.Background(), params)

	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return newUser, fmt.Errorf("error creating user please try it later")
	}

	return newUser, nil

}

func (u *User) GetByID(id int32) error {

	user := User{}

	dbUser, err := db.Queries.GetUserById(context.Background(), id)
	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return fmt.Errorf("error getting user")
	}

	// copy matching field from dbUser to user type
	copier.Copy(&user, &dbUser)

	// update user
	*u = user

	return nil
}

func (u *User) GetByEmail(email string) error {
	user := User{}

	dbUser, err := db.Queries.GetUser(context.Background(), email)
	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return fmt.Errorf("wrong credentials") //wrong email
	}

	// copy matching field from dbUser to user type
	copier.Copy(&user, &dbUser)

	*u = user
	return nil
}

func (u *User) ValidateLogin(password string) error {
	if !u.Verified.Bool {
		return fmt.Errorf("user not verified")
	}

	if err := utils.CheckPassword(password, u.Password); err != nil {
		utils.Log.ErrorLog(err, pack)
		return fmt.Errorf("wrong credentials") // wrong password
	}
	return nil
}

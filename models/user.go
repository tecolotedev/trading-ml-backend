package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jinzhu/copier"

	"github.com/tecolotedev/trading-ml-backend/db"
	sqlc "github.com/tecolotedev/trading-ml-backend/sqlc/code"
	"github.com/tecolotedev/trading-ml-backend/utils"
)

var pack = "models"

type User struct {
	sqlc.User
}
type SafeUser struct {
	ID          int32
	Name        string
	LastName    string
	Username    string
	Email       string
	Verified    pgtype.Bool
	CreatedAt   pgtype.Timestamp
	LastUpdated pgtype.Timestamp
	PlanID      pgtype.Int4
}

type CreateUserInput struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
	Name     string `json:"name"  form:"name"`
	LastName string `json:"last_name"  form:"last_name"`
}

func (u *User) CreateUser(input CreateUserInput) (output SafeUser, err error) {

	if input.Name == "" || input.LastName == "" || input.Username == "" || input.Email == "" || input.Password == "" {
		return output, fmt.Errorf("username, email, password, name and last_name can't be null")
	}

	// TODO:  validate email

	// hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return output, fmt.Errorf("error creating user please try it later")
	}

	params := sqlc.CreateUserParams{
		Username: input.Username,
		Password: hashedPassword,
		Email:    input.Email,
		Name:     input.Name,
		LastName: input.LastName,
	}

	dbUser, err := db.Queries.CreateUser(context.Background(), params)

	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return output, fmt.Errorf("error creating user please try it later")
	}

	copier.Copy(&output, &dbUser)

	return output, nil

}

type LoginInput struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
}

func (u *User) Login(input LoginInput) (output SafeUser, err error) {

	if err = u.GetByEmail(input.Email); err != nil {
		return output, err
	}

	if err = u.ValidateLogin(input.Password); err != nil {
		return output, err
	}

	copier.Copy(&output, u)

	return

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

	dbUser, err := db.Queries.GetUserByEmail(context.Background(), email)
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

type UpdateUserInput struct {
	Name     string `json:"name" form:"name"`
	LastName string `json:"last_name" form:"last_name"`
	Username string `json:"username" form:"username"`
	Password string `json:"password"  form:"password"`
	Email    string `json:"email" form:"email"`
	PlanID   int    `json:"plan_id" form:"plan_id"`
	ID       int32
}

func (u *User) UpdateUser(input UpdateUserInput) (err error) {

	err = u.GetByID(input.ID)
	if err != nil {
		return fmt.Errorf("error getting user")
	}

	params := sqlc.UpdateUserParams{
		Name:     u.Name,
		LastName: u.LastName,
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		PlanID:   u.PlanID,
		ID:       u.ID,
	}

	if input.Name != "" {
		params.Name = input.Name
	}
	if input.LastName != "" {
		params.LastName = input.LastName
	}
	if input.Username != "" {
		params.Username = input.Username
	}
	if input.Password != "" {
		params.Password = input.Password
	}
	if input.Email != "" {
		params.Email = input.Email
	}
	if input.PlanID != 0 {
		params.PlanID = pgtype.Int4{Int32: int32(input.ID), Valid: true}
	}

	err = db.Queries.UpdateUser(context.Background(), params)
	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return fmt.Errorf("error getting user")
	}

	return

}

type DeleteUserInput struct {
	ID int32
}

func (u *User) DeleteUser(input DeleteUserInput) (err error) {

	err = db.Queries.DeleteUser(context.Background(), input.ID)
	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return fmt.Errorf("error deleting user")
	}

	return nil
}

func (u *User) GetUserPlanById(id int32) (output sqlc.GetUserPlanByIdRow, err error) {
	userPlan, err := db.Queries.GetUserPlanById(context.Background(), id)
	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return userPlan, fmt.Errorf("error getting user")
	}
	return userPlan, nil
}

package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/db"
	"github.com/tecolotedev/stori_back/db/sqlc_code"
	"github.com/tecolotedev/stori_back/email"
	"github.com/tecolotedev/stori_back/utils"
)

type loginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
}
type loginResponse struct {
	ID        int    `json:"id" form:"id"`
	Email     string `json:"email" form:"email"`
	Username  string `json:"username"  form:"username"`
	CreatedAt string `json:"created_at"  form:"created_at"`
}

func Login(c *fiber.Ctx) error {
	loginBody := new(loginRequest)

	if err := c.BodyParser(loginBody); err != nil {
		return utils.SendError(c, "Error in request", fiber.StatusBadRequest)
	}

	user, err := db.Queries.GetUser(context.Background(), loginBody.Email)
	if err != nil {
		return utils.SendError(c, "Wrong email or password", fiber.StatusBadRequest)
	}

	if !user.Verified.Bool {
		return utils.SendError(c, "Please verify your account", fiber.StatusBadRequest)
	}

	err = utils.CheckPassword(loginBody.Password, user.Password)
	if err != nil {
		return utils.SendError(c, "Wrong email or password", fiber.StatusBadRequest)
	}

	token, err := utils.CreateToken(user.ID, 24*time.Hour)
	if err != nil {
		return utils.SendError(c, "Error processing, please try it later", fiber.StatusInternalServerError)
	}

	userResponse := loginResponse{
		ID:        int(user.ID),
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Time.String(),
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "access_token"
	cookie.Value = token
	cookie.SameSite = "None"
	cookie.Secure = false
	cookie.Domain = ".tecolotedev.com"
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.Cookie(cookie)

	return utils.SendResponse(c, userResponse)

}

func VerifyToken(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")
	payload, err := utils.VerifyToken(accessToken)
	if err != nil {
		return err
	}
	return utils.SendResponse(c, payload)
}

type signupRequest struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
}

func Signup(c *fiber.Ctx) error {
	signupBody := new(signupRequest)

	if err := c.BodyParser(signupBody); err != nil {
		return utils.SendError(c, "Error in request", fiber.StatusBadRequest)
	}

	hashedPassword, err := utils.HashPassword(signupBody.Password)
	if err != nil {
		fmt.Println(err)
		return utils.SendError(c, "Error processing, please try it later", fiber.StatusInternalServerError)
	}

	params := sqlc_code.CreateUserParams{
		Username: signupBody.Username,
		Email:    signupBody.Email,
		Password: hashedPassword,
	}

	userCreated, err := db.Queries.CreateUser(context.Background(), params)
	if err != nil {
		fmt.Println(err)
		return utils.SendError(c, "Email already exists", fiber.StatusBadRequest)
	}

	email.SendSignupEmail(userCreated.Username, userCreated.ID, userCreated.Email)

	return utils.SendResponse(c, userCreated)
}

func VerifyAccount(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id", ""))

	if err != nil {
		return utils.SendError(c, "Wrong id", fiber.StatusBadRequest)
	}
	_, err = db.Queries.VerifyUser(c.Context(), int32(id))

	if err != nil {
		return err
	}

	return utils.SendResponse(c, struct{}{})
}

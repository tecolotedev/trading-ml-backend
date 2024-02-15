package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/trading-ml-backend/db"
	"github.com/tecolotedev/trading-ml-backend/email"
	"github.com/tecolotedev/trading-ml-backend/models"
	"github.com/tecolotedev/trading-ml-backend/utils"
)

type loginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password"  form:"password"`
}
type loginResponse struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Verified  bool   `json:"verified"`
	CreatedAt string `json:"created_at"`
}

func Login(c *fiber.Ctx) error {
	loginBody := new(loginRequest)

	user := models.User{}

	if err := c.BodyParser(loginBody); err != nil {
		return utils.SendError(c, fmt.Errorf("error in body request"), fiber.StatusBadRequest)
	}

	if err := user.GetByEmail(loginBody.Email); err != nil {
		return utils.SendError(c, err, fiber.StatusBadRequest)
	}

	if err := user.ValidateLogin(loginBody.Password); err != nil {
		return utils.SendError(c, err, fiber.StatusBadRequest)
	}

	token, err := utils.CreateToken(user.ID, time.Hour*24)
	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return utils.SendError(c, fmt.Errorf("error login, please try it later"), fiber.StatusInternalServerError)
	}

	userResponse := loginResponse{
		ID:        int(user.ID),
		Email:     user.Email,
		Username:  user.Username,
		Verified:  user.Verified.Bool,
		CreatedAt: user.CreatedAt.Time.String(),
	}

	cookie := fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		SameSite: "None",
		HTTPOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	}

	c.Cookie(&cookie)

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
	// Parsing request from client
	signupBody := new(signupRequest)
	if err := c.BodyParser(signupBody); err != nil {
		utils.Log.ErrorLog(err, pack)
		return utils.SendError(c, err, fiber.StatusBadRequest)
	}

	// Init new User
	user := models.User{}
	user.Username = signupBody.Username
	user.Password = signupBody.Password
	user.Email = signupBody.Email

	// Calling query to create user
	newUser, err := user.CreateUser()
	if err != nil {
		return utils.SendError(c, err, fiber.StatusBadRequest)
	}

	// Sending email concurrently
	mail := email.Mail{
		Template: email.SignupTemplate,
		Data: map[string]any{
			"name": newUser.Username,
			"id":   newUser.ID,
		},
		To: newUser.Email,
	}
	email.Mailer.MailChan <- mail

	return utils.SendResponse(c, newUser)
}

func VerifyAccount(c *fiber.Ctx) error {
	// get id of user from url
	id, err := strconv.Atoi(c.Query("id", ""))
	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return utils.SendError(c, err, fiber.StatusBadRequest)
	}

	// get token of user from url
	token := c.Query("token", "")

	//verify token
	_, err = utils.VerifyToken(token)

	if err != nil {
		// Resend email if token if something is wrong with the token
		user := models.User{}
		err = user.GetByID(int32(id))

		if err != nil {
			return utils.SendError(c, err, fiber.StatusBadRequest)
		}

		// Re-sending email concurrently
		mail := email.Mail{
			Template: email.SignupTemplate,
			Data: map[string]any{
				"name": user.Username,
				"id":   user.ID,
			},
			To: user.Email,
		}
		email.Mailer.MailChan <- mail

		return utils.SendResponse(c, fiber.Map{"Message": "Token expired, resending email"})

	} else {
		// verify user if token is ok
		_, err = db.Queries.VerifyUser(c.Context(), int32(id))

		if err != nil {
			utils.Log.ErrorLog(err, pack)
			return utils.SendError(c, err, fiber.StatusBadRequest)
		}

		return utils.SendResponse(c, struct{}{})
	}

}

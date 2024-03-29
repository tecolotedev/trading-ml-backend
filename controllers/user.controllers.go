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

func Login(c *fiber.Ctx) error {
	loginBody := new(models.LoginInput)

	user := models.User{}

	if err := c.BodyParser(loginBody); err != nil {
		utils.Log.ErrorLog(err, pack)
		return utils.SendError(c, fmt.Errorf("error in body request"), fiber.StatusBadRequest)
	}

	output, err := user.Login(*loginBody)
	if err != nil {
		return utils.SendError(c, err, fiber.StatusBadRequest)

	}

	token, err := utils.CreateToken(user.ID, time.Hour*24)
	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return utils.SendError(c, fmt.Errorf("error login, please try it later"), fiber.StatusInternalServerError)
	}
	cookie := fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		SameSite: "None",
		HTTPOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	}

	c.Cookie(&cookie)

	return utils.SendResponse(c, output)

}

func Signup(c *fiber.Ctx) error {
	// Parsing request from client
	signupBody := new(models.CreateUserInput)
	if err := c.BodyParser(signupBody); err != nil {
		utils.Log.ErrorLog(err, pack)
		return utils.SendError(c, fmt.Errorf("wrong body request"), fiber.StatusBadRequest)
	}

	// Init new User
	user := models.User{}

	// Calling query to create user
	newUser, err := user.CreateUser(*signupBody)
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

func UpdateUser(c *fiber.Ctx) error {
	updateUserBody := new(models.UpdateUserInput)

	if err := c.BodyParser(updateUserBody); err != nil {
		utils.Log.ErrorLog(err, pack)
		return utils.SendError(c, fmt.Errorf("error in body request"), fiber.StatusBadRequest)
	}

	user := models.User{}

	updateUserBody.ID = c.Locals("userID").(int32)

	err := user.UpdateUser(*updateUserBody)

	if err != nil {
		return utils.SendError(c, err, fiber.StatusInternalServerError)
	}
	return utils.SendResponse(c, struct{}{})
}

func DeleteUser(c *fiber.Ctx) error {

	deleteUserBody := new(models.DeleteUserInput)

	if err := c.BodyParser(deleteUserBody); err != nil {
		return utils.SendError(c, fmt.Errorf("error in body request"), fiber.StatusBadRequest)
	}

	user := models.User{}

	deleteUserBody.ID = c.Locals("userID").(int32)

	err := user.DeleteUser(*deleteUserBody)
	if err != nil {
		return utils.SendError(c, err, fiber.StatusInternalServerError)
	}

	return utils.SendResponse(c, struct{}{})
}

func UpdateUserPlan(c *fiber.Ctx) error {
	updateUserPlanBody := new(models.UpdateUserPlanInput)

	if err := c.BodyParser(updateUserPlanBody); err != nil {
		return utils.SendError(c, fmt.Errorf("error in body request"), fiber.StatusBadRequest)
	}

	updateUserPlanBody.ID = c.Locals("userID").(int32)

	user := models.User{}
	err := user.UpdateUserPlan(*updateUserPlanBody)
	if err != nil {
		return utils.SendError(c, err, fiber.StatusInternalServerError)
	}

	return utils.SendResponse(c, struct{}{})

}

package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/trading-ml-backend/models"
	sqlc "github.com/tecolotedev/trading-ml-backend/sqlc/code"
	"github.com/tecolotedev/trading-ml-backend/utils"
)

func FinancialMiddleware(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int32)

	user := models.User{}
	userPlan, _ := user.GetUserPlanById(userID)

	areValid, message := areParametersValid(ParamsToCheck{outputsize: int32(c.QueryInt("outputsize"))}, userPlan)

	if !areValid {
		return utils.SendError(c, fmt.Errorf(message), fiber.StatusBadRequest)
	}

	return c.Next()
}

type ParamsToCheck struct {
	outputsize int32
}

func areParametersValid(inputParams ParamsToCheck, userPlan sqlc.GetUserPlanByIdRow) (areValid bool, message string) {
	if inputParams.outputsize > userPlan.MaxHistoricalBars.Int32 {
		areValid = false
		message = "outputsize are greater than your plan"
		return
	}
	areValid = true
	return
}

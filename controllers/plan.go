package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/trading-ml-backend/models"
	"github.com/tecolotedev/trading-ml-backend/utils"
)

func GetPlans(c *fiber.Ctx) error {
	plan := models.Plan{}
	plans, err := plan.GetAllPlans()

	if err != nil {
		utils.SendError(c, err, fiber.StatusInternalServerError)
	}

	return utils.SendResponse(c, plans)

}

func GetPlan(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return utils.SendError(c, fmt.Errorf("wrong id"), fiber.StatusBadRequest)
	}

	plan := models.Plan{}

	plan.GetPlanById(int32(id))

	return utils.SendResponse(c, plan)
}

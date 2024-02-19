package models

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/tecolotedev/trading-ml-backend/db"
	sqlc "github.com/tecolotedev/trading-ml-backend/sqlc/code"
	"github.com/tecolotedev/trading-ml-backend/utils"
)

type Plan struct {
	sqlc.Plan
}

func (p *Plan) GetPlanById(id int32) (err error) {
	plan := Plan{}
	dbPlan, err := db.Queries.GetPlanById(context.Background(), id)

	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return fmt.Errorf("error getting plan")
	}

	copier.Copy(&plan, &dbPlan)

	*p = plan
	return
}

func (p *Plan) GetAllPlans() ([]sqlc.Plan, error) {
	plans, err := db.Queries.GetAllPlans(context.Background())

	if err != nil {
		utils.Log.ErrorLog(err, pack)
		return plans, fmt.Errorf("error getting plans")
	}

	return plans, nil
}

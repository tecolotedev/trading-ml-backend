package controllers

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tecolotedev/stori_back/db"
	"github.com/tecolotedev/stori_back/db/sqlc_code"
	"github.com/tecolotedev/stori_back/email"
	"github.com/tecolotedev/stori_back/utils"
)

func ListAccounts(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int32)

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return utils.SendError(c, "Error in request", fiber.StatusBadRequest)
	}

	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return utils.SendError(c, "Error in request", fiber.StatusBadRequest)
	}

	params := sqlc_code.ListAccountsParams{
		UserID: pgtype.Int4{Int32: userId, Valid: true},
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	accounts, err := db.Queries.ListAccounts(context.Background(), params)
	if err != nil {
		return utils.SendError(c, "Error processing, please try it later", fiber.StatusInternalServerError)
	}

	return utils.SendResponse(c, accounts)

}

func GetAccount(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int32)

	accountId, err := strconv.Atoi(c.Params("account_id"))
	if err != nil {
		return utils.SendError(c, "Error in request", fiber.StatusBadRequest)
	}

	account, err := db.Queries.GetAccount(context.Background(), int32(accountId))
	if err != nil {
		return utils.SendError(c, "Account doesn't exist", fiber.StatusInternalServerError)
	}

	if account.UserID.Int32 != userId {
		return utils.SendError(c, "User not authorized to access this account", fiber.StatusBadRequest)
	}

	transfers, err := db.Queries.ListTransfers(context.Background(), pgtype.Int4{Int32: int32(accountId), Valid: true})
	if err != nil {
		return utils.SendError(c, "Error in request", fiber.StatusBadRequest)
	}

	return utils.SendResponse(c, fiber.Map{"account": account, "transfers": transfers})

}

type createAccountRequest struct {
	Balance  float64 `json:"balance" form:"balance"`
	Currency string  `json:"currency" form:"currency"`
}

func CreateAccount(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int32)

	createAccountBody := new(createAccountRequest)

	if err := c.BodyParser(createAccountBody); err != nil {
		return utils.SendError(c, "Error in request", fiber.StatusBadRequest)
	}

	params := sqlc_code.CreateAccountParams{
		Balance:  pgtype.Float8{Float64: createAccountBody.Balance, Valid: true},
		Currency: createAccountBody.Currency,
		UserID:   pgtype.Int4{Int32: userId, Valid: true},
	}

	accountCreated, err := db.Queries.CreateAccount(context.Background(), params)
	if err != nil {
		return utils.SendError(c, "Error processing, please try it later", fiber.StatusInternalServerError)
	}

	return utils.SendResponse(c, accountCreated)

}

func UpdateBalanceAccount(c *fiber.Ctx) error {
	ctx := context.Background()
	userId := c.Locals("userId").(int32)

	accountId, err := strconv.Atoi(c.Params("account_id"))
	if err != nil {
		fmt.Println("err1: ", err)
		return utils.SendError(c, "Error in request", fiber.StatusBadRequest)
	}

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("err2: ", err)

		return utils.SendError(c, "Error with the file", fiber.StatusBadRequest)

	}

	open, err := file.Open()
	if err != nil {
		return utils.SendError(c, "Error with the file", fiber.StatusBadRequest)
	}

	csvReader := csv.NewReader(open)

	data, err := csvReader.ReadAll()
	if err != nil {
		return utils.SendError(c, "Error processing, please try it later", fiber.StatusInternalServerError)
	}

	var records []email.Record

	for i, line := range data {
		if i > 0 { // omit header line
			date, err := time.Parse("2006-01-02", line[1])
			if err != nil {
				continue
			}
			amount, err := strconv.ParseFloat(line[2], 64)
			if err != nil {
				continue
			}
			reason := line[3]

			//transaction block
			err = db.MakeTx(context.Background(), func() error {
				createTransferParams := sqlc_code.CreateTransferParams{
					Amount:    amount,
					Reason:    pgtype.Text{String: reason, Valid: true},
					AccountID: pgtype.Int4{Int32: int32(accountId), Valid: true},
					CreatedAt: pgtype.Timestamp{Time: date, Valid: true},
				}
				_, err := db.Queries.CreateTransfer(ctx, createTransferParams)
				if err != nil {
					fmt.Println("err CreateTransfer: ", err)
					return err
				}

				account, err := db.Queries.GetAccountForUpdate(context.Background(), int32(accountId))
				if err != nil {
					log.Println(err)
					return err
				}

				if account.UserID.Int32 != userId {
					return fmt.Errorf("User not authorized to access this account")
				}

				updateAccountParmas := sqlc_code.UpdateAccountParams{
					ID:      int32(accountId),
					Balance: pgtype.Float8{Float64: account.Balance.Float64 + amount, Valid: true},
				}

				_, err = db.Queries.UpdateAccount(ctx, updateAccountParmas)
				if err != nil {
					fmt.Println("err UpdateAccount: ", err)
					return err
				}
				record := email.Record{Date: line[1], Transaction: amount, Reason: reason}
				records = append(records, record)
				return nil
			})
			if err != nil {
				fmt.Println("err MakeTx: ", err)
			}

		}
	}

	account, err := db.Queries.GetAccountForUpdate(context.Background(), int32(accountId))
	if err != nil {
		return utils.SendError(c, "Error in request", fiber.StatusBadRequest)
	}

	transfers, err := db.Queries.ListTransfers(context.Background(), pgtype.Int4{Int32: int32(accountId), Valid: true})
	if err != nil {
		return utils.SendError(c, "Error in request", fiber.StatusBadRequest)
	}

	user, _ := db.Queries.GetUserById(context.Background(), userId)

	summaries := utils.MakeSummary(transfers)

	email.SendReportEmail(user.Email, account.ID, account.Balance.Float64, records, summaries)

	return utils.SendResponse(c, fiber.Map{"account": account, "transfers": transfers})

}

package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/tecolotedev/trading-ml-backend/config"
	sqlc "github.com/tecolotedev/trading-ml-backend/sqlc/code"
)

var Queries sqlc.Queries
var Conn *pgx.Conn

func InitDb() {
	ctx := context.Background()

	sslConn := "allow"

	if config.EnvVars.IS_LOCAL {
		sslConn = "disable"
	}

	configConn := fmt.Sprintf("host=%s password=%s user=%s dbname=%s sslmode=%s",
		config.EnvVars.DB_HOST,
		config.EnvVars.DB_PASSWORD,
		config.EnvVars.DB_USER,
		config.EnvVars.DB_NAME,
		sslConn,
	)

	conn, err := pgx.Connect(ctx, configConn)
	if err != nil {
		log.Fatal(err)
	}

	Queries = *sqlc.New(conn) //sqlc_.New(conn)
	Conn = conn
}

// Postgres transactions to make safetly money's movements
func MakeTx(ctx context.Context, transactions func() error) error {

	tx, err := Conn.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return fmt.Errorf("error init Tx")
	}

	tx.Begin(ctx)

	err = transactions()

	if err != nil {
		if rbError := tx.Rollback(ctx); rbError != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbError)
		}
		return err
	}
	return tx.Commit(ctx)

}

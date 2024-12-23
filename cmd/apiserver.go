package cmd

import (
	"context"
	"os"

	"github.com/greed-verse/greed/internal/account"
	"github.com/greed-verse/greed/internal/shared"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func Execute() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	ctx := context.Background()
	dbConn, err := pgx.Connect(ctx, os.Getenv("DB_URL"))
	if err != nil {
		return err
	}

	appContext := shared.New(dbConn, os.Getenv("APP_ADDRESS"))

	account.InitModule(appContext)

	appContext.Serve()
	return nil
}

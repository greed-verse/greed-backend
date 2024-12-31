package cmd

import (
	"context"
	"fmt"

	"github.com/greed-verse/greed/internal/account"
	"github.com/greed-verse/greed/internal/payment"
	"github.com/greed-verse/greed/internal/shared"
	"github.com/greed-verse/greed/pkg/env"
	"github.com/jackc/pgx/v5"
)

func Execute() error {
	environment := env.GetEnv()
	fmt.Println(env.GetEnv().DB_URL())
	fmt.Println(env.GetEnv().APP_ADDRESS())

	ctx := context.Background()
	dbConn, err := pgx.Connect(ctx, environment.DB_URL())
	if err != nil {
		return err
	}

	appContext := shared.New(dbConn, environment.APP_ADDRESS())

	payment.New(appContext)
	account.New(appContext)

	appContext.Serve()
	return nil
}

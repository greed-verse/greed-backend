package cmd

import (
	"context"

	"github.com/greed-verse/greed/internal/account"
	"github.com/greed-verse/greed/internal/payment"
	"github.com/greed-verse/greed/internal/shared"
	"github.com/greed-verse/greed/pkg/env"
	"github.com/jackc/pgx/v5"
	"github.com/stripe/stripe-go/v81"
)

func Execute() error {
	environment := env.GetEnv()

	stripe.Key = environment.STRIPE_KEY()

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

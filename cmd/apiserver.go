package cmd

import (
	"context"

	"github.com/greed-verse/greed/internal/account"
	"github.com/greed-verse/greed/internal/shared"
	"github.com/greed-verse/greed/pkg/env"
	"github.com/jackc/pgx/v5"
)

func Execute() error {
	environment := env.GetEnv()

	ctx := context.Background()
	dbConn, err := pgx.Connect(ctx, environment.DB_URL())

	if err != nil {
		return err
	}

  appContext := shared.New(dbConn, environment.APP_ADDRESS())

	account.InitModule(appContext)

	appContext.Serve()
	return nil
}

func ResolveEnv(env string) (string, error) {
	filepath, exists := os.LookupEnv(env + "_FILE")

	if exists {
		content, err := os.ReadFile(filepath)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(content)), nil
	}
	return os.Getenv(env), nil
}

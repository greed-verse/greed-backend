package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/greed-verse/greed/internal/account"
	"github.com/greed-verse/greed/internal/shared"
	"github.com/jackc/pgx/v5"
)

func Execute() error {
	url, err := ResolveEnv("DB_URL")
	if err != nil {
		return err
	}
	fmt.Println(url)

	addr, err := ResolveEnv("DB_URL")
	if err != nil {
		return err
	}
	fmt.Println(addr)

	ctx := context.Background()
	dbConn, err := pgx.Connect(ctx, url)
	if err != nil {
		return err
	}

	appContext := shared.New(dbConn, addr)

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

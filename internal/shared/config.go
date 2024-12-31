package shared

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5"
)

type Server interface {
	Serve()
}

type AppContext struct {
	API    *API
	Logger *Logger
	Repo   *pgx.Conn
	PubSub *PubSub

	Addr string
}

func New(db *pgx.Conn, addr string) *AppContext {
	logger := NewLogger()

	api := NewApi(logger)
	api.InitMiddlewares(logger.Middleware())

	pubsub := NewPubSub(logger)
	// Add pubsub middlewares here
	pubsub.InitMiddlewares()

	return &AppContext{
		API:    api,
		Logger: logger,
		Repo:   db,
		PubSub: pubsub,

		Addr: addr,
	}
}

func (app *AppContext) Serve() {
	api := app.API.Core()
	logger := app.Logger.Core()
	pubsub := app.PubSub.Core()

	logger.Info().Str("address", app.Addr).Msg("App running")
	go func() { _ = api.Listen(app.Addr) }()
	go func() {
		err := app.PubSub.Router().Run(context.Background())
		if err != nil {
			logger.Error().Str("error", err.Error()).Msg("Pubsub Failed")
		}
	}()

	var sig os.Signal
	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	sig = <-c                                       // This blocks the main thread until an interrupt is received

	logger.Info().Str("signal", sig.String()).Msg("Signal received")
	logger.Info().Msg("Shutting down app, waiting background process to finish")
	defer logger.Info().Msg("App shutdown")

	_ = api.ShutdownWithContext(context.Background())
	_ = app.PubSub.Router().Close()
	_ = pubsub.Close()
}

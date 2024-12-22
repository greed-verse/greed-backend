package engine

import "github.com/greed-verse/greed/core"

func (e *engine) initCore() {
	logger := core.NewLogger()
	e.Logger = logger

    server := core.NewApi(logger)
    server.InitMiddlewares(logger)

	logger.Core().Info().Msg("Hello World")
}

func (e *engine) initPkg() {
}

package shared

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

type PubSub struct {
	core   *gochannel.GoChannel
	router *message.Router
}

func NewPubSub(l *Logger) *PubSub {
	defer l.Core().Info().Msg("PubSub Initialized")

	logger := watermill.NewStdLoggerWithOut(l.Core().With().Logger(), true, true)
	pubsub := gochannel.NewGoChannel(gochannel.Config{}, logger)

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}
	return &PubSub{
		core:   pubsub,
		router: router,
	}
}

func (ps *PubSub) Core() *gochannel.GoChannel {
	return ps.core
}

func (ps *PubSub) Router() *message.Router {
	return ps.router
}

func (ps *PubSub) InitMiddlewares(mws ...message.HandlerMiddleware) {
	for _, mw := range mws {
		ps.Router().AddMiddleware(mw)
	}
}

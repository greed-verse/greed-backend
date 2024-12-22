package engine

import "github.com/greed-verse/greed/core"

type Engine interface {
	Serve()
}

type engine struct {
	Addr string

	API    *core.API
	Logger *core.Logger
}

func New(addr string) Engine {
	var engine = new(engine)

	engine.Addr = addr
	engine.initCore()
	engine.initPkg()
	return engine
}

func (e *engine) Serve() {

}

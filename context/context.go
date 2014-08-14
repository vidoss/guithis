package context

import (
	"appengine"
	"github.com/unrolled/render"
	"log"
	"path/filepath"
)

type AppContext struct {
	Render     *render.Render
	GaeContext appengine.Context
}

func NewAppContext(options render.Options) *AppContext {
	log.Println(filepath.Abs("../templates"))
	c := AppContext{
		Render: render.New(options),
	}
	return &c
}

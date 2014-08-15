package views

import (
	"github.com/vidoss/guithis/context"
	"net/http"
)

func HomeHandler(c *context.AppContext, w http.ResponseWriter, r *http.Request) {
	c.Render.HTML(w, http.StatusOK, "home", nil)
}

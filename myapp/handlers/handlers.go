package handlers

import (
	"net/http"

	"github.com/RedHoodJT1988/velocidad"
)

type Handlers struct {
	App *velocidad.Velocidad
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}

}

package http

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *handler) CreateHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "create host!\n")
}

func (h *handler) QueryHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "query host!\n")
}

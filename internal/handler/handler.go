package handler

import (
	"fmt"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Letsgo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "let's go")
}

package middleware

import "net/http"

type Handler struct {
	http.Handler
}

func NewHandler() *Handler {
	return &Handler{}
}

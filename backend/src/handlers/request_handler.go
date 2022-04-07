package handlers

import (
    "net/http"
    "github.com/go-chi/chi"
    "github.com/go-chi/render"
)

func NewHandler() http.Handler {
    router := chi.NewRouter()
    router.MethodNotAllowed(methodNotAllowedHandler)
    router.NotFound(notFoundHandler)

    return router
}

func methodNotAllowedHandler(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-type", "application/json")
    writer.WriteHeader(405)
    render.Render(writer, request, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    w.WriteHeader(400)
    render.Render(w, r, ErrNotFound)
}

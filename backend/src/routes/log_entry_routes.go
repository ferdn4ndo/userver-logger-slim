package routes

import (
    "fmt"

    "github.com/go-chi/chi/v5"

    "github.com/ferdn4ndo/userver-logger-slim/backend/controllers"
)

func LogEntryRouter() chi.Router {
    router := chi.NewRouter()

    router.Get("/", controllers.GetListLogEntry)
    router.Post("/", controllers.PostLogEntry)
    router.Route(fmt.Sprintf("/{%s}", controllers.PARAM_LOG_ENTRY_ID), func(router chi.Router) {
        router.Use(controllers.LogEntryContext)
        router.Get("/", controllers.GetLogEntry)
        router.Put("/", controllers.PutLogEntry)
        router.Delete("/", controllers.DeleteLogEntry)
    })

    return router
}

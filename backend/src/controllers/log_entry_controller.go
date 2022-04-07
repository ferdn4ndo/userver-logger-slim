package controllers

import (
    "context"
    "fmt"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/render"

    "github.com/ferdn4ndo/userver-logger-slim/backend/models"
    "github.com/ferdn4ndo/userver-logger-slim/backend/handlers"
    "github.com/ferdn4ndo/userver-logger-slim/backend/services"
)

var PARAM_LOG_ENTRY_ID = "logEntryId"
var CONTEXT_LOG_ENTRY = "contextLogEntry"

// Param converter (context) for the "logEntryId"
func LogEntryContext(next http.Handler) http.Handler {
    return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
        logEntryId := chi.URLParam(request, "logEntryId")
        if logEntryId == "" {
            render.Render(writer, request, handlers.ErrorRenderer(fmt.Errorf("Log entry ID is required.")))

            return
        }
        
        id, error := strconv.Atoi(logEntryId)
        if error != nil {
            render.Render(writer, request, handlers.ErrorRenderer(fmt.Errorf("Invalid log entry ID.")))

            return
        }

        logEntryPointer, error := services.GetLogEntryById(uint(id))
        if error != nil {
            if error == services.ErrNoMatch {
                render.Render(writer, request, handlers.ErrNotFound)
            } else {
                render.Render(writer, request, handlers.ErrorRenderer(error))
            }

            return
        }

        fmt.Printf("VALUE: %+v \n", logEntryPointer)
        fmt.Printf("VALUE2: %+v \n", *logEntryPointer)

        ctx := context.WithValue(request.Context(), CONTEXT_LOG_ENTRY, logEntryPointer)

        next.ServeHTTP(writer, request.WithContext(ctx))
    })
}

// GET /log-entries
func GetListLogEntry(writer http.ResponseWriter, request *http.Request) {
    logEntryListPointer, error := services.GetAllLogEntries()
    if error != nil {
        render.Render(writer, request, handlers.ServerErrorRenderer(error))

        return
    }

    error = render.RenderList(writer, request, services.NewLogEntryListResponse(logEntryListPointer))
    if error != nil {
        render.Render(writer, request, handlers.ErrorRenderer(error))
    }
}

// POST /log-entries
func PostLogEntry(writer http.ResponseWriter, request *http.Request) {
    logEntryRequest := &models.LogEntryRequest{}

    if error := render.Bind(request, logEntryRequest); error != nil {
        render.Render(writer, request, handlers.ErrBadRequest)

        return
    }

    if error := services.AddLogEntry(logEntryRequest.LogEntry); error != nil {
        render.Render(writer, request, handlers.ErrorRenderer(error))

        return
    }

    render.Status(request, http.StatusCreated)
    error := render.Render(writer, request, services.NewLogEntryResponse(logEntryRequest.LogEntry))
    if error != nil {
        render.Render(writer, request, handlers.ServerErrorRenderer(error))

        return
    }
}

// GET /log-entries/{id}
func GetLogEntry(writer http.ResponseWriter, request *http.Request) {
    logEntryPointer := request.Context().Value(CONTEXT_LOG_ENTRY).(*models.LogEntry)

    error := render.Render(writer, request, services.NewLogEntryResponse(logEntryPointer))
    if error != nil {
        render.Render(writer, request, handlers.ServerErrorRenderer(error))

        return
    }
}

// PATCH /log-entries/{id}
func PutLogEntry(writer http.ResponseWriter, request *http.Request) {
    logEntryPointer := request.Context().Value(CONTEXT_LOG_ENTRY).(*models.LogEntry)
    logEntryRequest := &models.LogEntryRequest{}

    if error := render.Bind(request, logEntryRequest); error != nil {
        render.Render(writer, request, handlers.ErrBadRequest)

        return
    }

    updatedLogEntryPointer, error := services.UpdateLogEntry((*logEntryPointer).ID, logEntryRequest)
    if error != nil {
        render.Render(writer, request, handlers.ServerErrorRenderer(error))

        return
    }

    error = render.Render(writer, request, services.NewLogEntryResponse(updatedLogEntryPointer))
    if error != nil {
        render.Render(writer, request, handlers.ServerErrorRenderer(error))

        return
    }
}

// DELETE /log-entries/{id}
func DeleteLogEntry(writer http.ResponseWriter, request *http.Request) {
    logEntry := request.Context().Value(CONTEXT_LOG_ENTRY).(*models.LogEntry)
    
    error := services.DeleteLogEntry(logEntry.ID)
    if error != nil {
        render.Render(writer, request, handlers.ServerErrorRenderer(error))

        return
    }
    
    writer.WriteHeader(http.StatusNoContent)
}

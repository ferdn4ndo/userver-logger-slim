package services

import (
    "github.com/go-chi/render"

    "github.com/ferdn4ndo/userver-logger-slim/backend/models"
)

func NewLogEntryResponse(logEntry *models.LogEntry) *models.LogEntryResponse {
    resp := &models.LogEntryResponse{LogEntry: logEntry}

    return resp
}

func NewLogEntryListResponse(logEntries *models.LogEntryList) []render.Renderer {
    list := []render.Renderer{}
    for _, logEntry := range logEntries.LogEntries {
        list = append(list, NewLogEntryResponse(logEntry))
    }
    return list
}

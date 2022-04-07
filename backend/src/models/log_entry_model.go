package models

import (
    "fmt"
    "net/http"
    "time"
)

type LogEntry struct {
    ID uint `gorm:"primaryKey" json:"id"`
    Producer string `json:"producer"`
    Message string `gorm:"type:text" json:"message"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type LogEntryList struct {
    LogEntries []*LogEntry `json:"log_entries" gorm:"-"`
}

func (*LogEntryList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}

func (logEntry *LogEntry) Bind(r *http.Request) error {
    if logEntry.Producer == "" {
        return fmt.Errorf("The field 'producer' is required.")
    }

    if logEntry.Message == "" {
        return fmt.Errorf("The field 'message' is required.")
    }

    return nil
}

func (*LogEntry) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}

type LogEntryResponse struct {
    *LogEntry

    // We could add an additional field to the response here, such as this
    // elapsed computed property:
    //Elapsed int64 `json:"elapsed"`
}

func (logEntry *LogEntryResponse) Render(writer http.ResponseWriter, request *http.Request) error {
    // Pre-processing before a response is marshalled and sent across the wire
    //logEntry.Elapsed = 10
    return nil
}

// LogEntryRequest is the request payload for LogEntry data model.
//
// NOTE: It's good practice to have well defined request and response payloads
// so you can manage the specific inputs and outputs for clients, and also gives
// you the opportunity to transform data on input or output, for example
// on request, we'd like to protect certain fields and on output perhaps
// we'd like to include a computed field based on other values that aren't
// in the data model. Also, check out this awesome blog post on struct composition:
// http://attilaolah.eu/2014/09/10/json-and-struct-composition-in-go/
type LogEntryRequest struct {
    *LogEntry

    ProtectedID string `json:"id"` // override 'id' json to have more control
}

func (logEntryRequest *LogEntryRequest) Bind(request *http.Request) error {
    if logEntryRequest.Producer == "" {
        return fmt.Errorf("The field 'producer' is required.")
    }

    if logEntryRequest.Message == "" {
        return fmt.Errorf("The field 'message' is required.")
    }

    // just a post-process after a decode..
    logEntryRequest.ProtectedID = "" // unset the protected ID
    //a.Article.Title = strings.ToLower(a.Article.Title) // as an example, we down-case
    return nil
}

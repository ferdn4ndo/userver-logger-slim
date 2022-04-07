package handlers

import (
    "fmt"
    "net/http"

    "github.com/go-chi/render"
)

type ErrorResponse struct {
    Err error `json:"-"`
    StatusCode int `json:"-"`
    StatusText string `json:"status_text"`
    Message string `json:"message"`
}

var (
    ErrMethodNotAllowed = &ErrorResponse{
        StatusCode: http.StatusMethodNotAllowed,
        StatusText: "Method not allowed.",
        Message: "Method not allowed. Check the documentation for the correct endpoints.",
    }
    ErrNotFound = &ErrorResponse{
        StatusCode: http.StatusNotFound,
        StatusText: "Resource not found.",
        Message: "Resource not found. Check the given identifier and try again.",
    }
    ErrBadRequest = &ErrorResponse{
        StatusCode: http.StatusBadRequest,
        Message: "Bad request. Check the data sent in the request an try again.",
     }
)

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
    render.Status(r, e.StatusCode)
    return nil
}

func ErrorRenderer(err error) *ErrorResponse {
    return &ErrorResponse{
        Err: err,
        StatusCode: 400,
        StatusText: "There was a bad request.",
        Message: err.Error(),
    }
}

func ServerErrorRenderer(err error) *ErrorResponse {
    return &ErrorResponse{
        Err: err,
        StatusCode: 500,
        StatusText: "An internal server error has occurred.",
        Message: err.Error(),
    }
}

func AddCustomErrorHandlerfunc(writer http.ResponseWriter, request *http.Request, value interface{}) {
    if err, ok := value.(error); ok {

        // We set a default error status response code if one hasn't been set.
        if _, ok := request.Context().Value(render.StatusCtxKey).(int); !ok {
            writer.WriteHeader(400)
        }

        // We log the error
        fmt.Printf("Logging err: %s\n", err.Error())

        // We change the response to not reveal the actual error message,
        // instead we can transform the message something more friendly or mapped
        // to some code / language, etc.
        render.DefaultResponder(writer, request, render.M{"status": "error"})
        return
    }

    render.DefaultResponder(writer, request, value)
}

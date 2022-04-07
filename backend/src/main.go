package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/go-chi/render"

    "github.com/ferdn4ndo/userver-logger-slim/backend/handlers"
    "github.com/ferdn4ndo/userver-logger-slim/backend/models"
    "github.com/ferdn4ndo/userver-logger-slim/backend/routes"
    "github.com/ferdn4ndo/userver-logger-slim/backend/services"
)

func main() {
    fmt.Println("Initializing uServer Logger backend database...")
    db, err := services.InitializeDatabase()
    if err != nil {
        panic(fmt.Sprintf("Error initializing database: %s", err))
    }

    fmt.Println("Migrating the schema...")
    db.Conn.AutoMigrate(&models.LogEntry{})

    StartServer()

    fmt.Println("The uServer Logger Backend service has started!")
}

// This is entirely optional, but I wanted to demonstrate how you could easily
// add your own logic to the render.Respond method.
func init() {
    render.Respond = handlers.AddCustomErrorHandlerfunc
}

func StartServer() {
    address := ":8080"
    listener, error := net.Listen("tcp", address)
    if error != nil {
        log.Fatalf("Error occurred: %s", error.Error())
    }

    database, error := services.InitializeDatabase()
    if error != nil {
        log.Fatalf("Could not set up database: %v", error)
    }
    defer database.Close()

    //httpHandler := handlers.NewHandler()
    server := &http.Server{
        Handler: routes.CreateRouter(),
    }
    go func() {
        server.Serve(listener)
    }()
    defer StopServer(server)

    log.Printf("Started server on %s", address)
    channel := make(chan os.Signal, 1)
    signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)
    log.Println(fmt.Sprint(<-channel))
    log.Println("Stopping API server.")
}

func StopServer(server *http.Server) {
    context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if error := server.Shutdown(context); error != nil {
        log.Printf("Could not shut down server correctly: %v\n", error)
        os.Exit(1)
    }
}

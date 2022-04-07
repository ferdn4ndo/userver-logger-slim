package services

import (
    "fmt"
    "time"

    "gorm.io/gorm"
    "gorm.io/driver/sqlite"

    "github.com/ferdn4ndo/userver-logger-slim/backend/models"
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("No matching record.")
type DatabaseService struct {
    Conn *gorm.DB
}

func (db *DatabaseService) Close() {
    sqlDB, err := db.Conn.DB()

    if err != nil {
        panic("Error closing DB connection!")
    }

    // Close
    sqlDB.Close()
}


func AddHeartbeatLog(db *DatabaseService) (error) {
    fmt.Println("Adding heartbeat entry")
    currentTime := time.Now().Format(time.RFC3339)
    result := db.Conn.Create(&models.LogEntry{
        Producer: "userver-logger-backend",
        Message: "Heartbeat from userver-logger-backend at " + currentTime})

    return result.Error
}

func InitializeDatabase() (*DatabaseService, error) {
    conn := sqlite.Open("/data/sqlite.db")

    db, err := gorm.Open(conn, &gorm.Config{})
    if err != nil {
        return nil, err
    }

    service := &DatabaseService{Conn: db}

    AddHeartbeatLog(service)

    fmt.Println("Database connection established!")

    return service, nil
}

func GetDatabaseService() (*DatabaseService, error) {
    conn := sqlite.Open("/data/sqlite.db")

    db, err := gorm.Open(conn, &gorm.Config{})
    if err != nil {
        return nil, err
    }

    service := &DatabaseService{Conn: db}

    return service, nil
}

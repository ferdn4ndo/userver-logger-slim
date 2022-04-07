package services

import (
    "gorm.io/gorm"

    "github.com/ferdn4ndo/userver-logger-slim/backend/models"
)

func GetAllLogEntries() (*models.LogEntryList, error) {
    var logEntries models.LogEntryList

    db, error := GetDatabaseService()
    if error != nil {
        return &logEntries, error
    }

    list := db.Conn.Model(&models.LogEntry{}).Find(&logEntries.LogEntries)
    db.Close()

    if list.Error != nil {
        return &logEntries, list.Error
    }

    return &logEntries, nil
}

func AddLogEntry(logEntry *models.LogEntry) error {
    db, error := GetDatabaseService()
    if error != nil {
        return error
    }

    result := db.Conn.Create(&logEntry)
    db.Close()

    return result.Error
}

func CheckIfIdExists(logEntryId uint) (bool, error) {
    var exists bool
    
    db, error := GetDatabaseService()
    if error != nil {
        return exists, error
    }

    error = db.Conn.Model(&models.LogEntry{}).
            Select("count(*) > 0").
            Where("id = ?", logEntryId).
            Find(&exists).Error
    db.Close()

    if error != nil {
        return exists, error
    }

    return exists, nil
}

func GetLogEntryById(logEntryId uint) (*models.LogEntry, error) {
    var logEntry models.LogEntry
    
    db, error := GetDatabaseService()
    if error != nil {
        return &logEntry, error
    }

    error = db.Conn.First(&logEntry, "id = ?", logEntryId).Error
    db.Close()

    if error != nil {
        if error == gorm.ErrRecordNotFound {
            return &logEntry, ErrNoMatch
        } else {
            return &logEntry, error
        }
    }

    return &logEntry, nil
}

func DeleteLogEntry(logEntryId uint) error {
    var logEntry models.LogEntry
    
    db, error := GetDatabaseService()
    if error != nil {
        return error
    }

    error = db.Conn.Delete(&logEntry, logEntryId).Error
    db.Close()

    return error
}

func UpdateLogEntry(logEntryId uint, logEntryData *models.LogEntryRequest) (*models.LogEntry, error) {
    var logEntry models.LogEntry
    
    db, error := GetDatabaseService()
    if error != nil {
        return &logEntry, error
    }

    error = db.Conn.First(&logEntry, "id = ?", logEntryId).Error
    if error != nil {
        return &logEntry, error
    }

    error = db.Conn.Model(&logEntry).Updates(logEntryData.LogEntry).Error
    if error != nil {
        return &logEntry, error
    }

    db.Close()

    return &logEntry, nil
}

#!/bin/sh

echo "Downloading dependencies..."
go get github.com/go-chi/chi/v5 github.com/go-chi/docgen github.com/go-chi/v5/middleware github.com/go-chi/render gorm.io/gorm gorm.io/driver/sqlite github.com/cosmtrek/air

echo "Checking package integrity"
go mod verify

echo "Starting AIR"
air

# Based on the sources:
# - https://techinscribed.com/5-ways-to-live-reloading-go-applications/
# - https://github.com/cosmtrek/air

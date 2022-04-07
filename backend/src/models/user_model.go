package models

import (
    "net/http"
    "strings"
    "time"
)

type User struct {
    ID int `json:"id"`
    Username string `gorm "unique" json:"username" binding:"required"`
    HashedPassword string `json:"-"`
    Email string `gorm:"unique" json:"email" binding:"required"`
    IsAdmin bool `json:"is_admin" binding:"required"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type UserList struct {
    Users []*User `json:"users" gorm:"-"`
}

func (*UserList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}

func (user *User) Bind(r *http.Request) error {
    return nil
}

func (*User) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}

// UserResponse is the response payload for User data model.
type UserResponse struct {
    *User

    // We could add an additional field to the response here, such as this
    // elapsed computed property:
    //Elapsed int64 `json:"elapsed"`
}

func (logEntry *UserResponse) Render(writer http.ResponseWriter, request *http.Request) error {
    // Pre-processing before a response is marshalled and sent across the wire
    //logEntry.Elapsed = 10
    return nil
}

// UserRequest is the request payload for User data model.
type UserRequest struct {
    *User

    Password string `json:"password" binding:"required"`
    ConfirmPassword string `json:"confirm_password" binding:"required"`

    ProtectedID string `json:"id"` // override 'id' json to have more control
}

func (userRequest *UserRequest) Bind(request *http.Request) error {
    // just a post-process after a decode..
    userRequest.ProtectedID = "" // unset the protected ID
    userRequest.User.Email = strings.ToLower(userRequest.User.Email)

    return nil
}

package main

import (
	"net/http"
	"time"

	"github.com/satori/go.uuid"
)

// TmpUsers tmp user data
// when server stop, to be disappear
var tmpUsers = map[string]User{}

// User user data struct
type User struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"update_at"`
}

// Users type user list
type Users []User

//RegistUser regist user data
func RegistUser(user User) (User, error) {

	date := time.Now()
	user.CreatedAt = date
	user.UpdatedAt = date

	id, err := uuid.NewV4()
	if err != nil {
		return user, NewError(http.StatusInternalServerError, err.Error())
	}
	user.ID = id.String()

	tmpUsers[user.ID] = user

	return user, nil
}

// ReadUser get user data
func ReadUser(id string) (User, error) {
	user, exist := tmpUsers[id]

	if !exist {
		return user, NewError(http.StatusNotFound, "not found user")
	}
	return user, nil
}

// ListUser get user list
func ListUser() (Users, error) {
	users := Users{}

	for _, user := range tmpUsers {
		users = append(users, user)
	}

	return users, nil
}

// UpdateUser update user info
func UpdateUser(newUser User) (User, error) {
	if _, err := ReadUser(newUser.ID); err != nil {
		return newUser, err
	}

	date := time.Now()
	newUser.CreatedAt = date
	newUser.UpdatedAt = date

	tmpUsers[newUser.ID] = newUser
	return newUser, nil
}

// DeleteUser delete user data
func DeleteUser(id string) error {
	if _, err := ReadUser(id); err != nil {
		return err
	}

	delete(tmpUsers, id)
	return nil
}

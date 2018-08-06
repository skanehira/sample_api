package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

// uuid v4 regexp
var validUUID = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

// UserHandler handle user request
func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		Register(w, r)
	case http.MethodGet:
		Reader(w, r)
	case http.MethodPut:
		Updater(w, r)
	case http.MethodDelete:
		Deleter(w, r)
	default:
		NotFoundResources(w, r)
	}
}

// Register regist user info
func Register(w http.ResponseWriter, r *http.Request) {
	if err := checkRequestHeader(r); err != nil {
		newHTTPError(w, err)
		return
	}

	newUser := User{}
	err := requestBodyToStruct(r, &newUser)
	if err != nil {
		newHTTPError(w, err)
		return
	}
	user, err := RegistUser(newUser)

	if err != nil {
		newHTTPError(w, err)
		return
	}

	newHTTPResponse(w, http.StatusOK, user)
}

// Updater update user info
func Updater(w http.ResponseWriter, r *http.Request) {
	if err := checkRequestHeader(r); err != nil {
		newHTTPError(w, err)
		return
	}

	id, err := getUserID(r)
	if err != nil {
		newHTTPError(w, err)
		return
	}
	newUser := User{}
	err = requestBodyToStruct(r, &newUser)
	newUser.ID = id

	if err != nil {
		newHTTPError(w, err)
		return
	}

	newUser, err = UpdateUser(newUser)

	if err != nil {
		newHTTPError(w, err)
		return
	}

	newHTTPResponse(w, http.StatusOK, newUser)
}

// Reader get user info
func Reader(w http.ResponseWriter, r *http.Request) {

	id, err := getUserID(r)
	if err != nil {
		newHTTPError(w, err)
		return
	}

	if id != "" {
		user, err := ReadUser(id)

		if err != nil {
			newHTTPError(w, err)
			return
		}

		newHTTPResponse(w, http.StatusOK, user)
		return
	}

	// if not specified user id, get all user
	users, err := ListUser()
	if err != nil {
		newHTTPError(w, err)
		return
	}

	newHTTPResponse(w, http.StatusOK, users)
}

// Deleter delete user info
func Deleter(w http.ResponseWriter, r *http.Request) {
	id, err := getUserID(r)

	if err != nil {
		newHTTPError(w, err)
		return
	}

	err = DeleteUser(id)
	if err != nil {
		newHTTPError(w, err)
		return
	}

	successMessage := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		http.StatusOK,
		fmt.Sprintf("Success delete user=[%s]", id),
	}

	newHTTPResponse(w, http.StatusOK, successMessage)
}

// NotFoundResources not found api
func NotFoundResources(w http.ResponseWriter, r *http.Request) {
	newHTTPError(w, NewError(http.StatusNotFound, "Not found resources"))
}

func newHTTPError(w http.ResponseWriter, err error) {
	errMessage := err.(ErrorMessage)
	w.WriteHeader(errMessage.Code)
	w.Write(structToResponseBody(err))
	return
}

// creat http response
func newHTTPResponse(w http.ResponseWriter, code int, body interface{}) {
	w.WriteHeader(code)
	w.Write(structToResponseBody(body))
	return
}

// get user id from query parameter
func getUserID(r *http.Request) (string, error) {
	query := r.URL.Query()
	id := query.Get("id")

	if id != "" {
		if !validUUID.MatchString(id) {
			return id, NewError(http.StatusBadRequest, "invalid user id format")
		}
	}

	return id, nil
}

// struct to response body
func structToResponseBody(data interface{}) []byte {
	json, err := json.Marshal(&data)

	if err != nil {
		return []byte(err.Error())
	}

	return json
}

// request body to struct
func requestBodyToStruct(r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return NewError(http.StatusBadRequest, err.Error())
	}

	return nil
}

// check post and put request header
func checkRequestHeader(r *http.Request) error {
	if r.Method == http.MethodPost || r.Method == http.MethodPut {
		code := http.StatusBadRequest
		if r.Header.Get("Content-type") != "application/json" {
			return NewError(code, "Content-type is not application/json")
		}

		if r.ContentLength == 0 {
			return NewError(code, "Request body length is 0")
		}
	}

	return nil
}

package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNotUserFound(t *testing.T) {
	//initialization:

	//Execution
	user, err := GetUser(0)

	//Validation
	assert.Nil(t, user,  "we were not expecting a user with id = 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, "user 0 does not exists", err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)

	/*
	if user != nil {
		t.Error("we were not expecting a user with id = 0")
	}


	if err == nil {
		t.Error("we were not expecting an error when user id is 0")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Error("we were not expecting 404 when user is not found")
	}
	 */
}

func TestGetUserNotError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Ali", user.FirstName)
	assert.EqualValues(t, "Laode", user.LastName)
	assert.EqualValues(t, "odealidj.go@gmail.com", user.Email)
}


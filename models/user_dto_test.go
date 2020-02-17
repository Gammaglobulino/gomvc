package models

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)
	assert.Nil(t, user, "we were not expecting a user here")
	assert.NotNil(t, err, "we were expecting an error here")
	assert.EqualValues(t, err.StatusCode, http.StatusNotFound, "we were expecting 404")
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserUserFound(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, user, "we were expecting a user here")
	assert.Nil(t, err, "we were expecting an error here")
	assert.EqualValues(t, user.Id, 123, "we are expecting a user id 123")

}

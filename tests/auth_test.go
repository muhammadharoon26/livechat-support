package tests

import (
	"livechat-support/controllers"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	response := controllers.Login()
	assert.NotNil(t, response)
}

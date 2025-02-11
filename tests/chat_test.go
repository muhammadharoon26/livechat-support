package tests

import (
	"livechat-support/controllers"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveMessage(t *testing.T) {
	response := controllers.SaveMessage()
	assert.NotNil(t, response)
}

package server

import (
	"github.com/stretchr/testify/assert"
	"port/pkg/db"
	"testing"
)

func TestSetupCorrectHandler(t *testing.T) {
	assert := assert.New(t)
	handler, _ := db.NewDbHandler()
	serviceConfig, err := SetupService(handler)

	assert.Equal(nil, err, "With correct handler err should be nil")
	assert.NotEqual(nil, serviceConfig, "Service config should not be nil")
}

func TestSetupWrongHandler(t *testing.T) {
	assert := assert.New(t)
	_, err := SetupService(nil)
	assert.NotEqual(nil, err, "Raise error if handler is nil")
}

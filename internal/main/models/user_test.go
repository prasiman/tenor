package models

import (
	_ "main/internal/config/env"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test create user
func TestCreate(t *testing.T) {
	created, _ := UserCreate("pras_iman@hotmail.com", "pras_iman@hotmail.com")

	assert.Equal(t, true, created)
}

// Test authenticate user
func TestAuthenticate(t *testing.T) {
	exist, _, _ := UserAuthenticate("myusername", "mypassword")

	assert.Equal(t, false, exist)
}

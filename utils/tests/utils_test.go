package utils

import (
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sundowndev/go-covermyass/utils"
)

func TestGetUserHomeDir(t *testing.T) {
	userProfile := &user.User{HomeDir: "/test"}

	UserCurrent := func() (*user.User, error) {
		return userProfile, nil
	}

	assert.Equal(t, utils.GetUserHomeDir(UserCurrent), "/test", "they should be equal")
}

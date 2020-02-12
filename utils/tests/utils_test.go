package utils

import (
	"os/user"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sundowndev/go-covermyass/utils"
)

func TestGetUserHomeDir(t *testing.T) {
	Convey("Should assert user home directory", t, func() {
		userProfile := &user.User{HomeDir: "/test"}

		UserCurrent := func() (*user.User, error) {
			return userProfile, nil
		}

		So(utils.GetUserHomeDir(UserCurrent), ShouldEqual, "/test")
	})
}

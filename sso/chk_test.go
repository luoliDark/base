package sso

import (
	"fmt"
	"testing"

	"github.com/luoliDark/base/sysmodel"
)

func TestCheckFileOAuthToken(t *testing.T) {
	to := "50c7def5845743248ab97f491b104448:2025-05-07"
	err := ValidateFileOAuthToken(to, "")
	fmt.Println(err)
}

func TestGetFileOAuthToken(t *testing.T) {
	to := GetFileOAuthTokenByUser(&sysmodel.SSOUser{UserID: "admin7", EntID: "1"})
	fmt.Println(to)
}

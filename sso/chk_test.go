package sso

import (
	"fmt"
	"testing"
)

func TestGetUserByToken(t *testing.T) {
	bean, err := GetUserByToken("jsz:a748689ffe344bea969e1a3a4e8244cc:2022-07-17")
	fmt.Print(bean, err)
}

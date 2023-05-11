package redishelper

import (
	"fmt"
	"paas/base/confighelper"
	"testing"
)

func TestGetList(t *testing.T) {
	//fmt.Println(GetList("aa", 0, "jsz888"))

	s := GetString(confighelper.GetEnterpriseID(), confighelper.GetSessionDbIndex(), "292606d8044d4faf83cf0aa181d62a9c")
	fmt.Println(s)

}

func TestIsExists(t *testing.T) {
	IsExists("", 0, "aaa")
}

package redishelper

import "testing"

func TestSetList(t *testing.T) {

	arr := []string{"aaaa", "bbbb", "cccc", "ddd"}

	SetList("aa", 0, "jsz888", arr)
}

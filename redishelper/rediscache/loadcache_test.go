package rediscache

import (
	"fmt"
	"testing"
)

func TestLoadRedisCache(t *testing.T) {
	//err := LoadRedisMain("", "1", "50201", "wf")
	//fmt.Println(err)

	err := LoadRedisMain("", "0", "0", "other")
	fmt.Println(err)

}

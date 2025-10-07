package rediscache

import (
	"fmt"
	"testing"
)

func TestGetList(t *testing.T) {

}

func TestGetHashMap(t *testing.T) {

	//re := GetHashMap(1, 50201, "sys_fpage", "50201")
	//fmt.Println(re)
	re := GetHashMap(1, 50201, "sys_wfflow", "50201_1")
	fmt.Println(re)
}

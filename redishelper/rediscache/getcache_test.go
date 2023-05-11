package rediscache

import (
	"fmt"
	"testing"
)

func TestGetList(t *testing.T) {
	fmt.Println(GetListMap("sys_wfsteprelation_targetstepid_innerjoin_sys_wfstep_50201_start-f7f7cce2493a4be8b964a99235cbb474"))

}

func TestGetHashMap(t *testing.T) {
	fmt.Println(GetHashMap("sys_fpagefield_4996"))
}

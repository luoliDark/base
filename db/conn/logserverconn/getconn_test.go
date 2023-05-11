package logserverconn

import (
	"fmt"
	"testing"
)

func TestGetDB(t *testing.T) {
	eng, _ := GetDB()
	fmt.Println(eng.QueryString("select * from eb_user limit 1 "))
}

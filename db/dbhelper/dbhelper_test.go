package dbhelper

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	//fmt.Println(Query("", true, "select * from eb_user limit 2"))
	fmt.Println(Query("", true, "select * from eb_dept limit 2"))

}

package commutil

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

//

func TestHasSpecialCharByStr(t *testing.T) {
	//var a int64
	//a := 45019

	fmt.Println(excelDateToDate("45019.55826388888888889"))

}

func excelDateToDate(number string) time.Time {
	excelTime := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	var days, _ = strconv.Atoi(number)
	return excelTime.Add(time.Second * time.Duration(days*86400))
}

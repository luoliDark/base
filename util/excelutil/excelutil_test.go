package excelutil

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/luoliDark/base/util/commutil"
)

func TestCreateExcelPath(t *testing.T) {

	//t2:=excelDateToDate("45039")
	//t3:=t2.Add(time.Duration(-1419328704))
	//fmt.Println( t3)

	fmt.Println(convertToFormatTime("45039.1419328704"))
	//fmt.Println(time.ParseDuration("1419328704"))

	var in float64 = 45039.9999884259

	excelEpoch := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)

	tm := excelEpoch.Add(time.Duration(in * float64(24*time.Hour)))

	fmt.Println(tm.Format(commutil.Time_Fomat12)) // 2021-12-01 13:17:10.000003072 +0000 UTC

}

func excelDateToDate(excelDate string) time.Time {
	excelTime := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	var days, _ = strconv.Atoi(excelDate)
	return excelTime.Add(time.Second * time.Duration(days*86400))
}

func excelDateToDate2(excelDate string) time.Time {
	excelTime := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	var days, _ = strconv.Atoi(excelDate)
	return excelTime.Add(time.Second * time.Duration(days*86400))
}

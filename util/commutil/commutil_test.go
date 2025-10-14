package commutil

import (
	"fmt"
	"testing"

	"github.com/luoliDark/base/util/commutil"
	"github.com/luoliDark/base/util/jsonutil"
)

func TestToFloat32(t *testing.T) {

	startdate := TimParse("2024-12-24", commutil.Time_Fomat03)
	enddate := TimParse("2025-01-22", commutil.Time_Fomat03)

	diffMonth := GetSubMonth(startdate, enddate)
	fmt.Println(diffMonth)

}

func TestAppendStr(t *testing.T) {
	fmt.Println(AppendStr("122", "222", "55554fdd", ToString(232)))

}

func TestToBool(t *testing.T) {
	fmt.Println(ToBool("true"), ToBool(1), ToBool(nil))
}

func TestGetNowMonth(t *testing.T) {
	fmt.Println(GetNowMonth())
}

func TestGetResultBeanByJson(t *testing.T) {
	s := `付成都你六姐（青浦宝龙店）马祥3月工资李如秀（小时工2月薪资）工时96小时。
银行卡号：620522001030693446`
	a, e := jsonutil.StrToJson(s)
	fmt.Println(a, e)
	b := GetResultBeanByJson(s)
	fmt.Println(b)
	fmt.Println(a.String())
}

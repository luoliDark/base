package commutil

import (
	"fmt"
	"testing"
)

func TestToFloat32(t *testing.T) {
	var a float64
	a = 33.44

	var b float32
	b = 33.4

	var c string
	c = "22.33434343434232323323"

	b = ToFloat32(a)
	fmt.Println(b)
	b = ToFloat32(c)

	fmt.Println(b)
}

func TestAppendStr(t *testing.T) {
	fmt.Println(AppendStr("122", "222", "55554fdd", ToString(232)))

}

func TestToBool(t *testing.T) {
	fmt.Println(ToBool("true"), ToBool(1), ToBool(nil))
}

func TestGetNowTime(t *testing.T) {
	fmt.Println(GetNowTime())
}

func TestToString(t *testing.T) {
	fmt.Println(ToString("1231.12"))

	var a float64 = 1231.12
	fmt.Println(ToString(a))
}

func TestGetHttpCatchErrMsg(t *testing.T) {
	//GetHttpCatchErrMsg(fmt.Errorf("123"),"")
	defer CatchError()
	//panic(123)
	panic("123")
}

func TestToBool1(t *testing.T) {
	var a int64 = 1
	var b int = 1
	var c int16 = 1
	fmt.Println(ToBool(a))
	fmt.Println(ToBool(b))
	fmt.Println(ToBool(c))
	fmt.Println(ToBool("1"))
	fmt.Println(ToBool("2"))
	fmt.Println(ToBool("0"))
	fmt.Println(ToBool(false))
	fmt.Println(ToBool(true))
}

func TestGetNowTimeByFormatStr(t *testing.T) {
	fmt.Println(GetNowTimeByFormatStr("20060102_150405"))
}

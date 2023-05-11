package pinyinutil

import (
	"fmt"
	"testing"
)

func TestConvertAbbr(t *testing.T) {
	fmt.Println(PinYinConvertAbbr("我是你爸爸", "-"))
}

func TestPinYinCovertFull(t *testing.T) {
	fmt.Println(PinYinCovertFull("我是你妈妈，他是爸爸", "-"))
}

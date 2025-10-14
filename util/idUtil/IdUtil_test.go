package idUtil

import (
	"fmt"
	"testing"

	"github.com/luoliDark/base/util/commutil"
)

func Test_main(t *testing.T) {
	var aaa float64
	aaa = 0.00
	toString := commutil.ToString(aaa)
	n, _ := fmt.Sscanf(commutil.ToString(aaa), "%e", &aaa)
	fmt.Print(toString)
	fmt.Print(n)

}

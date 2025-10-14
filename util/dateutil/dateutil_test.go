package dateutil

import (
	"fmt"
	"strings"
	"testing"
)

func TestSearchInterval(t *testing.T) {
	var list = []string{"", "2022-01-02 00:00:00", "2022-01-02 12:22:33"}
	s, i := SearchInterval(list)
	fmt.Println(s)
	fmt.Println(i)
	join := strings.Join(list, ",")
	fmt.Println(join)
	same := JudgeDayIsSame(list[0], list[1])
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
	fmt.Println(same)
}

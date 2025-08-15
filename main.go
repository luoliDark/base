package main

import (
	"fmt"

	"github.com/luoliDark/base/db/conn"
)

func main() {
	en, _ := conn.GetDianShangDB()

	fmt.Println(en)
}

package main

import (
	"fmt"

	"github.com/luoliDark/base/db/conn"
)

func main() {
	en, _ := conn.GetBusFaDb()

	fmt.Println(en)
}

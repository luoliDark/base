package main

import (
	"fmt"

	"github.com/luoliDark/base/db/enum"

	"github.com/luoliDark/base/db/conn"
)

func main() {
	en, _ := conn.GetDBConnection("", true, enum.DianShang)

	fmt.Println(en)
}

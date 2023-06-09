package main

import (
	"fmt"

	"github.com/luoliDark/base/db/conn"
)

func main() {
	db, _ := conn.GetBusFaDbOriginal()
	sesion := db.NewSession()
	res, _ := sesion.QueryString("select code from c_store limit  1")
	sesion.Close()
	fmt.Println(res)
}

package main

import (
	"github.com/luoliDark/base/db/dbhelper"
)

func main() {
	dbhelper.Query("", true, "select * from eb_user limit 1")
}

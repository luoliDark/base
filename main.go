package main

import (
	"fmt"

	"github.com/luoliDark/base/redishelper/rediscache"
	"github.com/luoliDark/base/util/commutil"
)

func main() {
	sqltable := rediscache.GetHashMap("sys_fpage_" + commutil.ToString("30202"))

	fmt.Println(sqltable)
}

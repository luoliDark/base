package menuloading

import (
	"fmt"
	"testing"
)

func TestQueryMenuMain(t *testing.T) {
	// 刷新redis 缓存
	//reloadredis.LoadRedisCache("zhongxinjian")

	//menu := QueryMenu("3", false)
	//fmt.Println(menu)
	menu := QueryMenuMainBySql("3", 0)
	fmt.Println(menu)
}

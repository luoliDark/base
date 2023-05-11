package excelutil

import (
	"fmt"
	"testing"

	"github.com/luoliDark/runingproject/services/public"
)

func TestBatchInsertListDataByExcelize2(t *testing.T) {
	fmt.Println(public.GetStringByBus("BigDataExport"))
}

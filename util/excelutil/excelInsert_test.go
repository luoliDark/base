package excelutil

import (
	"fmt"
	"paas/runingproject/services/public"
	"testing"
)

func TestBatchInsertListDataByExcelize2(t *testing.T) {
	fmt.Println(public.GetStringByBus("BigDataExport"))
}

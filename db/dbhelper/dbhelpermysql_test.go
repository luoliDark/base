package dbhelper

import (
	"base/sysmodel"
	"fmt"
	"strings"
	"testing"
)

func TestExecProc(t *testing.T) {
	//execProc_OutParamMysql("11",true,"UP_SYS_GetBillId",[]string{"billid","result"},"EB_Role",1)
}

func TestQueryPaging(t *testing.T) {

	result, err := QueryByJson("11", true, "select * from exp_normain where billid=?", "{primarykey}")

	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	fmt.Println(result)
	//queryPaging_mysql("1", true, 0, 10, "select * from sys_fpage where pid > ? order by pid desc ", 200000)
}

func TestConcatSqlserverExeProc(t *testing.T) {
	parList := make([]sysmodel.ProcPar, 0)

	par := new(sysmodel.ProcPar)
	par.DataType = "varchar(200)"
	par.ProcName = "billid"
	par.IsOutPut = true

	parList = append(parList, sysmodel.ProcPar{IsOutPut: false, ProcName: "TableName", DataType: "varchar(200)", ProcValue: "EB_Role"})
	parList = append(parList, sysmodel.ProcPar{IsOutPut: false, ProcName: "pri_count", DataType: "varchar(800)", ProcValue: "1"})
	parList = append(parList, sysmodel.ProcPar{IsOutPut: true, ProcName: "billid", DataType: "varchar(800)"})
	parList = append(parList, sysmodel.ProcPar{IsOutPut: true, ProcName: "result", DataType: "varchar(800)"})
	res, _ := concatSqlserverExeProc("UP_SYS_GetBillId", parList)
	fmt.Println(res)
	res, _ = concatMysqlCallProc("UP_SYS_GetBillId", parList)
	fmt.Println(res)
	result, outValues, err := ExecProc_ResultSetValue("11", true, "UP_SYS_GetBillId", parList)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(outValues)
}

func TestConcanSqlserverPaging(t *testing.T) {

	var s = "21232'ssds'ds"
	fmt.Println(strings.ReplaceAll(s, "'", "''"))

	//sql := concatSqlserverPageSql("pid",0,20,"select * from sys_fpage ")
	//fmt.Println(sql)
}

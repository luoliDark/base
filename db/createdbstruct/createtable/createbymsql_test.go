package createtable

import (
	"fmt"
	"testing"

	"github.com/luoliDark/base/sysmodel"
)

func TestCreateTable(t *testing.T) {
	//fields := make([]sysmodel.SqlField,0)
	//createTable_mysql("admin","test_insert01","TID", true,"",fields,false)
	//values := getProcParList_mysql("admin","UP_SYS_GetBillId")
	//fmt.Println(values)
	//res := getSqlTableStruct_mysql("admin","sys_fpage")
	//fmt.Println(res)

	fields := make([]sysmodel.SqlField, 0)

	fields = append(fields, sysmodel.SqlField{
		ColName:            "id",
		DataType:           sysmodel.SqlDataType{DataType: "int", Length: 0},
		DataLength:         0,
		ColMemo:            "主键自增",
		DefaultValue:       "",
		IsPrimaryKey:       true,
		IsUnique:           false,
		IsNotNull:          true,
		IsFK:               false,
		FK_TableName:       "",
		FK_TablePrimaryKey: "",
		CheckStr:           "",
	})
	fields = append(fields, sysmodel.SqlField{
		ColName:            "name",
		DataType:           sysmodel.SqlDataType{DataType: "varchar", Length: 50},
		DataLength:         0,
		ColMemo:            "名称",
		DefaultValue:       "",
		IsPrimaryKey:       true,
		IsUnique:           false,
		IsNotNull:          false,
		IsFK:               false,
		FK_TableName:       "",
		FK_TablePrimaryKey: "",
		CheckStr:           "",
	})
	fields = append(fields, sysmodel.SqlField{
		ColName:            "age",
		DataType:           sysmodel.SqlDataType{DataType: "int", Length: 0},
		DataLength:         0,
		ColMemo:            "年龄",
		DefaultValue:       "",
		IsPrimaryKey:       true,
		IsUnique:           false,
		IsNotNull:          false,
		IsFK:               false,
		FK_TableName:       "",
		FK_TablePrimaryKey: "",
		CheckStr:           "",
	})
	//res := createTable_sqlserver("admin","test_insert01","id",true, "", fields, false)

	//res := createTable_mysql("admin","test_insert01","id",true, "", fields, false)
	res := getSqlTableStruct_sqlserver("admin", "sys_fpage")
	fmt.Println(res)
}

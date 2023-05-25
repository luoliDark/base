package createmodel

import (
	"testing"
)

//调用方法生成结构
func TestCreateStructBySql(t *testing.T) {
	CreateStructBySql("glo_setprocess")
	//CreateStructBySql("Con_RegDetail_Renewal")
}

//增加包名，指定文件位置生成
/**
SELECT * FROM `fa_glmain`
SELECT * FROM `fa_glassdetail`
SELECT * FROM `fa_gldetail`
*/
func TestCreateStructBySql2(t *testing.T) {
	//tname := "ex_sendlog"    //表名
	//createmodel := "voucher" //包名
	////path := `E:\GoProject\src\easyfa\easyfa\model\voucher`   //文件生成存放的目录
	//path := `F:\GoFile\easyfa\easyfa\model\voucher`          //文件生成存放的目录
	//fliepath := path + "\\" + strings.ToLower(tname) + ".go" //生成文件名
	//
	////CreateStructBySql2(tname, createmodel,
	////	fliepath)
}

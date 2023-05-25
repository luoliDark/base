package createmodel

import (
	"fmt"
	"strings"

	"github.com/luoliDark/base/db/conn"

	"github.com/gohouse/converter"
)

// CreateStructBySql 根据SQL生成实体类
func CreateStructBySql(tame string) {

	tame = strings.ToLower(tame)

	// 初始化
	t2t := converter.NewTable2Struct()
	// 个性化配置
	t2t.Config(&converter.T2tConfig{
		// 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
		RmTagIfUcFirsted: false,
		// tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
		TagToLower: true,
		// 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
		UcFirstOnly: false,
		//// 每个struct放入单独的文件,默认false,放入同一个文件(暂未提供)
		//SeperatFile: false,
	})
	// 开始迁移转换
	err := t2t.
		// 指定某个表,如果不指定,则默认全部表都迁移
		Table(tame).
		// 表前缀
		//Prefix("sys_fgridfield").
		// 是否添加json tag
		EnableJsonTag(true).
		// 生成struct的包名(默认为空的话, 则取名为: package model)
		PackageName("createmodel").
		// tag字段的key值,默认是orm
		TagKey("xorm").
		// 是否添加结构体方法获取表名
		RealNameMethod("TableName").
		// 生成的结构体保存路径
		//SavePath("/paas/base/db/createdbstruct/dbtomodel/modelp.go").
		// 数据库dsn,这里可以使用 t2t.DB() 代替,参数为 *db.DB 对象
		Dsn(conn.GetConnStr("")).
		// 执行
		Run()
	if err != nil {
		fmt.Println(err)
	}
}

func CreateStructBySql2(tame, createmodel, path string) {

	tame = strings.ToLower(tame)
	if createmodel == "" {
		createmodel = "createmodel"
	}

	// 初始化
	t2t := converter.NewTable2Struct()
	// 个性化配置
	t2t.Config(&converter.T2tConfig{
		// 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
		RmTagIfUcFirsted: false,
		// tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
		TagToLower: true,
		// 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
		UcFirstOnly: false,
		//// 每个struct放入单独的文件,默认false,放入同一个文件(暂未提供)
		//SeperatFile: false,
	})
	// 开始迁移转换
	err := t2t.
		// 指定某个表,如果不指定,则默认全部表都迁移
		Table(tame).
		// 表前缀
		//Prefix("sys_fgridfield").
		// 是否添加json tag
		EnableJsonTag(true).
		// 生成struct的包名(默认为空的话, 则取名为: package model)
		PackageName(createmodel).
		// tag字段的key值,默认是orm
		TagKey("xorm").
		// 是否添加结构体方法获取表名
		RealNameMethod("TableName").
		// 生成的结构体保存路径
		SavePath(path).
		// 数据库dsn,这里可以使用 t2t.DB() 代替,参数为 *db.DB 对象
		Dsn(conn.GetConnStr("")).
		// 执行
		Run()
	if err != nil {
		fmt.Println(err)
	}
}

package sysmodel

import (
	"github.com/luoliDark/base/confighelper"
	"github.com/luoliDark/base/db/conn"
	"github.com/luoliDark/base/loghelper"
	"github.com/luoliDark/base/sysmodel/logtype"
	"github.com/luoliDark/base/util/commutil"
)

type SqlDataType struct {
	//数据类型
	DataType string
	//长度
	Length int
	//长度
	DecimalLength int
}

func (this SqlDataType) GetDatatype() string {
	currDB := confighelper.GetCurrdb()
	switch currDB {
	case conn.DBTYPE_MYSQL:
		return this.getMysqlDatatype()
	case conn.DBTYPE_SQLSERVER:
		return this.getSqlServerDatatype()
	default:
		loghelper.ByHighError(logtype.NoDBErr, "请检查config.ini", "")
		panic("请检查config.ini")
	}

}

//获取mysql中数据类型
func (this *SqlDataType) getMysqlDatatype() string {

	//todo 需完善
	switch this.DataType {
	case "int":
		if this.Length == 0 {
			return "int"
		} else {
			return commutil.AppendStr(this.DataType, "(", commutil.ToString(this.Length), ")")
		}
	case "bit":
		return "int(1)"
	case "decimal": // 小数
		return commutil.AppendStr(this.DataType, "(18,", commutil.ToString(this.DecimalLength), ")")
	default:
		if this.Length > 0 {
			return commutil.AppendStr(this.DataType, "(", commutil.ToString(this.Length), ")")
		} else {
			return this.DataType
		}
	}
}

//获取sqlServer中数据类型
func (this *SqlDataType) getSqlServerDatatype() string {
	//todo 需完善

	switch this.DataType {
	case "decimal":
		return commutil.AppendStr(this.DataType, "(18,", commutil.ToString(this.Length), ")")
	default:
		if this.Length > 0 {
			return commutil.AppendStr(this.DataType, "(", commutil.ToString(this.Length), ")")
		} else {
			return this.DataType
		}
	}

}

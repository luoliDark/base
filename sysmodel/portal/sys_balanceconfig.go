package portal

/**
 * @Author: weiyg
 * @Date: 2020/4/24 10:24
 * @describe:
 */

type Sys_balanceconfig struct {
	ID            int    `xorm:"id" json:"id"`
	ReadPid       string `xorm:"readpid" json:"readpid"`
	Readbilltype  string `xorm:"readbilltype" json:"readbilltype"`
	SqlTable      string `xorm:"sqltable" json:"sqltable"`
	BalanceSql    string `xorm:"balancesql" json:"balancesql"`
	TotalMoneySql string `xorm:"totalmoneysql" json:"totalmoneysql"`
	ChkCol        string `xorm:"chkcol" json:"chkcol"`
	Title         string `xorm:"title" json:"title"`
	IsByLogin     int    `xorm:"isbylogin" json:"isbylogin"`
	IsByForm      int    `xorm:"isbyform" json:"isbyform"`
	FormValueCol  string `xorm:"formvaluecol" json:"formvaluecol"`
	IsOpen        int    `xorm:"isopen" json:"isopen"`
	WhereSql      string `xorm:"wheresql" json:"wheresql"`
}

func (*Sys_balanceconfig) TableName() string {
	return "Sys_balanceconfig"
}

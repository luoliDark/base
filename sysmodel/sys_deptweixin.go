package sysmodel

type Sys_DeptWeiXin struct {
	Id          int    `xorm:"id" json:"id"`
	Dept_id     string `xorm:"dept_id" json:"dept_id"`
	Deptcode    string `xorm:"deptcode" json:"deptcode"`
	Name        string `xorm:"name" json:"name"`
	Name_en     string `xorm:"name_en" json:"name_en"`
	Parentid    string `xorm:"parentid" json:"parentid"`
	Orders      string `xorm:"orders" json:"orders"`
	Entid       int    `xorm:"entid" json:"entid"`
	Manageruid  string `xorm:"manageruid" json:"manageruid"`
	Create_time string `xorm:"create_time" json:"create_time"`
}

func (this *Sys_DeptWeiXin) TableName() string {
	return "sys_deptweixin"
}

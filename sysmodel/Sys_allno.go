package sysmodel

type SysAllno struct {
	BillId     string `xorm:"billid" json:"billid"`
	Pid        int    `xorm:"pid" json:"pid"`
	BillNo     string `xorm:"billno" json:"billno"`
	RealNumber string `xorm:"realnumber" json:"realnumber"`
}

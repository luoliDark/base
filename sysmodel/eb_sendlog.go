package sysmodel

type Eb_sendlog struct {
	Logid       int    `xorm:"logid" json:"logid"`
	Tbname      string `xorm:"tbname" json:"tbname"`
	Primarykey  string `xorm:"primarykey" json:"primarykey"`
	Create_UID  string `xorm:"create_uid" json:"create_uid"`
	Create_Date string `xorm:"create_Date" json:"create_Date"`
	Update_UID  string `xorm:"update_UID" json:"update_UID"`
	Update_Date string `xorm:"update_Date" json:"update_Date"`
	IsDiscard   string `xorm:"isDiscard" json:"isDiscard"`
	Errmsg      string `xorm:"errmsg" json:"errmsg"`
	Sendtext    string `xorm:"sendtext" json:"sendtext"`
}

func (*Eb_sendlog) TableName() string {
	return "eb_sendlog"
}

package model

type Redis_RefreshConfig struct {
	Id            string `xorm:"id" json:"id"`
	Ctype         string `xorm:"ctype" json:"ctype"`
	XmlName       string `xorm:"xmlName" json:"xmlName"`
	ParentXmlName string `xorm:"parentXmlName" json:"parentXmlName"`
	FilterStr     string `xorm:"FilterStr" json:"FilterStr"`
	TableBM       string `xorm:"TableBM" json:"TableBM"`
	SortId        string `xorm:"sortId" json:"sortId"`
	Memo          string `xorm:"Memo" json:"Memo"`
	IsHasChild    int    `xorm:"IsHasChild" json:"IsHasChild"`
}

func (*Redis_RefreshConfig) TableName() string {
	return "Redis_RefreshConfig"
}

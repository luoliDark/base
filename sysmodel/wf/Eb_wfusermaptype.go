package wf

type Eb_wfusermaptype struct {
	MapID       string             `xorm:"mapid" json:"mapid"`
	MapCode     string             `xorm:"mapcode" json:"mapcode"` // 编码
	MapName     string             `xorm:"mapname" json:"mapname"` // 名称
	Memo        string             `xorm:"memo" json:"memo"`       // 备注
	NewGuid     string             `xorm:"newguid" json:"newguid"`
	Entid       int                `xorm:"entid" json:"entid"`
	IsMain      int                `xorm:"ismain" json:"ismain"`
	MapSqlTable string             `xorm:"mapsqltable" json:"mapsqltable"`
	SqlColList  []*Eb_wfusermapcol `xorm:"sqlcolList" json:"sqlcolList"`
}

func (*Eb_wfusermaptype) TableName() string {
	return "eb_wfusermaptype"
}

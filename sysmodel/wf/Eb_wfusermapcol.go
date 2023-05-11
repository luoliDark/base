package wf

type Eb_wfusermapcol struct {
	DetailID   int    `xorm:"detailid" json:"detailid"`
	MapID      string `xorm:"mapid" json:"mapid"`     // 外键
	MapCode    string `xorm:"mapcode" json:"mapcode"` // 外键
	MapSqlCol  string `xorm:"mapsqlcol" json:"mapsqlcol"`
	MapColName string `xorm:"mapcolname" json:"mapcolname"`
	Entid      int    `xorm:"entid" json:"entid"`
	Memo       string `xorm:"memo" json:"memo"`
}

func (*Eb_wfusermapcol) TableName() string {
	return "eb_wfusermapcol"
}

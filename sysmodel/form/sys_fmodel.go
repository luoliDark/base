package form

type Sys_fmodel struct {
	ModelID      int     `xorm:"modelid" json:"modelid"`
	ModelName    string  `xorm:"modelname" json:"modelname"`
	ParentID     int     `xorm:"parentid" json:"parentid"`
	IsOpen       int     `xorm:"isopen" json:"isopen"`
	TablePre     string  `xorm:"tablepre" json:"tablepre"`
	SubSysID     int     `xorm:"subsysid" json:"subsysid"`
	SortID       float64 `xorm:"sortid" json:"sortid"`
	IsBackConfig int     `xorm:"isbackconfig" json:"isbackconfig"`
	LenMemo      string  `xorm:"lenmemo" json:"lenmemo"`
	Memo         string  `xorm:"memo" json:"memo"`
	MenuImg      string  `xorm:"menuimg" json:"menuimg"`
}

func (*Sys_fmodel) TableName() string {
	return "sys_fmodel"
}

package eb

type Eb_area struct {
	AreaID      string `xorm:"areaid" json:"areaid"`
	AreaCode    string `xorm:"areacode" json:"areacode"`
	AreaName    string `xorm:"areaname" json:"areaname"`
	ParentID    string `xorm:"parentid" json:"parentid"`
	FullName    string `xorm:"fullname" json:"fullname"`
	Memo        string `xorm:"memo" json:"memo"`
	CreateUID   string `xorm:"create_uid" json:"create_uid"`
	CreateDate  string `xorm:"create_date" json:"create_date"`
	UpdateUID   string `xorm:"update_uid" json:"update_uid"`
	UpdateDate  string `xorm:"update_date" json:"update_date"`
	IsDiscard   int    `xorm:"isdiscard" json:"isdiscard"`
	DisCardUID  string `xorm:"discard_uid" json:"discard_uid"`
	DisCardDate string `xorm:"discard_date" json:"discard_date"`
	Ver         string `xorm:"ver" json:"ver"`
	GLCode      string `xorm:"gl_code" json:"gl_code"`
	NewGuID     string `xorm:"newguid" json:"newguid"`
	IsModify    int    `xorm:"ismodify" json:"ismodify"`
	CurrPID     int    `xorm:"currpid" json:"currpid"`
	SaveSource  string `xorm:"savesource" json:"savesource"`
	ParentCode  string `xorm:"parentcode" json:"parentcode"`
	Entid       int    `xorm:"entid" json:"entid"`
	Isimport    int    `xorm:"isimport" json:"isimport"`
	Ispub       int    `xorm:"ispub" json:"ispub"`
	Ishaschild  int    `xorm:"ishaschild" json:"ishaschild"`
}

func (*Eb_area) TableName() string {
	return "eb_area"
}

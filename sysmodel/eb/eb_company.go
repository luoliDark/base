package eb

type Eb_company struct {
	CompID         string `xorm:"compid" json:"compid"`
	CompCode       string `xorm:"compcode" json:"compcode"`
	CompName       string `xorm:"compname" json:"compname"`
	ParentID       string `xorm:"parentid" json:"parentid"`
	FullName       string `xorm:"fullname" json:"fullname"`
	CeoManagerUID  string `xorm:"ceomanageruid" json:"ceomanageruid"`
	CWUID          string `xorm:"cwuid" json:"cwuid"`
	CWManagerUID   string `xorm:"cwmanageruid" json:"cwmanageruid"`
	AdminUID       string `xorm:"adminuid" json:"adminuid"`
	DefBankAccount string `xorm:"defbankaccount" json:"defbankaccount"`
	RateCode       string `xorm:"ratecode" json:"ratecode"`
	Memo           string `xorm:"memo" json:"memo"`
	CreateUID      string `xorm:"create_uid" json:"create_uid"`
	CreateDate     string `xorm:"create_date" json:"create_date"`
	UpdateUID      string `xorm:"update_uid" json:"update_uid"`
	UpdateDate     string `xorm:"update_date" json:"update_date"`
	IsDiscard      int    `xorm:"isdiscard" json:"isdiscard"`
	DisCardUID     string `xorm:"discard_uid" json:"discard_uid"`
	DisCardDate    string `xorm:"discard_date" json:"discard_date"`
	Ver            string `xorm:"ver" json:"ver"`
	GLCode         string `xorm:"gl_code" json:"gl_code"`
	NewGuID        string `xorm:"newguid" json:"newguid"`
	IsModify       int    `xorm:"ismodify" json:"ismodify"`
	CurrPID        int    `xorm:"currpid" json:"currpid"`
	SaveSource     string `xorm:"savesource" json:"savesource"`
	ParentCode     string `xorm:"parentcode" json:"parentcode"`
	CFO            string `xorm:"cfo" json:"cfo"`
	CEO            string `xorm:"ceo" json:"ceo"`
	BookingTicket  string `xorm:"bookingticket" json:"bookingticket"`
	Entid          int    `xorm:"entid" json:"entid"`
	Isimport       int    `xorm:"isimport" json:"isimport"`
	CompType       string `xorm:"comptype" json:"comptype"`
}

func (*Eb_company) TableName() string {
	return "eb_company"
}

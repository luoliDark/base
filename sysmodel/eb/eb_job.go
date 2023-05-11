package eb

type Eb_job struct {
	JobID       string `xorm:"jobid" json:"jobid"`
	JobCode     string `xorm:"jobcode" json:"jobcode"`
	JobName     string `xorm:"jobname" json:"jobname"`
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
	Entid       int    `xorm:"entid" json:"entid"`
}

func (*Eb_job) TableName() string {
	return "eb_job"
}

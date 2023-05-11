package portal

type Eb_banklhh struct {
	BankAdd string `xorm:"bankadd" json:"bankadd"`
	LHH     string `xorm:"lhh" json:"lhh"`
}

func (*Eb_banklhh) TableName() string {
	return "eb_banklhh"
}

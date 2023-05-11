package eb

type Eb_userimage struct {
	Userid string `xorm:"userid" json:"userid"`
	Image  string `xorm:"image" json:"image"`
}

func (*Eb_userimage) TableName() string {
	return "eb_userimage"
}

package sysmodel

type Eb_uservsjob struct {
	UserVsJobID int    `xorm:"UserVsJobID" json:"UserVsJobID"`
	UserID      string `xorm:"UserID" json:"UserID"`
	JobID       string `xorm:"JobID" json:"JobID"`
	EntId       int    `xorm:"entid" json:"entid"`
}

func (*Eb_uservsjob) TableName() string {
	return "eb_uservsjob"
}

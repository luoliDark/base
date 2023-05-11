package eb

type Eb_UserVsJob struct {
	UserVsJobID string `xorm:"uservsjobid" json:"uservsjobid"`
	UserID      string `xorm:"userid" json:"userid"`
	JobID       string `xorm:"jobid" json:"jobid"`
	EntId       string `xorm:"entid" json:"entid"`
	InsertDate  string `xorm:"insertdate" json:"insertdate"`
}

func (*Eb_UserVsJob) TableName() string {
	return "Eb_UserVsJob"
}

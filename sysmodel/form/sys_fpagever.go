package form

type Sys_fpagever struct {
	ID          string `xorm:"id" json:"id"`
	Pid         int    `xorm:"pid" json:"pid"`
	VerId       int    `xorm:"verid" json:"verid"`
	ReleaseUid  string `xorm:"releaseuid" json:"releaseuid"`
	ReleaseDate string `xorm:"releasedate" json:"releasedate"`
	EntId       string `xorm:"entid" json:"entid"`
	MenuUrl     string `xorm:"menuurl" json:"menuurl"`
	EditUrl     string `xorm:"editurl" json:"editurl"`
}

func (*Sys_fpagever) TableName() string {
	return "sys_fpagever"
}

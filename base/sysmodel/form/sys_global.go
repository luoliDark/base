package form

type Sys_global struct {
	Gid        int    `xorm:"gid" json:"gid"`
	SoftVer    string `xorm:"softver" json:"softver"`
	SoftUpDate string `xorm:"softupdate" json:"softupdate"`
	MenuVer    int    `xorm:"menuver" json:"menuver"`
}

func (*Sys_global) TableName() string {
	return "sys_global"
}

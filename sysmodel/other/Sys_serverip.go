package other

type Sys_serverip struct {
	Gid        string `xorm:"gid" json:"gid"`
	IP         string `xorm:"ip" json:"ip"`
	ServerName string `xorm:"servername" json:"servername"`
}

func (*Sys_serverip) TableName() string {
	return "sys_serverip"
}

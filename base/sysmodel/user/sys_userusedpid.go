package user

/**
 * @Author: yix
 * @Date: 2020/5/2 21:19
 * @describe:
 */
type Sys_userusedpid struct {
	UseId   int    `xorm:"useid" json:"useid"`
	UserId  string `xorm:"userid" json:"userid"`
	Pid     int    `xorm:"pid" json:"pid"`
	UseTime string `xorm:"usetime" json:"usetime"`
	EntId   int    `xorm:"entid" json:"entid"`
}

func (*Sys_userusedpid) TableName() string {
	return "sys_userusedpid"
}

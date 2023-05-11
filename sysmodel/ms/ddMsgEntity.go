package ms

//用于接收钉钉的消息实体类
type DDMsgEntity struct {
	Pid            int               //单据类型
	PrimaryKey     string            //主健
	UserIdList     []string          `xorm:"useridlist" json:"useridlist"`
	To_all_user    bool              `xorm:"to_all_user" json:"to_all_user"`
	Templateid     int               `xorm:"templateid" json:"templateid"`
	Templateparams map[string]string `xorm:"templateparams" json:"templateparams"`
	Entid          int               `xorm:"entid" json:"entid"`
}

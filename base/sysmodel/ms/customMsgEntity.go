package ms

//用于接收自定义的消息实体类
type CustomMsgEntity struct {
	Pid        int      `xorm:"pid" json:"pid"`               //单据类型
	PrimaryKey string   `xorm:"primarykey" json:"primarykey"` //主健
	Title      string   `xorm:"title" json:"title"`
	Body       string   `xorm:"body" json:"body"`
	Bodys      []string `xorm:"bodys" json:"bodys"`
	LinkUrl    string   `xorm:"linkUrl" json:"linkUrl"`
	ToUsers    []string `xorm:"toUsers" json:"toUsers"`
	EntId      int      `xorm:"entid" json:"entid"`
}

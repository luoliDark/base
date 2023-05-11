package bbs

/**
 * @Author: weiyg
 * @Date: 2020/4/29 8:50
 * @describe:
 */
type Bbsbypidreciver struct {
	ID          int    `xorm:"id" json:"id"`
	MsgId       string `xorm:"msgid" json:"msgid"`
	MsgType     string `xorm:"msgtype" json:"msgtype"`
	ReciverUid  string `xorm:"reciver_uid" json:"reciver_uid"`
	IsViewed    int    `xorm:"isviewed" json:"isviewed"`
	ReciverTime string `xorm:"recivertime" json:"recivertime"`
	ViewTime    string `xorm:"viewtime" json:"viewtime"`
	EntId       int    `xorm:"entid" json:"entid"`
}

func (*Bbsbypidreciver) TableName() string {
	return "bbsbypidreciver"
}

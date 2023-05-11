package wf

/**
 * @describe: 历史抄送用户
 *
 * @Author: YiXin
 * @Date: 2021/3/1
 */

type Sys_wfcchistoryuser struct {
	Userid  string `xorm:"userid" json:"userid"`   // 用户id
	Ccusers string `xorm:"ccusers" json:"ccusers"` // 最近一次抄送用户id
	Entid   int    `xorm:"entid" json:"entid"`     // 企业id
}

func (*Sys_wfcchistoryuser) TableName() string {
	return "sys_wfcchistoryuser"
}

package form

/**
 * @Author: yix
 * @Date: 2020/5/1 20:32
 * @describe: 系统字典项分类主表	映射类
 */
type Sys_infomain struct {
	GroupID   int    `xorm:"groupid" json:"groupid"`
	GroupCode string `xorm:"groupcode" json:"groupcode"`
	GroupName string `xorm:"groupname" json:"groupname"`
	Memo      string `xorm:"memo" json:"memo"`
}

func (*Sys_infomain) TableName() string {
	return "sys_infomain"
}

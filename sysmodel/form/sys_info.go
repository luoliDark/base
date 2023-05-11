package form

/**
 * @Author: yix
 * @Date: 2020/5/1 20:32
 * @describe: 系统字典项明细	映射类
 */
type Sys_info struct {
	InfoID     int     `xorm:"infoid" json:"infoid"`
	InfoKey    string  `xorm:"infokey" json:"infokey"`
	InfoValue  string  `xorm:"infovalue" json:"infovalue"`
	GroupID    int     `xorm:"groupid" json:"groupid"`
	Memo       string  `xorm:"memo" json:"memo"`
	SortID     float64 `xorm:"sortid" json:"sortid"`
	CreateUid  string  `xorm:"create_uid" json:"create_uid"`
	CreateDate string  `xorm:"create_date" json:"create_date"`
	UpdateUid  string  `xorm:"update_uid" json:"update_uid"`
	UpdateDate string  `xorm:"update_date" json:"update_date"`
}

func (*Sys_info) TableName() string {
	return "sys_info"
}

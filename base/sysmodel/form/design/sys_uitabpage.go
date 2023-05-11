package design

type Sys_uitabpage struct {
	TabID      string  `xorm:"tabid" json:"tabid"`
	TabName    string  `xorm:"tabname" json:"tabname"`
	SortID     float64 `xorm:"sortid" json:"sortid"`
	TabGroupID string  `xorm:"tabgroupid" json:"tabgroupid"`
	TabPic     string  `xorm:"tabpic" json:"tabpic"`
	Pid        int     `xorm:"pid" json:"pid"`
	EntId      int     `xorm:"entid" json:"entid"`
}

func (*Sys_uitabpage) TableName() string {
	return "sys_uitabpage"
}

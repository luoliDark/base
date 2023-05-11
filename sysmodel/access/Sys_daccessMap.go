package access

type Sys_daccessmap struct {
	Id             int    `xorm:"id" json:"id"`
	Pid            int    `xorm:"pid" json:"pid"`
	Userid         string `xorm:"userid" json:"userid"`
	Primarykey     string `xorm:"primarykey" json:"primarykey"`             // 单据数据ID
	DataSyncSource string `xorm:"data_sync_source" json:"data_sync_source"` // 数据来源，默认system_add为系统内添加，和外部钉钉同步的数据做区分
	Entid          int    `xorm:"entid" json:"entid"`
}

func (*Sys_daccessmap) TableName() string {
	return "sys_daccessmap"
}

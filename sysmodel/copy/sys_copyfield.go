package copy

type Sys_copyfield struct {
	DetailID        string  `xorm:"detailid" json:"detailid"`
	CPFromID        string  `xorm:"cpfromid" json:"cpfromid"`
	SourceCol       string  `xorm:"sourcecol" json:"sourcecol"`
	SourceReadSql   string  `xorm:"sourcereadsql" json:"sourcereadsql"`
	SourceShow      string  `xorm:"source_show" json:"source_show"`
	SourceName      string  `xorm:"source_name" json:"source_name"`
	TargetCol       string  `xorm:"targetcol" json:"targetcol"`
	IsSourceMain    int     `xorm:"issourcemain" json:"issourcemain"`
	IsGroupCtr      int     `xorm:"isgroupctr" json:"isgroupctr"`
	ISort           float32 `xorm:"isort" json:"isort"`
	IsHide          int     `xorm:"ishide" json:"ishide"`
	IsCopyTo        int     `xorm:"iscopyto" json:"iscopyto"`
	IsGroupField    int     `xorm:"isgroupfield" json:"isgroupfield"`
	IsCheckGroupCtr int     `xorm:"ischeckgroupctr" json:"ischeckgroupctr"`
	ColFunc         string  `xorm:"colfunc" json:"colfunc"`
	IsSum           int     `xorm:"issum" json:"issum"`
	IsQuery         int     `xorm:"isquery" json:"isquery"`
	ControlType     string  `xorm:"controltype" json:"controltype"`
	DataSource      string  `xorm:"datasource" json:"datasource"`
	DataSourceWhere string  `xorm:"datasourcewhere" json:"datasourcewhere"`
	WriteBackCol    string  `xorm:"writebackcol" json:"writebackcol"`
	OverCtrWhere    string  `xorm:"overctrwhere" json:"overctrwhere"`
	DsGroupId       int     `xorm:"dsgroupid" json:"dsgroupid"`
	IsMobile        int     `xorm:"ismobile" json:"ismobile"`
	IsWriteBack     int     `xorm:"iswriteback" json:"iswriteback"`
	IsCtrBySubmit   int     `xorm:"isctrbysubmit" json:"isctrbysubmit"`
	SqlDataType     string  `xorm:"sqldatatype" json:"sqldatatype"`
	WriteBalanceCol string  `xorm:"writebalancecol" json:"writebalancecol"` //反写字段余额取值
}

func (*Sys_copyfield) TableName() string {
	return "sys_copyfield"
}

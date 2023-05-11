package design

type Sys_uiformcol struct {
	ID                 int     `xorm:"id" json:"id"`
	Pid                int     `xorm:"pid" json:"pid"`
	Name               string  `xorm:"name" json:"name"`
	SqlCol             string  `xorm:"sqlcol" json:"sqlcol"`
	TabID              string  `xorm:"tabid" json:"tabid"`
	ObjPkid            string  `xorm:"obj_pkid" json:"obj_pkid"`
	ObjType            string  `xorm:"objtype" json:"objtype"`
	Sortid             float64 `xorm:"sortid" json:"sortid"`
	Ishide             int     `xorm:"ishide" json:"ishide"`
	ControlType        string  `xorm:"controltype" json:"controltype"`
	Width              string  `xorm:"width" json:"width"`
	Height             string  `xorm:"height" json:"height"`
	IsRequired         int     `xorm:"isrequired" json:"isrequired"`
	IsReadOnly         int     `xorm:"isreadonly" json:"isreadonly"`
	HideMath           string  `xorm:"hidemath" json:"hidemath"`
	RequireMath        string  `xorm:"requiremath" json:"requiremath"`
	ReadonlyMath       string  `xorm:"readonlymath" json:"readonlymath"`
	IsQueryShow        int     `xorm:"isqueryshow" json:"isqueryshow"`
	IsOnlyListPageShow int     `xorm:"isonlylistpageshow" json:"isonlylistpageshow"`
	SqlDataType        string  `xorm:"sqldatatype" json:"sqldatatype"`
	DataSource         string  `xorm:"datasource" json:"datasource"`
	DataSourceWhere    string  `xorm:"datasourcewhere" json:"datasourcewhere"`
	DsGroupId          string  `xorm:"dsgroupid" json:"dsgroupid"`
	IsSum              int     `xorm:"issum" json:"issum"`
	IsSingle           int     `xorm:"issingle" json:"issingle"`
	IsMobile           int     `xorm:"ismobile" json:"ismobile"`
	ColDirection       string  `xorm:"coldirection" json:"coldirection"`
	MathStr            string  `xorm:"mathstr" json:"mathstr"`
	DefaultValue       string  `xorm:"defaultvalue " json:"defaultvalue"`
	EventByAuto        string  `xorm:"eventbyauto " json:"eventbyauto"`
	TargetRefCol       string  `xorm:"targetrefcol" json:"targetrefcol"` //关联的目标字段
	RefCol             string  `xorm:"refcol" json:"refcol"`             //关联的来源字段
	IsOnlyLast         int     `xorm:"isonlylast" json:"isonlylast"`     //只能选择末级，仅树形档案
	OpenType           string  `xorm:"opentype" json:"opentype"`         // 弹窗显示方式   针对有数据源的字段 下拉（tree,tree2,treevslist,list)
	DataControlMode    string  `xorm:"datacontrolmode" json:"datacontrolmode"`
	SumbyFilterMath    string  `xorm:"sumbyfiltermath" json:"sumbyfiltermath"`
	DsWindowShowCol    string  `xorm:"dswindowshowcol" json:"dswindowshowcol"`
	ChkByMaxLen        int     `xorm:"chkbymaxlen" json:"chkbymaxlen"`
	ChkByRegular       string  `xorm:"chkbyregular" json:"chkbyregular"`
	Placeholder        string  `xorm:"placeholder" json:"placeholder"`
	IsZQ               int     `xorm:"iszq" json:"iszq"`         // 是否链接跳转
	ZQParams           string  `xorm:"zqparams" json:"zqparams"` // 链接跳转Url格式
	EntId              int     `xorm:"entid" json:"entid"`
}

func (*Sys_uiformcol) TableName() string {
	return "sys_uiformcol"
}

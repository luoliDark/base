package uitemplete

type Sys_uiformgrid struct {
	ID                 int    `xorm:"id" json:"id"`
	Pid                int    `xorm:"pid" json:"pid"`
	GridID             int    `xorm:"gridid" json:"gridid"`
	Name               string `xorm:"name" json:"name"`
	SqlCol             string `xorm:"sqlcol" json:"sqlcol"`
	DetailID           string `xorm:"detailid" json:"detailid"`
	SortID             int    `xorm:"sortid" json:"sortid"`
	Ishide             int    `xorm:"ishide" json:"ishide"`
	ControlType        string `xorm:"controltype" json:"controltype"`
	Isreadonly         int    `xorm:"isreadonly" json:"isreadonly"`
	Isrequired         int    `xorm:"isrequired" json:"isrequired"`
	RequireMath        string `xorm:"requiremath" json:"requiremath"`
	ReadonlyMath       string `xorm:"readonlymath" json:"readonlymath"`
	HideMath           string `xorm:"hidemath" json:"hidemath"`
	Width              string `xorm:"width" json:"width"`
	SqlDataType        string `xorm:"sqldatatype" json:"sqldatatype"`
	DataSource         string `xorm:"datasource" json:"datasource"`
	DataSourceWhere    string `xorm:"datasourcewhere" json:"datasourcewhere"`
	DsGroupId          string `xorm:"dsgroupid" json:"dsgroupid"`
	IsSum              int    `xorm:"issum" json:"issum"`
	IsSingle           int    `xorm:"issingle" json:"issingle"`
	IsMobile           int    `xorm:"ismobile" json:"ismobile"`
	MathStr            string `xorm:"mathstr" json:"mathstr"`
	DefaultValue       string `xorm:"defaultvalue " json:"defaultvalue"`
	EventByAuto        string `xorm:"eventbyauto " json:"eventbyauto"`
	TargetRefCol       string `xorm:"targetrefcol" json:"targetrefcol"`             //关联的目标字段
	RefCol             string `xorm:"refcol" json:"refcol"`                         //关联的来源字段
	IsOnlyLast         int    `xorm:"isonlylast" json:"isonlylast"`                 //只能选择末级，仅树形档案
	OpenType           string `xorm:"opentype" json:"opentype"`                     // 弹窗显示方式   针对有数据源的字段 下拉（tree,tree2,treevslist,list)
	IsOnlyEditPageShow int    `xorm:"isonlyeditpageshow" json:"isonlyeditpageshow"` // 字段仅编辑列表页显示
	EntId              int    `xorm:"entid" json:"entid"`
	ColDirection       string `xorm:"coldirection" json:"coldirection"`
	Ismobilelistshow   int    `xorm:"ismobile_listshow" json:"ismobilelistshow"`
	Ismobilemerge      int    `xorm:"ismobile_merge" json:"ismobilemerge"`
	Ismobileshowlab    int    `xorm:"ismobile_showlab" json:"ismobileshowlab"`
	Ismobilelisttop    int    `xorm:"ismobile_listtop" json:"ismobilelisttop"`
	DataControlMode    string `xorm:"datacontrolmode" json:"datacontrolmode"`
	SumbyFilterMath    string `xorm:"sumbyfiltermath" json:"sumbyfiltermath"`
	DsWindowShowCol    string `xorm:"dswindowshowcol" json:"dswindowshowcol"`
	ChkByMaxLen        int    `xorm:"chkbymaxlen" json:"chkbymaxlen"`
	ChkByRegular       string `xorm:"chkbyregular" json:"chkbyregular"`
	Placeholder        string `xorm:"placeholder" json:"placeholder"`
}

func (*Sys_uiformgrid) TableName() string {
	return "sys_uiformgrid"
}

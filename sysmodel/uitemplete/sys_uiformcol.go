package uitemplete

type Sys_uiformcol struct {
	ID                 int     `xorm:"id" json:"id"`
	Pid                int     `xorm:"pid" json:"pid"`
	UITempID           string  `xorm:"uitempid" json:"uitempid"`
	Name               string  `xorm:"name" json:"name"`
	SqlCol             string  `xorm:"sqlcol" json:"sqlcol"`
	TabID              string  `xorm:"tabid" json:"tabid"`
	ObjPkid            string  `xorm:"obj_pkid" json:"obj_pkid"`
	ObjType            string  `xorm:"objtype" json:"objtype"`
	DataSource         string  `xorm:"datasource" json:"datasource"`
	DataSourceWhere    string  `xorm:"datasourcewhere" json:"datasourcewhere"`
	SqlDataType        string  `xorm:"sqldatatype" json:"sqldatatype"`
	EventbyAuto        string  `xorm:"eventbyauto" json:"eventbyauto"`
	TargetRefCol       string  `xorm:"targetrefcol" json:"targetrefcol"` //关联的目标字段
	RefCol             string  `xorm:"refcol" json:"refcol"`             //关联的来源字段
	IsOnlyLast         int     `xorm:"isonlylast" json:"isonlylast"`     //只能选择末级，仅树形档案
	OpenType           string  `xorm:"opentype" json:"opentype"`         // 弹窗显示方式   针对有数据源的字段 下拉（tree,tree2,treevslist,list)
	IsZQ               int     `xorm:"iszq" json:"iszq"`                 // 是否链接跳转
	ZQParams           string  `xorm:"zqparams" json:"zqparams"`         // 链接跳转Url格式
	DefaultValue       string  `xorm:"defaultvalue" json:"defaultvalue"`
	MathStr            string  `xorm:"mathstr" json:"mathstr"`
	SortID             float64 `xorm:"sortid" json:"sortid"`
	Ishide             int     `xorm:"ishide" json:"ishide"`
	IsReadOnly         int     `xorm:"isreadonly" json:"isreadonly"`
	IsAllowUserInput   int     `xorm:"isallowuserinput" json:"isallowuserinput"` //下拉框时是否允许用户直接输入 例：一次性供应商要输入户名
	IsRequired         int     `xorm:"isrequired" json:"isrequired"`
	IsMobile           int     `xorm:"ismobile" json:"ismobile"`
	IsOnlyListPageShow int     `xorm:"isonlylistpageshow" json:"isonlylistpageshow"`
	IsOnlyEditPageShow int     `xorm:"isonlyeditpageshow" json:"isonlyeditpageshow"`
	Width              string  `xorm:"width" json:"width"`
	Height             string  `xorm:"height" json:"height"`
	ControlType        string  `xorm:"controltype" json:"controltype"`
	IsCustom           string  `xorm:"iscustom" json:"iscustom"`
	//子表字段专用
	GridID          int    `xorm:"gridid" json:"gridid"`
	DetailID        string `xorm:"detailid" json:"detailid"`
	ColDirection    string `xorm:"coldirection" json:"coldirection"`
	DataControlMode string `xorm:"datacontrolmode" json:"datacontrolmode"`
	SumbyFilterMath string `xorm:"sumbyfiltermath" json:"sumbyfiltermath"`
	DsWindowShowCol string `xorm:"dswindowshowcol" json:"dswindowshowcol"`
	ChkByMaxLen     int    `xorm:"chkbymaxlen" json:"chkbymaxlen"`
	ChkByRegular    string `xorm:"chkbyregular" json:"chkbyregular"`
	Placeholder     string `xorm:"placeholder" json:"placeholder"`
}

func (*Sys_uiformcol) TableName() string {
	return "sys_uiformcol"
}

package design

import "reflect"

type FormDesignResult struct {
	Pid    string           `json:"pid"`
	List   []FormDesignItem `json:"list"`
	Config struct {
		LabelWidth    int    `json:"labelWidth"`
		LabelPosition string `json:"labelPosition"`
		Size          string `json:"size"`
		CustomClass   string `json:"customClass"`
		UI            string `json:"ui"`
		Layout        string `json:"layout"`
		LabelCol      int    `json:"labelCol"`
		Width         string `json:"width"`
		FormName      string `json:"formName"`
	} `json:"config"`
}
type FormDesignItem struct {
	Type           string            `json:"type"`
	Icon           string            `json:"icon"`
	Options        FormDesignOptions `json:"options"`
	ParrentOptions FormDesignOptions `json:"parrentOptions"`
	Name           string            `json:"name"`
	Key            string            `json:"key"`
	Model          string            `json:"model"`
	Rules          []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"rules"`
	Columns []struct {
		Span int           `json:"span"`
		Xs   int           `json:"xs"`
		Sm   int           `json:"sm"`
		Md   int           `json:"md"`
		Lg   int           `json:"lg"`
		Xl   int           `json:"xl"`
		List []interface{} `json:"list"`
	} `json:"columns,omitempty"`
	Tabs         []FormDesignTab  `json:"tabs,omitempty"`
	TableColumns []FormDesignItem `json:"tableColumns"`
}
type FormDesignOptions struct {
	DataSourceWhere    string          `json:"datasourcewhere"`
	Grid               string          `json:"grid"`
	SelectorStyle      string          `json:"selectorStyle"`
	MathFormula        string          `json:"mathFormula"`
	EventByAuto        string          `json:"eventbyauto"`  //事件执行代码
	TargetRefCol       string          `json:"targetrefcol"` //数据关联字段
	RefCol             string          `json:"refcol"`       //关联的来源字段
	IsOnlyLast         bool            `json:"isonlylast"`   //只能选择末级，仅树形档案
	IsBatchEdit        bool            `json:"isbatchedit"`  //只能选择末级，仅树形档案
	OpenType           string          `json:"opentype"`     // 弹窗显示方式   针对有数据源的字段 下拉（tree,tree2,treevslist,list)
	DsWindowShowCol    string          `json:"dswindowshowcol"`
	ChkByMaxLen        int             `json:"chkbymaxlen"`
	ChkByRegular       string          `json:"chkbyregular"`
	DataControlMode    string          `json:"datacontrolmode"`
	SumbyFilterMath    string          `json:"sumbyfiltermath"`
	IsZQ               bool            `json:"iszq"`     // 是否链接跳转
	ZQParams           string          `json:"zqparams"` // 链接跳转Url格式
	OnChangedformula   string          `json:"onChangedformula"`
	RemoteRule         string          `json:"remoteRule"`
	DataSource         string          `json:"datasource"`
	DsGroupId          string          `json:"dsgroupid"`
	HideFormula        string          `json:"hideFormula"`
	OnlyReadFormula    string          `json:"onlyReadFormula"`
	RequiredFormula    string          `json:"requiredFormula"`
	SqlCode            string          `json:"sqlCode"`
	DefaultValue       interface{}     `json:"defaultValue"`
	Width              interface{}     `json:"width"`
	CustomClass        string          `json:"customClass"`
	LabelWidth         int             `json:"labelWidth"`
	IsLabelWidth       bool            `json:"isLabelWidth"`
	Hidden             bool            `json:"hidden"`
	IsQueryShow        bool            `json:"isQueryShow"`
	IsOnlyListPageShow bool            `json:"isOnlyListPageShow"` //仅主表列表页显示
	IsOnlyEditPageShow bool            `json:"isOnlyEditPageShow"` //仅子表编辑页显示
	DataBind           bool            `json:"dataBind"`
	CustomToolbar      [][]interface{} `json:"customToolbar"`
	Disabled           bool            `json:"disabled"`
	IsMobile           bool            `json:"ismobile"`
	ColDirection       string          `json:"coldirection"`
	Ismobilelistshow   bool            `json:"ismobilelistshow"`
	Ismobilemerge      bool            `json:"ismobilemerge"`
	Ismobileshowlab    bool            `json:"ismobileshowlab"`
	Ismobilelisttop    bool            `json:"ismobilelisttop"`
	Formula            string          `json:"formula"`
	RemoteFunc         string          `json:"remoteFunc"`
	RemoteOption       string          `json:"remoteOption"`
	Placeholder        string          `json:"placeholder"`
	Clearable          bool            `json:"clearable"`
	Remote             bool            `json:"remote"`
	RemoteType         string          `json:"remoteType"`
	RemoteOptions      []interface{}   `json:"remoteOptions"`
	Props              struct {
		Value    string `json:"value"`
		Label    string `json:"label"`
		Children string `json:"children"`
	} `json:"props"`
	Size struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"size"`
	TokenFunc       string        `json:"tokenFunc"`
	Token           string        `json:"token"`
	Domain          string        `json:"domain"`
	Readonly        bool          `json:"readonly"`
	Limit           int           `json:"limit"`
	Multiple        bool          `json:"multiple"`
	IsQiniu         bool          `json:"isQiniu"`
	IsDelete        bool          `json:"isDelete"`
	Min             int           `json:"min"`
	IsEdit          bool          `json:"isEdit"`
	Action          string        `json:"action"`
	Headers         []interface{} `json:"headers"`
	ContentPosition string        `json:"contentPosition"`
	Gutter          int           `json:"gutter"`
	Justify         string        `json:"justify"`
	Align           string        `json:"align"`
	Flex            bool          `json:"flex"`
	Responsive      bool          `json:"responsive"`
	Tip             string        `json:"tip"`
	DataType        string        `json:"dataType"`
	Pattern         string        `json:"pattern"`
	ShowPassword    bool          `json:"showPassword"`
	CountAll        bool          `json:"countAll"`
	Inline          bool          `json:"inline"`
	ShowLabel       bool          `json:"showLabel"`
	Options         []struct {
		Value string `json:"value"`
	} `json:"options"`
	Required bool `json:"required"`
}

type FormDesignTab struct {
	Label string           `json:"label"`
	Name  string           `json:"name"`
	List  []FormDesignItem `json:"list"`
}

func (a FormDesignResult) IsEmpty() bool {

	return reflect.DeepEqual(a, FormDesignResult{})
}

package form

type Sys_fpagefield struct {
	DetailID        string  `xorm:"detailid" json:"detailid"`
	Pid             int     `xorm:"pid" json:"pid"`
	SqlCol          string  `xorm:"sqlcol" json:"sqlcol"`
	SelectStr       string  `xorm:"selectstr" json:"selectstr"`
	SelectShow      string  `xorm:"select_show" json:"select_show"`
	Name            string  `xorm:"name" json:"name"`
	SortID          float64 `xorm:"sortid" json:"sortid"`
	IsEnable        int     `xorm:"isenable" json:"isenable"`
	IsSave          int     `xorm:"issave" json:"issave"`
	ControlType     string  `xorm:"controltype" json:"controltype"`
	DataSource      string  `xorm:"datasource" json:"datasource"`
	DataSourceWhere string  `xorm:"datasourcewhere" json:"datasourcewhere"`
	IsCodeName      int     `xorm:"iscodename" json:"iscodename"`
	SqlDataType     string  `xorm:"sqldatatype" json:"sqldatatype"`
	EventByAuto     string  `xorm:"eventbyauto" json:"eventbyauto"`
	MathStr         string  `xorm:"mathstr" json:"mathstr"`
	DefaultValue    string  `xorm:"defaultvalue" json:"defaultvalue"`
	TargetRefCol    string  `xorm:"targetrefcol" json:"targetrefcol"`
	RefCol          string  `xorm:"refcol" json:"refcol"`
	IsSum           int     `xorm:"issum" json:"issum"`
	Fnumber         int     `xorm:"fnumber" json:"fnumber"`
	IsOnlyLast      int     `xorm:"isonlylast" json:"isonlylast"`
	IsSingle        int     `xorm:"issingle" json:"issingle"`
	IsAutoEx        int     `xorm:"isautoex" json:"isautoex"`
	IsExcelImport   int     `xorm:"isexcelimport" json:"isexcelimport"`
	IsNeedConvert   int     `xorm:"isneedconvert" json:"isneedconvert"`
	IsConvertByCode int     `xorm:"isconvertbycode" json:"isconvertbycode"`
	IsConvertByName int     `xorm:"isconvertbyname" json:"isconvertbyname"`
	DataFormat      string  `xorm:"dataformat" json:"dataformat"`
	SpecialSet      string  `xorm:"specialset" json:"specialset"`
	IsFinance       int     `xorm:"isfinance" json:"isfinance"`
	ShowDiscard     int     `xorm:"showdiscard" json:"showdiscard"`
	IsStanDs        int     `xorm:"isstands" json:"isstands"`
	DSGroupID       int     `xorm:"dsgroupid" json:"dsgroupid"`
	IsMobile        int     `xorm:"ismobile" json:"ismobile"`
	OpenType        string  `xorm:"opentype" json:"opentype"`
	IsPrint         int     `xorm:"isprint" json:"isprint"`
	IsMobileSSP     int     `xorm:"ismobilessp" json:"ismobilessp"`
	MobileRangeID   string  `xorm:"mobilerangeid" json:"mobilerangeid"`
	IsDisCard       int     `xorm:"isdiscard" json:"isdiscard"`
	DisCard_Date    string  `xorm:"discard_date" json:"discard_date"`
}

type Sys_fpagefieldShow struct {
	DetailID        string  `xorm:"detailid" json:"detailid"`
	Pid             int     `xorm:"pid" json:"pid"`
	SqlCol          string  `xorm:"sqlcol" json:"sqlcol"`
	SelectStr       string  `xorm:"selectstr" json:"selectstr"`
	SelectShow      string  `xorm:"select_show" json:"select_show"`
	Name            string  `xorm:"name" json:"name"`
	SortID          float64 `xorm:"sortid" json:"sortid"`
	IsEnable        int     `xorm:"isenable" json:"isenable"`
	IsSave          int     `xorm:"issave" json:"issave"`
	ControlType     string  `xorm:"controltype" json:"controltype"`
	DataSource      string  `xorm:"datasource" json:"datasource"`
	DataSourceShow  string  `xorm:"datasourceshow" json:"datasourceshow"`
	DataSourceWhere string  `xorm:"datasourcewhere" json:"datasourcewhere"`
	IsCodeName      int     `xorm:"iscodename" json:"iscodename"`
	SqlDataType     string  `xorm:"sqldatatype" json:"sqldatatype"`
	EventByAuto     string  `xorm:"eventbyauto" json:"eventbyauto"`
	MathStr         string  `xorm:"mathstr" json:"mathstr"`
	DefaultValue    string  `xorm:"defaultvalue" json:"defaultvalue"`
	TargetRefCol    string  `xorm:"targetrefcol" json:"targetrefcol"`
	RefCol          string  `xorm:"refcol" json:"refcol"`
	IsSum           int     `xorm:"issum" json:"issum"`
	Fnumber         int     `xorm:"fnumber" json:"fnumber"`
	IsOnlyLast      int     `xorm:"isonlylast" json:"isonlylast"`
	IsSingle        int     `xorm:"issingle" json:"issingle"`
	IsAutoEx        int     `xorm:"isautoex" json:"isautoex"`
	IsExcelImport   int     `xorm:"isexcelimport" json:"isexcelimport"`
	IsNeedConvert   int     `xorm:"isneedconvert" json:"isneedconvert"`
	IsConvertByCode int     `xorm:"isconvertbycode" json:"isconvertbycode"`
	IsConvertByName int     `xorm:"isconvertbyname" json:"isconvertbyname"`
	DataFormat      string  `xorm:"dataformat" json:"dataformat"`
	SpecialSet      string  `xorm:"specialset" json:"specialset"`
	IsFinance       int     `xorm:"isfinance" json:"isfinance"`
	ShowDiscard     int     `xorm:"showdiscard" json:"showdiscard"`
	IsStanDs        int     `xorm:"isstands" json:"isstands"`
	DSGroupID       int     `xorm:"dsgroupid" json:"dsgroupid"`
	IsMobile        int     `xorm:"ismobile" json:"ismobile"`
	OpenType        string  `xorm:"opentype" json:"opentype"`
	IsPrint         int     `xorm:"isprint" json:"isprint"`
	IsMobileSSP     int     `xorm:"ismobilessp" json:"ismobilessp"`
	MobileRangeID   string  `xorm:"mobilerangeid" json:"mobilerangeid"`
	IsDisCard       int     `xorm:"isdiscard" json:"isdiscard"`
	DisCard_Date    string  `xorm:"discard_date" json:"discard_date"`
}

func (*Sys_fpagefield) TableName() string {
	return "sys_fpagefield"
}

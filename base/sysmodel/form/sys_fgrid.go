package form

type Sys_fgrid struct {
	GridID              int     `xorm:"gridid" json:"gridid"`
	Pid                 int     `xorm:"pid" json:"pid"`
	GridName            string  `xorm:"gridname" json:"gridname"`
	PrimaryKey          string  `xorm:"primarykey" json:"primarykey"`
	RelationKey         string  `xorm:"relationkey" json:"relationkey"`
	SqlTableName        string  `xorm:"sqltablename" json:"sqltablename"`
	TableBM             string  `xorm:"tablebm" json:"tablebm"`
	SelectFromSql       string  `xorm:"selectfromsql" json:"selectfromsql"`
	WhereSql            string  `xorm:"wheresql" json:"wheresql"`
	OrderSql            string  `xorm:"ordersql" json:"ordersql"`
	GridType            string  `xorm:"gridtype" json:"gridtype"`
	IsAllowSave         int     `xorm:"isallowsave" json:"isallowsave"`
	IsEnable            int     `xorm:"isenable" json:"isenable"`
	IsExcelImport       int     `xorm:"isexcelimport" json:"isexcelimport"`
	SortID              float64 `xorm:"sortid" json:"sortid"`
	IsMustInput         int     `xorm:"ismustinput" json:"ismustinput"`
	Toolbars            string  `xorm:"toolbars" json:"toolbars"`
	IsHide              int     `xorm:"ishide" json:"ishide"`
	BusModel            string  `xorm:"busmodel" json:"busmodel"`
	MobilePhotoPath     string  `xorm:"mobilephotopath" json:"mobilephotopath"`
	Width               string  `xorm:"width" json:"width"`
	Height              string  `xorm:"height" json:"height"`
	Memo                string  `xorm:"memo" json:"memo"`
	ConfigWhereFormat   string  `xorm:"configwhereformat" json:"configwhereformat"`
	ConfigWhereFromSql  string  `xorm:"configwherefromsql" json:"configwherefromsql"`
	IsMobile            int     `xorm:"ismobile" json:"ismobile"`
	SumCol              string  `xorm:"sumcol" json:"sumcol"`
	ExcelReadStartIndex int     `xorm:"excelreadstartindex" json:"excelreadstartindex"`
}

func (*Sys_fgrid) TableName() string {
	return "sys_fgrid"
}

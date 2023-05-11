package form

type Sys_fpageid struct {
	Pid                 int     `xorm:"pid" json:"pid"`
	Pname               string  `xorm:"pname" json:"pname"`
	ModelID             int     `xorm:"modelid" json:"modelid"`
	FormType            string  `xorm:"formtype" json:"formtype"`
	IsExcelImport       int     `xorm:"isexcelimport" json:"isexcelimport"`
	EditType            string  `xorm:"edittype" json:"edittype"`
	MenuUrl             string  `xorm:"menuurl" json:"menuurl"`
	EditUrl             string  `xorm:"editurl" json:"editurl"`
	SortID              float64 `xorm:"sortid" json:"sortid"`
	IsPhysicsDelete     int     `xorm:"isphysicsdelete" json:"isphysicsdelete"`
	IsBillHelper        int     `xorm:"isbillhelper" json:"isbillhelper"`
	LeftPid             int     `xorm:"leftpid" json:"leftpid"`
	LeftJoinCol         string  `xorm:"leftjoincol" json:"leftjoincol"`
	MenuImg             string  `xorm:"menuimg" json:"menuimg"`
	IsYearBack          int     `xorm:"isyearback" json:"isyearback"`
	DataAccessLevel     int     `xorm:"dataaccesslevel" json:"dataaccesslevel"`
	NoFirstChar         string  `xorm:"nofirstchar" json:"nofirstchar"`
	IsSysFrame          int     `xorm:"issysframe" json:"issysframe"`
	IsOpen              int     `xorm:"isopen" json:"isopen"`
	FormAdminUid        string  `xorm:"formadminuid" json:"formadminuid"`
	IsMustUpLoad        int     `xorm:"ismustupload" json:"ismustupload"`
	IsMobile            int     `xorm:"ismobile" json:"ismobile"`
	ExcelReadStartIndex int     `xorm:"excelreadstartindex" json:"excelreadstartindex"`
	SubSysId            int     `xorm:"subsysid" json:"subsysid"`
	VerId               int     `xorm:"verid" json:"verid"`
	ReleaseDate         string  `xorm:"releasedate" json:"releasedate"`
	IsDiscard           int     `xorm:"isdiscard" json:"isdiscard"`
	Hqty                int     `xorm:"hqty" json:"hqty"`
	UseNoIntoPK         int     `xorm:"usenointopk" json:"usenointopk"`
	WindType            string  `xorm:"windtype" json:"windtype"`

	IsShowPayState     int `xorm:"isshowpaystate" json:"isshowpaystate"`
	IsDataCenter       int `xorm:"isdatacenter" json:"isdatacenter"`
	IsSaveDocSystem    int `xorm:"issavedocsystem" json:"issavedocsystem"`
	IsOpenRealCtr      int `xorm:"isopenrealctr" json:"isopenrealctr"`
	IsOpenMatrixAccess int `xorm:"isopenmatrixaccess" json:"isopenmatrixaccess"`
}

func (*Sys_fpageid) TableName() string {
	return "sys_fpageid"
}

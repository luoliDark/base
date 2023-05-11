package sysmodel

type EB_userwaitetranaccess struct {
	WaiteID      string `xorm:"WaiteID" json:"waiteid"`
	AccName      string `xorm:"AccName" json:"accname"`
	SqlTableName string `xorm:"SqlTableName" json:"sqltablename"`
	UserCol      string `xorm:"UserCol" json:"usercol"`
	AccCol       string `xorm:"AccCol" json:"acccol"`
	UserID       string `xorm:"UserID" json:"userid"`
	AccValue     string `xorm:"AccValue" json:"accvalue"`
	NewUserID    string `xorm:"-" json:"newuserid"`
	EntID        int    `xorm:"EntID" json:"-"`
}

func (*EB_userwaitetranaccess) TableName() string {
	return "eb_userwaitetranaccess"
}

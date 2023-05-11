package rule

type Sys_ruleformgroupdetail struct {
	RuleDetailID        string `xorm:"ruledetailid pk" json:"ruledetailid"`
	RuleGroupID         string `xorm:"rulegroupid" json:"rulegroupid"`
	ObjectID            string `xorm:"objectid" json:"objectid"`
	CType               string `xorm:"ctype" json:"ctype"`
	CheckType           string `xorm:"checktype" json:"checktype"`
	FormSqlTB           string `xorm:"formsqltb" json:"formsqltb"`
	FormSqlCol          string `xorm:"formsqlcol" json:"formsqlcol"`
	ByLoginType         string `xorm:"bylogintype" json:"bylogintype"`
	ByLoginDataSource   string `xorm:"bylogindatasource" json:"bylogindatasource"`
	LeftDataSource      int    `xorm:"leftdatasource" json:"leftdatasource"`
	LeftDataSourceCol   string `xorm:"leftdatasourcecol" json:"leftdatasourcecol"`
	LeftPid             int    `xorm:"leftpid" json:"leftpid"`
	LeftGridid          int    `xorm:"leftgridid" json:"leftgridid"`
	OP                  string `xorm:"op" json:"op"`
	JSOP                string `xorm:"jsop" json:"jsop"`
	WhereValueType      string `xorm:"wherevaluetype" json:"wherevaluetype"`
	WhereByStaticValue  string `xorm:"wherebystaticvalue" json:"wherebystaticvalue"`
	WhereByStaticText   string `xorm:"wherebystatictext" json:"wherebystatictext"`
	WhereByFormSqlTB    string `xorm:"wherebyformsqltb" json:"wherebyformsqltb"`
	WhereByFormSqlCol   string `xorm:"wherebyformsqlcol" json:"wherebyformsqlcol"`
	WhereByFormKeyWorkd string `xorm:"wherebyformkeyworkd" json:"wherebyformkeyworkd"`
	RightPid            int    `xorm:"rightpid" json:"rightpid"`
	RightGridid         int    `xorm:"rightgridid" json:"rightgridid"`
	IsCustom            int    `xorm:"iscustom" json:"iscustom"`
	CustomSqlScript     string `xorm:"customsqlscript" json:"customsqlscript"`
	WhereFullText       string `xorm:"wherefulltext" json:"wherefulltext"`
	EntId               int    `xorm:"entid" json:"entid"`
}

func (*Sys_ruleformgroupdetail) TableName() string {
	return "sys_ruleformgroupdetail"
}

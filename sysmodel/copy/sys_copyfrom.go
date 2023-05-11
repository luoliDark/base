package copy

type Sys_copyfrom struct {
	CPFromID            string `xorm:"cpfromid" json:"cpfromid"`
	CPID                string `xorm:"cpid" json:"cpid"`
	SourcePid           int    `xorm:"sourcepid" json:"sourcepid"`
	TargetPid           int    `xorm:"targetpid" json:"targetpid"`
	SourceGridID        int    `xorm:"sourcegridid" json:"sourcegridid"`
	TargetGridID        int    `xorm:"targetgridid" json:"targetgridid"`
	CopyType            string `xorm:"copytype" json:"copytype"`
	PrimaryKey          string `xorm:"primarykey" json:"primarykey"`
	TableBM             string `xorm:"tablebm" json:"tablebm"`
	FromSql             string `xorm:"fromsql" json:"fromsql"`
	WhereSql            string `xorm:"wheresql" json:"wheresql"`
	CopyWhere           string `xorm:"copywhere" json:"copywhere"`
	AutoCopyWhere       string `xorm:"autocopywhere" json:"autocopywhere"`
	SourceWriteCheckSql string `xorm:"sourcewritechecksql" json:"sourcewritechecksql"`
	AutoCreateWhereSql  string `xorm:"autocreatewheresql" json:"autocreatewheresql"` //提交或终审创建下游where条件
	Memo                string `xorm:"memo" json:"memo"`
	CopyFromShort       int    `xorm:"copyfromshort" json:"copyfromshort"`
	Orderby             string `xorm:"orderby" json:"orderby"`
	IsFirst             int    `xorm:"isfirst" json:"isfirst"`
	IsGroup             int    `xorm:"isgroup" json:"isgroup"`
	//反写模式：默认 ，0 金额反写 1 下游变更上游(新增,更新) 2 下游变更上游(新增,更新，删除)
	Backwritemode int              `xorm:"backwritemode" json:"backwritemode"`
	IsOpen        int              `xorm:"isopen" json:"isopen"`
	Fields        *[]Sys_copyfield `xorm:"-"`
}

type Sys_copyfromShow struct {
	CPFromID            string `xorm:"cpfromid" json:"cpfromid"`
	CPID                string `xorm:"cpid" json:"cpid"`
	SourcePid           int    `xorm:"sourcepid" json:"sourcepid"`
	TargetPid           int    `xorm:"targetpid" json:"targetpid"`
	SourceGridID        int    `xorm:"sourcegridid" json:"sourcegridid"`
	TargetGridID        int    `xorm:"targetgridid" json:"targetgridid"`
	SourceName          string `xorm:"sourcename" json:"sourcename"`
	TargetName          string `xorm:"targetname" json:"targetname"`
	SourceGridName      string `xorm:"sourcegridname" json:"sourcegridname"`
	TargetGridName      string `xorm:"targetgridname" json:"targetgridname"`
	CopyType            string `xorm:"copytype" json:"copytype"`
	PrimaryKey          string `xorm:"primarykey" json:"primarykey"`
	TableBM             string `xorm:"tablebm" json:"tablebm"`
	FromSql             string `xorm:"fromsql" json:"fromsql"`
	WhereSql            string `xorm:"wheresql" json:"wheresql"`
	CopyWhere           string `xorm:"copywhere" json:"copywhere"`
	AutoCopyWhere       string `xorm:"autocopywhere" json:"autocopywhere"`
	SourceWriteCheckSql string `xorm:"sourcewritechecksql" json:"sourcewritechecksql"`
	AutoCreateWhereSql  string `xorm:"autocreatewheresql" json:"autocreatewheresql"` //自动创建where条件
	Memo                string `xorm:"memo" json:"memo"`
	CopyFromShort       int    `xorm:"copyfromshort" json:"copyfromshort"`
	Orderby             string `xorm:"orderby" json:"orderby"`
	IsFirst             int    `xorm:"isfirst" json:"isfirst"`
	IsGroup             int    `xorm:"isgroup" json:"isgroup"`
	Backwritemode       int    `xorm:"backwritemode" json:"backwritemode"`
	IsOpen              int    `xorm:"isopen" json:"isopen"`
	Fields              *[]Sys_copyfield
}

func (*Sys_copyfrom) TableName() string {
	return "sys_copyfrom"
}

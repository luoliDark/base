package sysmodel

// SQL字段对象 (注创建表
type SqlField struct {
	ColName            string      // 字段名称
	DataType           SqlDataType // 数据类型
	DataLength         int         // 字段长度
	ColMemo            string      // 描述信息
	DefaultValue       string      //默认值
	IsPrimaryKey       bool
	IsUnique           bool   // 是否唯一
	IsNotNull          bool   // 是否必填
	IsFK               bool   // 是否有外键
	FK_TableName       string // 外键关联表名
	FK_TablePrimaryKey string // 外键关联表主键
	CheckStr           string // 字段检查约束 money > 100
	IsRename           bool   // 当前字段是否是重命名
	RenameName         string // 重命名名称
}

// SQL表结构对象
type SqlTableStruct struct {
	TableName            string //表名
	PrimaryKey           string //主健列名
	ForeignKey           string //外健列名
	IsIdentity           bool   //主健是否为自增列
	IsStringByPrimaryKey bool   //主健是否为字符串
	Pid                  int
	GridID               int
	ColList              []SqlField //字段清单
}

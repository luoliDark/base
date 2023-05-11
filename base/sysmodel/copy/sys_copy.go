package copy

type Sys_copy struct {
	CPID       string `xorm:"cpid" json:"cpid"`
	SourcePid  int    `xorm:"sourcepid" json:"sourcepid"`
	SourceName string `xorm:"sourcename" json:"sourcename"` //来源单据名称 可一个单据引入到不同子表
	TargetPid  int    `xorm:"targetpid" json:"targetpid"`
	FlagCol    string `xorm:"flagcol" json:"flagcol"`
	IsOnlyOne  int    `xorm:"isonlyone" json:"isonlyone"` //只能被引用一次
	//上游提交后，下游自动创建
	IsAutoCreateBySourceSubmited int `xorm:"isautocreatebysourcesubmited" json:"isautocreatebysourcesubmited"` //是否自动创建并提交
	//上游终审后，下游自动创建
	IsAutoCreateBySourceApproved int    `xorm:"isautocreatebysourceapproved" json:"isautocreatebysourceapproved"` //是否自动创建并审批
	AutoCreateWhereSql           string `xorm:"autocreatewheresql" json:"autocreatewheresql"`                     //自动创建where条件
	Sourcesplitgridid            int    `xorm:"sourcesplitgridid" json:"sourcesplitgridid"`
	Sourcesplitcols              string `xorm:"sourcesplitcols" json:"sourcesplitcols"`
	Copymemo                     string `xorm:"copymemo" json:"copymemo"`
	IsOpen                       int    `xorm:"isopen" json:"isopen"`
	IsHide                       int    `xorm:"ishide" json:"ishide"` // 是否隐藏，部分拷贝关系只需关联上下游，不需要编辑页引入显示。
	//数据时不进行数据权限过虑，用于可以引入所有人提交的单据
	IsCantUseDaccess int `xorm:"iscantusedaccess" json:"iscantusedaccess"`
	//创建下游后，自动提交下游
	Iscreatebyautosubmit int `xorm:"iscreatebyautosubmit" json:"iscreatebyautosubmit"`
	//下游自动提交时，下游单据创建人 是否取上游单据制单人， 默认取登录人审批人
	IsGetCreate_uidByAutoSubmited int             `xorm:"isgetcreate_uidbyautosubmited" json:"isgetcreate_uidbyautosubmited"`
	Autocreatenext                int             `xorm:"autocreatenext" json:"autocreatenext"` //系统自动创建下游(定时)
	CopyFromList                  *[]Sys_copyfrom `xorm:"-" json:"copyfromlist"`                //数据库不使用，删除xorm标记 - 表示不使用xorm
}

func (*Sys_copy) TableName() string {
	return "sys_copy"
}

//业务显示对象对象
type Sys_copyShow struct {
	Targetname string `json:"targetname"`
	Sourcename string `json:"sourcename"`
	CPID       string `xorm:"cpid" json:"cpid"`
	SourcePid  int    `xorm:"sourcepid" json:"sourcepid"`
	TargetPid  int    `xorm:"targetpid" json:"targetpid"`
	FlagCol    string `xorm:"flagcol" json:"flagcol"`
	IsOnlyOne  int    `xorm:"isonlyone" json:"isonlyone"` //只能被引用一次
	//上游提交后，下游自动创建
	IsAutoCreateBySourceSubmited int `xorm:"isautocreatebysourcesubmited" json:"isautocreatebysourcesubmited"` //是否自动创建并提交
	//上游终审后，下游自动创建
	IsAutoCreateBySourceApproved int    `xorm:"isautocreatebysourceapproved" json:"isautocreatebysourceapproved"` //是否自动创建并审批
	AutoCreateWhereSql           string `xorm:"autocreatewheresql" json:"autocreatewheresql"`                     //自动创建where条件
	Sourcesplitgridid            int    `xorm:"sourcesplitgridid" json:"sourcesplitgridid"`
	Sourcesplitcols              string `xorm:"sourcesplitcols" json:"sourcesplitcols"`
	Copymemo                     string `xorm:"copymemo" json:"copymemo"`
	IsOpen                       int    `xorm:"isopen" json:"isopen"`
	//数据时不进行数据权限过虑，用于可以引入所有人提交的单据
	IsCantUseDaccess int `xorm:"iscantusedaccess" json:"iscantusedaccess"`
	//创建下游后，自动提交下游
	Iscreatebyautosubmit int `xorm:"iscreatebyautosubmit" json:"iscreatebyautosubmit"`
	//下游自动提交时，下游单据创建人 是否取上游单据制单人， 默认取登录人审批人
	IsGetCreate_uidByAutoSubmited int `xorm:"isgetcreate_uidbyautosubmited" json:"isgetcreate_uidbyautosubmited"`
	Autocreatenext                int `xorm:"autocreatenext" json:"autocreatenext"`
}

package restentity

//前台传入的参数
type QueryParam struct {
	Pageindex int    `xorm:"pageindex" json:"pageindex"` //页码
	Pagesize  int    `xorm:"pagesize" json:"pagesize"`   //一页显示行数
	KeyWord   string `xorm:"keyword" json:"keyword"`     //查询关健字 用于不传Field对象的情况
}

package restentity

//查询列表页主数据rest
type TreeVo struct {
	Title      string    `xorm:"title" json:"title"`
	Value      string    `xorm:"value" json:"value"`
	Code       string    `xorm:"code" json:"code"`
	IsHasChild int       `xorm:"ishaschild" json:"ishaschild"`
	Children   []*TreeVo `xorm:"children" json:"children"`
}

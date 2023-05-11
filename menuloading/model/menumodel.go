package model

// 一级菜单
type MenuModel struct {
	ModelID   string
	ModelName string
	SortID    float32
	Child     []*PageModel
	MenuImg   string
}

// 二级菜单
type PageModel struct {
	ParentID  string
	ModelID   string
	ModelName string
	MenuImg   string
	SortID    float32
	Child     []*Page // 三级菜单
}

//三级菜单
type Page struct {
	ParentID string
	Pid      int
	MenuName string
	Ver      int
	MenuUrl  string
	WindType string
	MenuImg  string
	SortID   float32
	OpenType string
}

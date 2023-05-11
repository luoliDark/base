package form

type Sys_fgridbtn struct {
	Rid           int     `xorm:"rid" json:"rid"`
	BarID         string  `xorm:"barid" json:"barid"`
	BarText       string  `xorm:"bartext" json:"bartext"`
	IsSystem      int     `xorm:"issystem" json:"issystem"`
	PageStateShow string  `xorm:"pagestateshow" json:"pagestateshow"`
	Isort         float64 `xorm:"isort" json:"isort"`
}

func (*Sys_fgridbtn) TableName() string {
	return "sys_fgridbtn"
}

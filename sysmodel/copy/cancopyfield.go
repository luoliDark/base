package copy

type Cancopyfield struct {
	Sqlcol     string `xorm:"sqlcol" json:"sqlcol"`
	Name       string `xorm:"name" json:"name"`
	TargetName string `xorm:"targetname" json:"targetname"`
}

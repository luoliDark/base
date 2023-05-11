package design

type Eb_colstaticds struct {
	InfoID  int     `xorm:"infoid" json:"infoid"`
	InfoKey string  `xorm:"infokey" json:"infokey"`
	Pid     int     `xorm:"pid" json:"pid"`
	GridId  int     `xorm:"gridid" json:"gridid"`
	SqlCol  string  `xorm:"sqlcol" json:"sqlcol"`
	SortID  float64 `xorm:"sortid" json:"sortid"`
}

func (*Eb_colstaticds) TableName() string {
	return "eb_colstaticds"
}

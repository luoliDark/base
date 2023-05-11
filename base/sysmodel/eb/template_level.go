package eb

/**
 * @describe:层级档案的level表 的创建模板
 */
type Template_Level_Table struct {
	RefID           string `xorm:"refid" json:"refid"`
	InnerCode       string `xorm:"innercode" json:"innercode"`
	ParentInnerCode string `xorm:"parentinnercode" json:"parentinnercode"`
	IsHasChild      int    `xorm:"ishaschild" json:"ishaschild"`
	ILevel          int    `xorm:"ilevel" json:"ilevel"`
	Level1ID        string `xorm:"level1_id" json:"level1_id"`
	Level2ID        string `xorm:"level2_id" json:"level2_id"`
	Level3ID        string `xorm:"level3_id" json:"level3_id"`
	Level4ID        string `xorm:"level4_id" json:"level4_id"`
	Level5ID        string `xorm:"level5_id" json:"level5_id"`
	Level6ID        string `xorm:"level6_id" json:"level6_id"`
	Level7ID        string `xorm:"level7_id" json:"level7_id"`
	Level8ID        string `xorm:"level8_id" json:"level8_id"`
	Level9ID        string `xorm:"level9_id" json:"level9_id"`
	Level10ID       string `xorm:"level10_id" json:"level10_id"`
	Level11ID       string `xorm:"level11_id" json:"level11_id"`
	Level12ID       string `xorm:"level12_id" json:"level12_id"`
	Level13ID       string `xorm:"level13_id" json:"level13_id"`
	Level14ID       string `xorm:"level14_id" json:"level14_id"`
	Level15ID       string `xorm:"level15_id" json:"level15_id"`
	Entid           int    `xorm:"entid" json:"entid"`
}

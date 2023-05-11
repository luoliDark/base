package eb

/**
 * @Author: weiyg
 * @Date: 2020/3/20 15:12
 * @describe:刷新层级档案的各级levelid专用结构体
 */
type EbRefreshJson struct {
	Code       string `xorm:"Code" json:"Code"`
	Name       string `xorm:"Name" json:"Name"`
	PrimaryKey string `xorm:"PrimaryKey" json:"PrimaryKey"`
	ParentID   string `xorm:"ParentID" json:"ParentID"`

	Innercode       string `xorm:"Innercode" json:"Innercode"`
	FullName        string `xorm:"FullName" json:"FullName"`
	ParentFullName  string `xorm:"ParentFullName" json:"ParentFullName"`
	ParentInnerCode string `xorm:"Code" json:"Code"`
	Level1_ID       string `xorm:"Level1_ID" json:"Level1_ID"`
	Level2_ID       string `xorm:"Level2_ID" json:"Level2_ID"`
	Level3_ID       string `xorm:"Level3_ID" json:"Level3_ID"`
	Level4_ID       string `xorm:"Level4_ID" json:"Level4_ID"`
	Level5_ID       string `xorm:"Level1_ID" json:"Level5_ID"`
	Level6_ID       string `xorm:"Level6_ID" json:"Level6_ID"`
	Level7_ID       string `xorm:"Level7_ID" json:"Level7_ID"`
	Level8_ID       string `xorm:"Level1_ID" json:"Level8_ID"`
	Level9_ID       string `xorm:"Level9_ID" json:"Level9_ID"`
	Level10_ID      string `xorm:"Level10_ID" json:"Level10_ID"`
	Level11_ID      string `xorm:"Level11_ID" json:"Level11_ID"`
	Level12_ID      string `xorm:"Level12_ID" json:"Level12_ID"`
	Level13_ID      string `xorm:"Level13_ID" json:"Level13_ID"`
	Level14_ID      string `xorm:"Level14_ID" json:"Level14_ID"`
	Level15_ID      string `xorm:"Level15_ID" json:"Level15_ID"`

	Level1_Name  string `xorm:"Level1_Name" json:"Level1_Name"`
	Level2_Name  string `xorm:"Level2_Name" json:"Level2_Name"`
	Level3_Name  string `xorm:"Level3_Name" json:"Level3_Name"`
	Level4_Name  string `xorm:"Level4_Name" json:"Level4_Name"`
	Level5_Name  string `xorm:"Level1_Name" json:"Level5_Name"`
	Level6_Name  string `xorm:"Level6_Name" json:"Level6_Name"`
	Level7_Name  string `xorm:"Level7_Name" json:"Level7_Name"`
	Level8_Name  string `xorm:"Level1_Name" json:"Level8_Name"`
	Level9_Name  string `xorm:"Level9_Name" json:"Level9_Name"`
	Level10_Name string `xorm:"Level10_Name" json:"Level10_Name"`
	Level11_Name string `xorm:"Level11_Name" json:"Level11_Name"`
	Level12_Name string `xorm:"Level12_Name" json:"Level12_Name"`
	Level13_Name string `xorm:"Level13_Name" json:"Level13_Name"`
	Level14_Name string `xorm:"Level14_Name" json:"Level14_Name"`
	Level15_Name string `xorm:"Level15_Name" json:"Level15_Name"`
}

package ssomodel

type EB_EntVsUser struct {
	EntID              string `xorm:"entid" json:"entid"`
	Logo               string `xorm:"logo" json:"logo"`
	EntName            string `xorm:"entname" json:"entname"`
	FormEntId          string `xorm:"formentid" json:"formentid"`
	IsOpenOcr          bool   //是否启用OCR识别功能
	SsoTime            int    `xorm:"ssotime" json:"ssotime"`           //登录信息保存时间 -1表示永久，24表示保留一天
	FileView_Ver       int    `xorm:"fileview_ver" json:"fileview_ver"` //登录信息保存时间 -1表示永久，24表示保留一天
	IsCostCenter       bool
	IsVatDetail        bool   //是否启用增值税专票明细数据识别
	IsExWf             bool   //是否使用外部流程系统进行审批 例：外部BPM 外部OA
	IsNotCtrWFNextUser bool   //不强控必须选择下一步审批人
	Mver               string `xorm:"mver" json:"mver"` // 菜单版本号，用于浏览器端缓存时使用
}

func (*EB_EntVsUser) TableName() string {
	return "EB_EntVsUser"
}

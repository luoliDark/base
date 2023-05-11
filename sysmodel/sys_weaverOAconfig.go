package sysmodel

type Sys_WeaverOAConfig struct {
	Config_id     int    `xorm:"config_id" json:"config_id"`
	Entid         int    `xorm:"entid" json:"entid"`
	Entname       string `xorm:"entname" json:"entname"`
	IsOpenSendMsg int    `xorm:"isopensendmsg" json:"isopensendmsg"`
	OutServerIp   string `xorm:"outserverip" json:"outserverip"`
	Code          string `xorm:"code" json:"code"`
	Key           string `xorm:"key" json:"key"` //秘钥
	BasePushUrl   string `xorm:"basepushurl" json:"basepushurl"`
	Messagetypeid string `xorm:"messagetypeid" json:"messagetypeid"`
	Badge         string `xorm:"badge" json:"badge"`
	OAType        string `xorm:"oatype" json:"oatype"` //OA实现类型
	AppID         string `xorm:"appid" json:"appid"`   //OA软件 ID
	Spk           string `xorm:"spk" json:"spk"`       // 非对称加密，公钥
}

func (this *Sys_WeaverOAConfig) TableName() string {
	return "sys_weaverOAconfig"
}

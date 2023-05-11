package sysmodel

type Sys_WeiXinConfig struct {
	Config_id        int    `xorm:"config_id" json:"config_id"`
	Code             string `xorm:"code" json:"code"`
	AgentId          string `xorm:"agentid" json:"agentid"`
	AppKey           string `xorm:"appkey" json:"appkey"`
	AppSecret        string `xorm:"appsecret" json:"appsecret"`
	Accesstokenurl   string `xorm:"accesstokenurl" json:"accesstokenurl"`
	Getuserlisturl   string `xorm:"getuserlisturl" json:"getuserlisturl"`
	Getdeptlisturl   string `xorm:"getdeptlisturl" json:"getdeptlisturl"`
	Sendmsgurl       string `xorm:"sendmsgurl" json:"sendmsgurl"`
	Entid            int    `xorm:"entid" json:"entid"`
	Corpid           string `xorm:"corpid" json:"corpid"`
	CorpSecret       string `xorm:"corpsecret" json:"corpsecret"`
	Getdeptdetailurl string `xorm:"getdeptdetailurl" json:"getdeptdetailurl"`
	Is_thirdparty    int    `xorm:"is_thirdparty" json:"is_thirdparty"`
	MsgServerUrl     string `xorm:"msgserverurl" json:"msgserverurl"`
	CallbackToken    string `xorm:"callbacktoken" json:"callbacktoken"`
	CallbackKey      string `xorm:"callbackkey" json:"callbackkey"`
	SuiteKey         string `xorm:"suitekey" json:"suitekey"`
	SuiteSecret      string `xorm:"suitesecret" json:"suitesecret"`
	OutServerIp      string `xorm:"outserverip" json:"outserverip"`
}

func (this *Sys_WeiXinConfig) TableName() string {
	return "sys_weixinconfig"
}
func (this *Sys_WeiXinConfig) GetAgentId() string {
	return this.AgentId
}
func (this *Sys_WeiXinConfig) GetAppKey() string {
	return this.AppKey
}
func (this *Sys_WeiXinConfig) GetAppSecret() string {
	return this.AppSecret
}
func (this *Sys_WeiXinConfig) GetAccesstokenurl() string {
	return this.Accesstokenurl
}
func (this *Sys_WeiXinConfig) GetGetuserlisturl() string {
	return this.Getuserlisturl
}

func (this *Sys_WeiXinConfig) GetGetdeptlisturl() string {
	return this.Getdeptlisturl
}

func (this *Sys_WeiXinConfig) GetEntid() int {
	return this.Entid
}

func (this *Sys_WeiXinConfig) GetCorpSecret() string {
	return this.CorpSecret
}

func (this *Sys_WeiXinConfig) GetCorpid() string {
	return this.Corpid
}
func (this *Sys_WeiXinConfig) GetGetdeptdetailurl() string {
	return this.Getdeptdetailurl
}
func (this *Sys_WeiXinConfig) GetIs_thirdparty() int {
	return this.Is_thirdparty
}
func (this *Sys_WeiXinConfig) GetSendMsgUrl() string {
	return this.Sendmsgurl
}
func (this *Sys_WeiXinConfig) GetMsgServerUrl() string {
	return this.MsgServerUrl
}

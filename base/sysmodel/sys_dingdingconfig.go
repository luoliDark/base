package sysmodel

type Sys_dingdingConfig struct {
	Config_id               int    `xorm:"config_id" json:"config_id"`
	Code                    string `xorm:"code" json:"code"`
	AgentId                 string `xorm:"agentid" json:"agentid"`
	AppKey                  string `xorm:"appkey" json:"appkey"`
	AppSecret               string `xorm:"appsecret" json:"appsecret"`
	Accesstokenurl          string `xorm:"accesstokenurl" json:"accesstokenurl"`
	Getuserlisturl          string `xorm:"getuserlisturl" json:"getuserlisturl"`
	Getdeptlisturl          string `xorm:"getdeptlisturl" json:"getdeptlisturl"`
	Sendmsgurl              string `xorm:"sendmsgurl" json:"sendmsgurl"`
	Entid                   int    `xorm:"entid" json:"entid"`
	Corpid                  string `xorm:"corpid" json:"corpid"`
	Getdeptdetailurl        string `xorm:"getdeptdetailurl" json:"getdeptdetailurl"`
	Is_thirdparty           int    `xorm:"is_thirdparty" json:"is_thirdparty"`
	MsgServerUrl            string `xorm:"msgserverurl" json:"msgserverurl"`
	CallbackToken           string `xorm:"callbacktoken" json:"callbacktoken"`
	CallbackKey             string `xorm:"callbackkey" json:"callbackkey"`
	SuiteKey                string `xorm:"suitekey" json:"suitekey"`
	SuiteSecret             string `xorm:"suitesecret" json:"suitesecret"`
	OutServerIp             string `xorm:"outserverip" json:"outserverip"`
	Instancecreateurl       string `xorm:"instancecreateurl" json:"instancecreateurl"`             //创建实例url
	Taskcreateurl           string `xorm:"taskcreateurl" json:"taskcreateurl"`                     //创建代办url
	Templatecreateurl       string `xorm:"templatecreateurl" json:"templatecreateurl"`             //创建模板url
	UpdateInstancecreateurl string `xorm:"updateinstancecreateurl" json:"updateinstancecreateurl"` //修改实例状态url
	UpdateTaskcreateurl     string `xorm:"updatetaskcreateurl" json:"updatetaskcreateurl"`         //修改待办状态url
	Rolelist                string `xorm:"rolelist" json:"rolelist"`                               //获取角色列表url
	Rolesimplelist          string `xorm:"rolesimplelist" json:"rolesimplelist"`                   //获取角色下人员列表url
}

func (this *Sys_dingdingConfig) TableName() string {
	return "sys_dingdingconfig"
}
func (this *Sys_dingdingConfig) GetAgentId() string {
	return this.AgentId
}
func (this *Sys_dingdingConfig) GetAppKey() string {
	return this.AppKey
}
func (this *Sys_dingdingConfig) GetAppSecret() string {
	return this.AppSecret
}
func (this *Sys_dingdingConfig) GetAccesstokenurl() string {
	return this.Accesstokenurl
}
func (this *Sys_dingdingConfig) GetGetuserlisturl() string {
	return this.Getuserlisturl
}

func (this *Sys_dingdingConfig) GetGetdeptlisturl() string {
	return this.Getdeptlisturl
}

func (this *Sys_dingdingConfig) GetEntid() int {
	return this.Entid
}

func (this *Sys_dingdingConfig) GetCorpid() string {
	return this.Corpid
}
func (this *Sys_dingdingConfig) GetGetdeptdetailurl() string {
	return this.Getdeptdetailurl
}
func (this *Sys_dingdingConfig) GetIs_thirdparty() int {
	return this.Is_thirdparty
}
func (this *Sys_dingdingConfig) GetSendMsgUrl() string {
	return this.Sendmsgurl
}
func (this *Sys_dingdingConfig) GetMsgServerUrl() string {
	return this.MsgServerUrl
}

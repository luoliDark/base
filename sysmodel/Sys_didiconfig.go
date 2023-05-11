package sysmodel

type Sys_didiconfig struct {
	Config_id          int    `xorm:"config_id" json:"config_id"`
	Entid              int    `xorm:"entid" json:"entid"`
	Entname            string `xorm:"entname" json:"entname"`
	IsOpenSendMsg      int    `xorm:"isOpenSendMsg" json:"isOpenSendMsg"`
	OutServerIp        string `xorm:"outServerIp" json:"outServerIp"`
	Code               string `xorm:"code" json:"code"`
	ClientID           string `xorm:"clientID" json:"clientID"`
	ClientSecret       string `xorm:"clientSecret" json:"clientSecret"`
	SignKey            string `xorm:"signKey" json:"signKey"`
	Phone              string `xorm:"phone" json:"phone"`
	Companyid          string `xorm:"companyid" json:"companyid"`
	Accesstokenurl     string `xorm:"accesstokenurl" json:"accesstokenurl"`
	AddUserurl         string `xorm:"addUserurl" json:"addUserurl"`
	UpdateUserurl      string `xorm:"updateUserurl" json:"updateUserurl"`
	Deleteuserurl      string `xorm:"deleteuserurl" json:"deleteuserurl"`
	AddDepturl         string `xorm:"addDepturl" json:"addDepturl"`
	UpdateDepturl      string `xorm:"updateDepturl" json:"updateDepturl"`
	Deletedepturl      string `xorm:"deletedepturl" json:"deletedepturl"`
	GetInvoiceurl      string `xorm:"getInvoiceurl" json:"getInvoiceurl"`
	Msgserverurl       string `xorm:"msgserverurl" json:"msgserverurl"`
	Callbacktoken      string `xorm:"callbacktoken" json:"callbacktoken"`
	DownloadInvoiceurl string `xorm:"downloadinvoiceurl" json:"downloadinvoiceurl"` //下载滴滴发票地址
}

func (this *Sys_didiconfig) TableName() string {
	return "sys_didiconfig"
}

func (this *Sys_didiconfig) GetAccesstokenurl() string {
	return this.Accesstokenurl
}

func (this *Sys_didiconfig) GetConfig_id() int {
	return this.Config_id
}

func (this *Sys_didiconfig) GetEntid() int {
	return this.Entid
}

func (this *Sys_didiconfig) GetClientID() string {
	return this.ClientID
}

func (this *Sys_didiconfig) GetClientSecret() string {
	return this.ClientSecret
}

func (this *Sys_didiconfig) GetPhone() string {
	return this.Phone
}

func (this *Sys_didiconfig) GetSignKey() string {
	return this.SignKey
}

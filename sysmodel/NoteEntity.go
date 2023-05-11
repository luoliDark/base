package sysmodel

/**
 * 通知对象
 */
type NoteEntity struct {
	MessageCode string `json:"messageCode"`
	SrcSystem   string `json:"srcSystem"`
	BusiType    string `json:"busiType"`
	OrderNum    string `json:"orderNum"`
	OptType     string `json:"optType"`
	Url         string `json:"url"`
	TenantId    int    `json:"tenantId"`
	//Messge Messge `json:"message"`
	ReceiverAddressList []ReceiverAddress `json:"receiverAddressList"`
	EntId               int               `json:entId`
	Args                map[string]string `json:"args"`
}

/**
 * 消息对象
 */
type Messge struct {
	subject      string `json:"messageCode"`
	content      string `json:"messageCode"`
	externalCode string `json:"messageCode"`
}

/**
 * 接收人信息
 */
type ReceiverAddress struct {
	TargetUserTenantId int               `json:"targetUserTenantId"`
	AdditionInfo       map[string]string `json:"additionInfo"`
}

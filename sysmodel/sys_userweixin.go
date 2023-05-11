package sysmodel

type Sys_UserWeiXin struct {
	Id                int    `xorm:"id" json:"id"`
	Userid            string `xorm:"userid" json:"userid"`
	Usercode          string `xorm:"usercode" json:"usercode"`
	Name              string `xorm:"name" json:"name"`
	Mobile            string `xorm:"mobile" json:"mobile"`
	Department        string `xorm:"department" json:"department"`
	Deptid            string `xorm:"deptid" json:"deptid"`
	Orders            string `xorm:"orders" json:"orders"`
	Entid             string `xorm:"entid" json:"entid"`
	Position          string `xorm:"position" json:"position"`
	Gender            string `xorm:"gender" json:"gender"`
	Email             string `xorm:"email" json:"email"`
	Is_leader_in_dept string `xorm:"is_leader_in_dept" json:"is_leader_in_dept"`
	Avatar            string `xorm:"avatar" json:"avatar"`
	Thumb_avatar      string `xorm:"thumb_avatar" json:"thumb_avatar"`
	Telephone         string `xorm:"telephone" json:"telephone"`
	Alias             string `xorm:"alias" json:"alias"`
	Extattr           string `xorm:"extattr" json:"extattr"`
	Status_           string `xorm:"status_" json:"status_"`
	Qr_code           string `xorm:"qr_code" json:"qr_code"`
	External_profile  string `xorm:"external_profile" json:"external_profile"`
	External_position string `xorm:"external_position" json:"external_position"`
	Address           string `xorm:"address" json:"address"`
	Open_userid       string `xorm:"open_userid" json:"open_userid"`
	Main_department   string `xorm:"main_department" json:"main_department"`
	Create_time       string `xorm:"create_time" json:"create_time"`
}

func (this *Sys_UserWeiXin) TableName() string {
	return "sys_userweixin"
}

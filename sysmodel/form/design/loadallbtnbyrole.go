package design

/**
 * @Author: lvxuanye
 * @Date: 2020/5/3 08:29
 * @describe:取出所有菜单
 */
type Loadallbtnbyrole struct {
	BtnVSFormID string `xorm:"BtnVSFormID" json:"BtnVSFormID"`
	ModelID     int    `xorm:"ModelID" json:"ModelID"`
	Modelname   string `xorm:"modelname" json:"modelname"`
	Pid         int    `xorm:"Pid" json:"Pid"`
	Pname       string `xorm:"Pname" json:"Pname"`
	Btncode     string `xorm:"Btncode" json:"Btncode"`
	BtnText     string `xorm:"BtnText" json:"BtnText"`
	Isenable    int    `xorm:"isenable" json:"isenable"`
	IsEditPage  int    `xorm:"IsEditPage" json:"IsEditPage"`
}

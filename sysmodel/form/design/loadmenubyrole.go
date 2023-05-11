package design

/**
 * @Author: lvxuanye
 * @Date: 2020/5/3 08:29
 * @describe:取出所有菜单
 */
type Loadmenubyrole struct {
	Modelid   int    `xorm:"modelid" json:"modelid"`
	Modelname string `xorm:"modelname" json:"modelname"`
	Pname     string `xorm:"pname" json:"pname"`
	Pid       int    `xorm:"Pid" json:"Pid"`
	Isenable  int    `xorm:"isenable" json:"isenable"`
}

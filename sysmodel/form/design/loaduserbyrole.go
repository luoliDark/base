package design

/**
 * @Author: lvxuanye
 * @Date: 2020/4/29
 * @describe:查看审批信息 -结构体
 */
type Loaduserbyrole struct {
	Userid   string `xorm:"Userid" json:"Userid"`
	Usercode string `xorm:"Usercode" json:"Usercode"`
	Username string `xorm:"Username" json:"Username"`
	ImgSrc   string `xorm:"ImgSrc" json:"ImgSrc"`
	DeptCode string `xorm:"DeptCode" json:"DeptCode"`
	DeptName string `xorm:"DeptName" json:"DeptName"`
}

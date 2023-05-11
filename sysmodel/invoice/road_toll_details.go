package invoice

type Road_toll_details struct {
	Id           string  `xorm:"id" json:"id"`
	InvoiceCode  string  `xorm:"invoice_code" json:"invoice_code"`   // 发票代码
	InvoiceNo    string  `xorm:"invoice_no" json:"invoice_no"`       // 发票号码
	Date         string  `xorm:"date" json:"date"`                   // 日期
	Time         string  `xorm:"time" json:"time"`                   // 时间
	AmountLittle float64 `xorm:"amount_little" json:"amount_little"` // 合计人民币（小写）
	Exit         string  `xorm:"exit" json:"exit"`                   // 出口
	Entrance     string  `xorm:"entrance" json:"entrance"`           // 入口
	CreateDate   string  `xorm:"create_date" json:"create_date"`     // 创建时间
	Billno       string  `xorm:"billno" json:"billno"`               // 单号
	IsSend       int     `xorm:"is_send" json:"is_send"`             // 是否已发送到外部系统 0：未发送 1：已发送
	Pid          string  `xorm:"pid" json:"pid"`                     // pid
	BillType     string  `xorm:"bill_type" json:"bill_type"`         // 单据类型
	AttachmentId int     `xorm:"attachment_id" json:"attachment_id"` // 影像图片id
	CompanyId    int     `xorm:"company_id" json:"company_id"`       // 公司ID
	Userid       string  `xorm:"userid" json:"userid"`
}

func (*Road_toll_details) TableName() string {
	return "road_toll_details"
}

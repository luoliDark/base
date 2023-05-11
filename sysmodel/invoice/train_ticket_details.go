package invoice

type Train_ticket_details struct {
	Id             string  `xorm:"id" json:"id"`
	InvoiceNo      string  `xorm:"invoice_no" json:"invoice_no"`           // 票号
	DateTime       string  `xorm:"date_time" json:"date_time"`             // 发车时间
	AmountLittle   float64 `xorm:"amount_little" json:"amount_little"`     // 票价
	TaxAmount      float64 `xorm:"tax_amount" json:"tax_amount"`           // 税额
	NoTaxAmount    float64 `xorm:"no_tax_amount" json:"no_tax_amount"`     // 不含税金额
	TrainNo        string  `xorm:"train_no" json:"train_no"`               // 车次
	StationFrom    string  `xorm:"station_from" json:"station_from"`       // 出发站
	StationTo      string  `xorm:"station_to" json:"station_to"`           // 到达站
	Seat           string  `xorm:"seat" json:"seat"`                       // 座位号
	SeatClass      string  `xorm:"seat_class" json:"seat_class"`           // 座位类型
	EntranceInfo   string  `xorm:"entrance_info" json:"entrance_info"`     // 进站信息
	PurchaseType   string  `xorm:"purchase_type" json:"purchase_type"`     // 购票方式
	IdName         string  `xorm:"id_name" json:"id_name"`                 // 身份证号码姓名
	IssuingNo      string  `xorm:"issuing_no" json:"issuing_no"`           // 售票信息
	IssuingStation string  `xorm:"issuing_station" json:"issuing_station"` // 售票站
	CreateDate     string  `xorm:"create_date" json:"create_date"`         // 创建时间
	Billno         string  `xorm:"billno" json:"billno"`                   // 单据号
	IsSend         int     `xorm:"is_send" json:"is_send"`                 // 是否已发送到外部系统 0：未发送 1：已发送
	Pid            string  `xorm:"pid" json:"pid"`                         // pid
	BillType       string  `xorm:"bill_type" json:"bill_type"`             // 单据类型
	AttachmentId   int     `xorm:"attachment_id" json:"attachment_id"`     // 影像图片id
	CompanyId      int     `xorm:"company_id" json:"company_id"`           // 公司ID
	Userid         string  `xorm:"userid" json:"userid"`
}

func (*Train_ticket_details) TableName() string {
	return "train_ticket_details"
}

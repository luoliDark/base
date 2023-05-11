package invoice

type Air_ticket_details struct {
	Id                  string  `xorm:"id" json:"id"`
	Insurance           float64 `xorm:"insurance" json:"insurance"`                         // 保险费
	Fare                float64 `xorm:"fare" json:"fare"`                                   // 票价
	IssuedBy            string  `xorm:"issued_by" json:"issued_by"`                         // 填开单位
	Name                string  `xorm:"name" json:"name"`                                   // 旅客姓名
	NameId              string  `xorm:"name_id" json:"name_id"`                             // 有效身份证件号码
	CaacDevelopmentFund float64 `xorm:"caac_development_fund" json:"caac_development_fund"` // 民航发展基金
	Date                string  `xorm:"date" json:"date"`                                   // 日期
	Time                string  `xorm:"time" json:"time"`                                   // 时间
	FlightNo            string  `xorm:"flight_no" json:"flight_no"`                         // 航班号
	FareBasis           string  `xorm:"fare_basis" json:"fare_basis"`                       // 客票级别
	From                string  `xorm:"from" json:"from"`                                   // 出发地
	SeatClass           string  `xorm:"seat_class" json:"seat_class"`                       // 座位等级
	To                  string  `xorm:"to" json:"to"`                                       // 目的地
	ETicketNo           string  `xorm:"e_ticket_no" json:"e_ticket_no"`                     // 电子客票号码
	FuelSurcharge       float64 `xorm:"fuel_surcharge" json:"fuel_surcharge"`               // 燃油附加费
	OtherTaxes          float64 `xorm:"other_taxes" json:"other_taxes"`                     // 其他税费
	AmountLittle        float64 `xorm:"amount_little" json:"amount_little"`                 // 合计票价
	TaxAmount           float64 `xorm:"tax_amount" json:"tax_amount"`                       // 税额
	NoTaxAmount         float64 `xorm:"no_tax_amount" json:"no_tax_amount"`                 // 不含税金额
	DateOfIssue         string  `xorm:"date_of_issue" json:"date_of_issue"`                 // 填开日期
	AgentCode           string  `xorm:"agent_code" json:"agent_code"`                       // 销售单位代号
	Billno              string  `xorm:"billno" json:"billno"`                               // 单据号
	IsSend              int     `xorm:"is_send" json:"is_send"`                             // 是否已发送到外部系统 0：未发送 1：已发送
	Pid                 string  `xorm:"pid" json:"pid"`                                     // pid
	BillType            string  `xorm:"bill_type" json:"bill_type"`                         // 单据类型
	CompanyId           int     `xorm:"company_id" json:"company_id"`                       // 公司id
	AttachmentId        int     `xorm:"attachment_id" json:"attachment_id"`                 // 影像图片ID
	CreateDate          string  `xorm:"create_date" json:"create_date"`                     // 创建时间
	Userid              string  `xorm:"userid" json:"userid"`
}

func (*Air_ticket_details) TableName() string {
	return "air_ticket_details"
}

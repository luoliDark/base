package file

type Ocr_sys_fuploadfile struct {
	Fid            string `xorm:"fid" json:"fid"`
	Uid            string `xorm:"uid" json:"uid"`                       // 上传人id
	Filepath       string `xorm:"filepath" json:"filepath"`             // 文件路径
	Billid         string `xorm:"billid" json:"billid"`                 // 单据号
	FileSize       int    `xorm:"filesize" json:"filesize"`             // 文件大小
	OriginFileName string `xorm:"originfilename" json:"originfilename"` // 原始文件名
	Memo           string `xorm:"memo" json:"memo"`                     // 备注
	NewGuid        string `xorm:"newguid" json:"newguid"`
	GridID         int    `xorm:"gridid" json:"gridid"`             // 表单id
	GridDetailID   string `xorm:"griddetailid" json:"griddetailid"` // 表单子表id
	CreateDate     string `xorm:"createdate" json:"createdate"`     // 创建时间
	FormType       string `xorm:"formtype" json:"formtype"`
	OriginPath     string `xorm:"originpath" json:"originpath"`
	FileName       string `xorm:"filename" json:"filename"`     // 生成的文件名称
	FileType       string `xorm:"filetype" json:"filetype"`     // 文件类型
	AuthcType      string `xorm:"authctype" json:"authctype"`   // 查看权限类型
	CompanyId      int    `xorm:"company_id" json:"company_id"` // 公司id
	Pid            int    `xorm:"pid" json:"pid"`
	Entid          int    `xorm:"entid" json:"entid"`
	IsDiscard      int    `xorm:"isdiscard" json:"isdiscard"`
	DiscardUid     string `xorm:"discard_uid" json:"discard_uid"`
	DiscardDate    string `xorm:"discard_date" json:"discard_date"`
	Bbsid          string `xorm:"bbsid" json:"bbsid"`
}

func (*Ocr_sys_fuploadfile) TableName() string {
	return "ocr_sys_fuploadfile"
}

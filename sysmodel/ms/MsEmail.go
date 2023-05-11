package ms

type MsEmail struct {
	Pid        int      //单据类型
	PrimaryKey string   //主健
	To         []string `json:"to"`
	Body       string   `json:"body"`
	Title      string   `json:"title"`
}

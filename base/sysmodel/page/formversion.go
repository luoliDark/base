package page

/*
jsz by 2020.3.2 单据版本号实体
*/

type ReturnGrid struct {
	GridId       int
	AddDetailPks map[string]string
}

type FormVersion struct {
	PrimaryKey    string
	Version       string
	BillNo        string
	FlowStatus    int
	IsRefresh     int
	IsTemporary   bool
	AddRowsResult []*ReturnGrid
}

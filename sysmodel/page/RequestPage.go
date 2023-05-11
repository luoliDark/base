package page

//前端传参格式

type GridParms struct {
	GridID       int    //子表id
	GridDetailId string //子表流水ID
	Offset       int    //当前页码
	Limit        int    //当前页面容量
}

type MainParms struct {
	Pid          int
	PrimaryKeyId string
	Offset       int //当前页码
	Limit        int //当前页面容量
}

/**
编辑页查询返回格式
*/
type EditPageData struct {
	MainPageData MainPageData   // 主表数据
	GridPageData []GridPageData //子表数据
}

type MainPageData struct {
	Pid        int    //主表id
	ResultData string //主表的数据
}

type GridPageData struct {
	GridID     int    //子表id
	ResultData string //子表的数据
}

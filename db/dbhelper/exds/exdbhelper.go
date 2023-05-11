//jsa by 2020.2.6 用于读取外部数据源

package exds

//查询外部数据库
func QueryByEXDB(userID string, exDBID int, Args ...interface{}) ([]map[string]string, error) {

	//todo
	var re []map[string]string
	return re, nil
}

//分页查询外部数据
func QueryPagingByEXDB(userID string, exDBID int, pageIndex int, pageSize int, Args ...interface{}) (lst []map[string]string, rowCnt int, err error) {
	//todo
	return lst, 0, nil
}

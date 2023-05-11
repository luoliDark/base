package sysmodel

type DataEntity struct {
	// 记录类型
	rowType string
	// data 数据格式
	dataJson string
	// xml 数据格式
	dataXml string
	//每个表的行数,必填
	rownum int
	// listmap 数据格式
	dataMapList []map[string]interface{}
	// 表名 可为空
	tableName string
	//比表gridid ，只有子表有，且子表必须有值
	gridID string
	// 对应子表数据集，一个或多个
	childDataEntity []interface{}
	//list<Object> 格式的数据集
	listObj []interface{}

	isSelectByPaging bool
	field_json       string
	//子表主键字段
	tablePkName string
}

func (this *DataEntity) GetdataMapList() []map[string]interface{} {
	//todo
	return this.dataMapList
}
func (this *DataEntity) SetdataMapList(dataMapList []map[string]interface{}) {
	//todo
	this.dataMapList = dataMapList
}

func (this *DataEntity) SettableName(tableName string) {
	//todo
	this.tableName = tableName
}

func (this *DataEntity) SetgridID(gridID string) {
	//todo
	this.gridID = gridID
}

func (this *DataEntity) GetgridID() (gridID string) {
	//todo
	return this.gridID
}

func (this *DataEntity) SetchildDataEntity(childDataEntity []interface{}) {
	//todo
	this.childDataEntity = childDataEntity
}

func (this *DataEntity) GetchildDataEntity() (childDataEntity []interface{}) {
	//todo
	return this.childDataEntity
}

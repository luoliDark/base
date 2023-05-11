package sysmodel

type ExcelEntity struct {
	ColIndex        int    //这一列在EXCEL中的列索引
	titleName       string //标题名     原值=部门。按编码转换，标题名+编码 ， 例 ：部门编码。按名称转换，标题名+名称， 例：部门名称
	value           string //Excel中的值
	sqlCol          string //数据库列名
	sqlDataType     string //数据库字符类型
	dataSource      string //数据源  Pid 或者 9999_pid两种，9999开头，只能按名称为key值进行比对
	isConvertByCode bool   //按编码转换。数据源查找按编码为key值
	isConvertByName bool   //按名称转换。数据源查找按名称为key值
	isRequired      bool   //是否必输。True，验证是否value为空
	isMain          bool   //是否主表
	isPidSource     bool   //是否为 Pid型数据源
	isSingle        bool   //是否数据源单选
	showname        string
}

func (this *ExcelEntity) SetTitle(titleName string) {
	this.titleName = titleName
}

func (this *ExcelEntity) SetValue(value string) {
	this.value = value
}

func (this *ExcelEntity) SetSqlCol(sqlCol string) {
	this.sqlCol = sqlCol
}

func (this *ExcelEntity) SetSqlDataType(sqlDataType string) {
	this.sqlDataType = sqlDataType
}
func (this *ExcelEntity) SetDataSource(dataSource string) {
	this.dataSource = dataSource
}

func (this *ExcelEntity) SetIsConvertByCode(isConvertByCode bool) {
	this.isConvertByCode = isConvertByCode
}

func (this *ExcelEntity) SetIsConvertByName(isConvertByName bool) {
	this.isConvertByName = isConvertByName
}
func (this *ExcelEntity) SetIsRequired(isRequired bool) {
	this.isRequired = isRequired
}
func (this *ExcelEntity) SetIsMain(isMain bool) {
	this.isMain = isMain
}
func (this *ExcelEntity) SetIsPidSource(isPidSource bool) {
	this.isPidSource = isPidSource
}
func (this *ExcelEntity) SetIsSingle(isSingle bool) {
	this.isSingle = isSingle
}
func (this *ExcelEntity) SetShowname(showname string) {
	this.showname = showname
}

func (this *ExcelEntity) GetTitle() string {
	return this.titleName
}

func (this *ExcelEntity) GetValue() string {
	return this.value
}

func (this *ExcelEntity) GetSqlCol() string {
	return this.sqlCol
}

func (this *ExcelEntity) GetSqlDataType() string {
	return this.sqlDataType
}
func (this *ExcelEntity) GetDataSource() string {
	return this.dataSource
}

func (this *ExcelEntity) GetIsConvertByCode() bool {
	return this.isConvertByCode
}
func (this *ExcelEntity) GetIsConvertByName() bool {
	return this.isConvertByName
}
func (this *ExcelEntity) GetIsRequired() bool {
	return this.isRequired
}
func (this *ExcelEntity) GetIsMain() bool {
	return this.isMain
}
func (this *ExcelEntity) GetIsPidSource() bool {
	return this.isPidSource
}
func (this *ExcelEntity) GetIsSingle() bool {
	return this.isSingle
}
func (this *ExcelEntity) GetShowname() string {
	return this.showname
}

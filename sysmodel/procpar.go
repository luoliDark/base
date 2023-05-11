package sysmodel

// 存储过程参数对象
type ProcPar struct {
	ProcName   string
	DataType   string
	DataLength int
	IsOutPut   bool
	ProcValue  string
}

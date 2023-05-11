package sysmodel

type Eb_deptdingtalk struct {
	id              int    `xorm:"id" json:"id"`
	createDeptGroup string `xorm:"createdeptgroup" json:"createdeptgroup"`
	name            string `xorm:"name" json:"name"`
	dept_id         string `xorm:"dept_id" json:"dept_id"`
	autoAddUser     string `xorm:"autoadduser" json:"autoadduser"`
	parent_depte_id string `xorm:"parent_depte_id" json:"parent_depte_id"`
	create_time     string `xorm:"create_time" json:"create_time"`
}

func (this *Eb_deptdingtalk) SetCreateDeptGroup(createDeptGroup string) {
	this.createDeptGroup = createDeptGroup
}
func (this *Eb_deptdingtalk) SetName(name string) {
	this.name = name
}
func (this *Eb_deptdingtalk) SetDept_id(deptid string) {
	this.dept_id = deptid
}
func (this *Eb_deptdingtalk) SetAutoAddUser(autoAddUser string) {
	this.autoAddUser = autoAddUser
}
func (this *Eb_deptdingtalk) SetParent_depte_id(parent_depte_id string) {
	this.parent_depte_id = parent_depte_id
}

func (*Eb_deptdingtalk) TableName() string {
	return "eb_deptdingtalk"
}

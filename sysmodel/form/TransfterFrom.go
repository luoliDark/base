package form

import "paas/base/sysmodel/form/design"

type TransfterFrom struct {
	Sys_fpageid     []Sys_fpageid
	Sys_fpage       []Sys_fpage
	Sys_fpagefield  []Sys_fpagefield
	Sys_fgrid       []Sys_fgrid
	Sys_fgridfield  []Sys_fgridfield
	Sys_fbtnvsform  []Sys_fbtnvsform
	Sys_fbtnaccess  []Sys_fbtnaccess
	Sys_fpageaccess []Sys_fpageaccess
	Sys_uiform      []design.Sys_uiform
	Sys_uiformcol   []design.Sys_uiformcol
	Sys_uiformgrid  []design.Sys_uiformgrid
	Sys_uitabs      []design.Sys_uitabs
	Sys_uitabpage   []design.Sys_uitabpage
	Sys_fpagever    []Sys_fpagever
}

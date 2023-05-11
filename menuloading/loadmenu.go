package menuloading

import (
	"base/db/dbhelper"
	"base/menuloading/model"
	"base/sysmodel"
	"base/util/commutil"
	"base/util/jsonutil"

	"github.com/gogf/gf/frame/g"
)

// create by zhongxinjian 2020年3月6日21:38:13

type pageuilevel struct {
	pid      string
	uitempid string
	level    int
}

//根据角色查询菜单并返回json
//注：从redis中根据角色判定读取 redis中存每个角色对应的菜单明细
func QueryMenuMainByRedis(token string, enterPriseId string, isBack bool) (result sysmodel.ResultBean) {
	////todo 菜单加载及redis加载规范讨论
	result = sysmodel.ResultBean{}
	return sysmodel.ResultBean{}
	//user, err := sso.GetUserByToken(token)
	//
	//if err != nil {
	//	// 用户为空
	//	result.SetError(errcode.LGERR_ACCEXPIRE, "用户token过期，请重新登录", "")
	//	return result
	//}
	//var backflag = "0"
	//if isBack {
	//	backflag = "1"
	//}
	//// 存储pid 对应模板信息，需要对比level值，如果level值大的保留
	//pageMap := make(map[string]pageuilevel)
	//// 先组装当前用户角色有多少菜单，然后根据pid，uitempid 去重,然后把level大的uitempid 保留再去查找表单数据
	//for roleid, _ := range user.UserRoleIds {
	//	// 根据roleid 获取pid_tempid_levelid
	//	values := redishelper.GetList(enterPriseId, confighelper.GetCacheDbIndex(), "sys_fpageaccess_"+roleid+"_"+backflag)
	//	for _, val := range values {
	//		// 将val 分割
	//		arr := strings.Split(val, "_")
	//		pidStr := arr[0]
	//		uitempid := arr[1]
	//		var level int
	//		if uitempid == "" {
	//			level = 0
	//		} else {
	//			if arr[2] == "" {
	//				level = 0
	//			}
	//		}
	//		if _, ok := pageMap[pidStr]; ok {
	//			// 如果有存在过pid表单数据
	//			if pageMap[pidStr].level < level {
	//				// 如果包含pid 则对比旧的level 如果新的level比较大，则用新的否则不变
	//				vg := pageMap[pidStr]
	//				vg.level = level
	//				vg.pid = pidStr
	//				vg.uitempid = uitempid
	//				pageMap[pidStr] = vg
	//			}
	//		} else {
	//			// 如果之前没有存在过表单数据，则设置新的
	//			vg := pageuilevel{uitempid: uitempid, level: level, pid: pidStr}
	//			pageMap[pidStr] = vg
	//		}
	//	}
	//}
	//// pid 已经是有序的
	//// 遍历pid map 逐个查询数据
	//MenuList := make([]*model.MenuModel, 0)
	//var lastlevel1id = "" // 上一个一级模块id
	//var lastlevel2id = "" // 上一个二级模块id
	//var level1index = 0   // 一级菜单当前下标
	//var level2index = 0   // 二级菜单当前下标
	//for _, value := range pageMap {
	//	// 取得了pid 和 uitempid 之后，然后从redis 获取表单信息
	//	pageInfo := redishelper.GetHashMap(enterPriseId, confighelper.GetCacheDbIndex(),
	//		"sys_fpagelevel_"+value.pid+"_"+value.uitempid)
	//	if pageInfo["listhtmlpath"] != "" {
	//		// 不为空
	//		pageInfo["MenuUrl"] = pageInfo["listhtmlpath"] + "&uitempid=" + pageInfo["uitempid"]
	//	}
	//	level1id := pageInfo["level1id"]
	//	level2id := pageInfo["level2id"]
	//	if (lastlevel1id == "" && lastlevel2id == "") || (lastlevel1id != level1id && lastlevel2id != level2id) {
	//		// 是第一个
	//		pageList := make([]map[string]string, 0)
	//		pageList = append(pageList, pageInfo)
	//		pagemodel := model.PageModel{ModelName: pageInfo["level2name"], ModelID: pageInfo["level2id"], ParentID: level1id, MenuImg: pageInfo["menuimg"], Child: pageList}
	//		level2List := make([]*model.PageModel, 0)
	//		level2List = append(level2List, &pagemodel)
	//		MenuList = append(MenuList, &model.MenuModel{ModelID: level1id, ModelName: pageInfo["level1name"], MenuImg: pageInfo["level1img"], Child: level2List})
	//		if lastlevel2id == "" && lastlevel1id == "" {
	//			// 第一个 下标都为0
	//			level1index = 0
	//		} else { // 否则下标增加
	//			level1index++
	//		}
	//		// 二级菜单下标需要归零
	//		level2index = 0
	//		lastlevel1id = level1id
	//		lastlevel2id = level2id
	//	} else {
	//		// 已经存在一级菜单
	//		// lastleve1id 是否和当前一致
	//		if lastlevel1id == level1id {
	//			// 在同一个一级模块下
	//			if lastlevel2id == level2id {
	//				// 是否在同一个二级模块下
	//				level2Menu := MenuList[level1index].Child[level2index]
	//				// 只需要添加三级菜单
	//				// 已经存在的三级菜单列表
	//				pageList := level2Menu.Child
	//				pageList = append(pageList, pageInfo)
	//				level2Menu.Child = pageList
	//			} else {
	//				// 不在同一个二级模块
	//				// 重新创建二级模块列表
	//				level2List := MenuList[level1index].Child
	//				pageList := make([]map[string]string, 0)
	//				pageList = append(pageList, pageInfo)
	//				pagemodel := model.PageModel{ModelName: pageInfo["level2name"], ModelID: pageInfo["level2id"], ParentID: level1id, MenuImg: pageInfo["menuimg"], Child: pageList}
	//				level2List = append(level2List, &pagemodel)
	//				MenuList[level1index].Child = level2List
	//				//MenuList = append(MenuList, &model.MenuModel{ModelID: level1id, ModelName: pageInfo["level1name"], MenuImg: pageInfo["level1img"], Child: level2List})
	//				level2index++
	//				lastlevel2id = level2id
	//			}
	//		}
	//	}
	//}
	//json, err := jsonutil.ObjToJson(MenuList)
	//if err != nil {
	//
	//	result.SetError(errcode.LGERROR, "系统错误,"+err.Error(), "")
	//}
	//// 排序之后组装成 菜单json [{level1id:1,level1name:"222",level2:[{level2id:101,level2name:"3ee",level3:[{level3id:101001,level3name:"3fd5"}]}]}]
	//// 根据levelid，level2id，pid sort 字段排序  数据结构不对
	//// 多字段排序
	//result.SetSuccess(json)
	//return result
}

/**
查询菜单数据，根据角色查询菜单数据
*/
func QueryMenuMainBySql(isBack int, user sysmodel.SSOUser) string {

	//查询用户角色
	qroleSql := "select roleid from eb_uservsrole where UserID=?"
	rows, _ := dbhelper.Query(user.UserID, true, qroleSql, user.UserID)
	roleids := commutil.RowsToIdStrByCol(rows, "roleid")
	if g.IsEmpty(roleids) {
		roleids = "('1')" //默认为普通用户
	} else {
		roleids = "(" + roleids + ")"
	}

	//ifnull := dbhelper.GetIFNull()

	// 查询当前拥有 的菜单权限
	//u.tempname 暂时取消了 作为菜单名
	var query = "  select pg.windtype,  pmd.modelid as level1id,pmd.modelname as level1name,pmd.menuimg as level1img," +
		"pg.pid,pg.formtype as opentype, case when pg.FormType ='onlymenu' then pg.menuurl ELSE ent.menuurl end as menuurl," +
		"ent.editurl,md.modelid as level2id,md.modelname as level2name," +
		"ifnull(u.tempname , pg.pname ) as  pname,pg.menuimg,pg.sortid ,ent.verid " +
		"from sys_fpageid pg   join sys_fmodel md on md.modelid=pg.modelid  " +
		" join sys_fmodel pmd on pmd.modelid=md.parentid  " +
		" left join sys_fpagever ent on pg.pid=ent.pid and ent.entid=? " +
		" left join sys_uiform U on pg.pid=U.pid and U.entid=? " +
		"where pmd.isopen=1 and md.isopen=1 and pg.isopen=1 and pg.isopen is not null " +
		" and pmd.isbackconfig=? and md.isbackconfig=? " +
		"and exists( select 1 from sys_fpageaccess pac where pac.pid=pg.pid " +
		" and pac.entid=? and pac.roleid in " + roleids + ") " +
		"and pg.FormType !='InnerForm'   and PG.SubSysID in (select SubSysid from eb_entVsSystem where entid=?)" +
		"order by pmd.sortid,md.sortid,pg.sortid "
	// 取到之后保留 pid, uiusage max(levelid) 的记录

	valList, err := dbhelper.Query("admin", true, query, user.FormEntId, user.FormEntId, isBack, isBack, user.EntID, user.EntID)

	MenuList := make([]*model.MenuModel, 0)
	var lastlevel1id = "" // 上一个一级模块id
	var lastlevel2id = "" // 上一个二级模块id
	var level1index = 0   // 一级菜单当前下标
	var level2index = 0   // 二级菜单当前下标
	for _, value := range valList {
		// 取得了pid 和 uitempid 之后，然后从redis 获取表单信息
		pageInfo := value
		// 不为空
		pageInfo["menuurl"] = value["menuurl"]
		pageInfo["windtype"] = value["windtype"]

		if chkPid(commutil.ToInt(value["pid"])) {
			pageInfo["opentype"] = value["opentype"]
		} else {
			pageInfo["opentype"] = ""
		}

		pageInfo["pid"] = value["pid"]

		//替换固定基础档案菜单URL
		pageInfo["menuurl"] = basePidUrl(pageInfo["pid"], pageInfo["menuurl"])

		level1id := pageInfo["level1id"]
		level2id := pageInfo["level2id"]
		if (lastlevel1id == "" && lastlevel2id == "") || (lastlevel1id != level1id && lastlevel2id != level2id) {
			// 是第一个
			page := model.Page{
				Pid:      commutil.ToInt(pageInfo["pid"]),
				Ver:      commutil.ToInt(pageInfo["verid"]),
				ParentID: pageInfo["level2id"],
				MenuName: pageInfo["pname"],
				MenuUrl:  pageInfo["menuurl"],
				WindType: pageInfo["windtype"],
				OpenType: pageInfo["opentype"],
				MenuImg:  pageInfo["menuimg"],
				SortID:   commutil.ToFloat32(pageInfo["sortid"]),
			}
			children := make([]*model.Page, 0)
			children = append(children, &page)
			pagemodel := model.PageModel{ModelName: pageInfo["level2name"], ModelID: pageInfo["level2id"], ParentID: level1id, MenuImg: pageInfo["menuimg"], Child: children}
			level2List := make([]*model.PageModel, 0)
			level2List = append(level2List, &pagemodel)
			MenuList = append(MenuList, &model.MenuModel{ModelID: level1id, ModelName: pageInfo["level1name"], MenuImg: pageInfo["level1img"], Child: level2List})
			if lastlevel2id == "" && lastlevel1id == "" {
				// 第一个 下标都为0
				level1index = 0
			} else { // 否则下标增加
				level1index++
			}
			// 二级菜单下标需要归零
			level2index = 0
			lastlevel1id = level1id
			lastlevel2id = level2id
		} else {
			// 已经存在一级菜单
			// lastleve1id 是否和当前一致
			if lastlevel1id == level1id {
				// 在同一个一级模块下
				if lastlevel2id == level2id {
					// 是否在同一个二级模块下
					level2Menu := MenuList[level1index].Child[level2index]
					// 只需要添加三级菜单
					// 已经存在的三级菜单列表
					page := model.Page{
						Pid:      commutil.ToInt(pageInfo["pid"]),
						Ver:      commutil.ToInt(pageInfo["verid"]),
						ParentID: pageInfo["level2id"],
						MenuName: pageInfo["pname"],
						MenuUrl:  pageInfo["menuurl"],
						WindType: pageInfo["windtype"],
						OpenType: pageInfo["opentype"],
						MenuImg:  pageInfo["menuimg"],
						SortID:   commutil.ToFloat32(pageInfo["sortid"]),
					}
					children := level2Menu.Child
					children = append(children, &page)
					level2Menu.Child = children
				} else {
					// 不在同一个二级模块
					// 重新创建二级模块列表
					level2List := MenuList[level1index].Child

					page := model.Page{
						Pid:      commutil.ToInt(pageInfo["pid"]),
						Ver:      commutil.ToInt(pageInfo["verid"]),
						ParentID: pageInfo["level2id"],
						MenuName: pageInfo["pname"],
						MenuUrl:  pageInfo["menuurl"],
						WindType: pageInfo["windtype"],
						OpenType: pageInfo["opentype"],
						MenuImg:  pageInfo["menuimg"],
						SortID:   commutil.ToFloat32(pageInfo["sortid"]),
					}
					children := make([]*model.Page, 0)
					children = append(children, &page)

					pagemodel := model.PageModel{ModelName: pageInfo["level2name"], ModelID: pageInfo["level2id"], ParentID: level1id, MenuImg: pageInfo["menuimg"], Child: children}
					level2List = append(level2List, &pagemodel)
					MenuList[level1index].Child = level2List
					//MenuList = append(MenuList, &model.MenuModel{ModelID: level1id, ModelName: pageInfo["level1name"], MenuImg: pageInfo["level1img"], Child: level2List})
					level2index++
					lastlevel2id = level2id
				}
			} else {
				// 不在同一个一级模块下，新创建一级模块,二级模块，三级模块
				page := model.Page{
					Pid:      commutil.ToInt(pageInfo["pid"]),
					Ver:      commutil.ToInt(pageInfo["verid"]),
					ParentID: pageInfo["level2id"],
					MenuName: pageInfo["pname"],
					MenuUrl:  pageInfo["menuurl"],
					WindType: pageInfo["windtype"],
					OpenType: pageInfo["opentype"],
					MenuImg:  pageInfo["menuimg"],
					SortID:   commutil.ToFloat32(pageInfo["sortid"]),
				}
				children := make([]*model.Page, 0)
				children = append(children, &page)

				pagemodel := model.PageModel{ModelName: pageInfo["level2name"], ModelID: pageInfo["level2id"], ParentID: level1id, MenuImg: pageInfo["menuimg"], Child: children}
				level2List := make([]*model.PageModel, 0)
				level2List = append(level2List, &pagemodel)
				MenuList = append(MenuList, &model.MenuModel{ModelID: level1id, ModelName: pageInfo["level1name"], MenuImg: pageInfo["level1img"], Child: level2List})
				level1index++
				level2index = 0
				lastlevel1id = level1id
				lastlevel2id = level2id
			}
		}
	}
	json, err := jsonutil.ObjToJson(MenuList)
	if err != nil {
		panic(err)
	}

	return json
}

// 固定基础档案
func basePidUrl(pid string, menuUrl string) string {

	if pid == "20101" {
		//用户
		menuUrl = "user/user-edit/1"
	} else if pid == "20102" {
		//部门
		menuUrl = "dept/dept-edit/1"
	} else if pid == "20103" {
		//公司
		menuUrl = "company/company-edit/1"
	} else if pid == "20104" {
		//角色
		menuUrl = "role/role-edit/1"
	} else if pid == "20105" {
		//岗位
		menuUrl = "job/job-edit/1"
	}
	return menuUrl
}

// 固定菜单不需要生成opentype
func chkPid(pid int) bool {
	if pid == 20101 || pid == 20102 || pid == 20103 || pid == 20104 || pid == 20105 {
		return false
	} else {
		return true
	}

}

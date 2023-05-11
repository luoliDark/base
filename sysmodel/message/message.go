package message

/**
 * @Author zhengjie
 * @Description 消息接口和消息公用方法
 * @Date 11:57 2020/2/21
 **/
import (
	"strings"

	"github.com/luoliDark/base/db/conn"
	"github.com/luoliDark/base/sysmodel"
	"github.com/luoliDark/base/sysmodel/eb"
	"github.com/luoliDark/base/util/commutil"
)

// 消息实体类
type SendMessage interface {
	Receivers() []int
	Message() string
}

//--类型为role
//SELECT userid from eb_uservsrole where RoleID in (....)
//
//--类型为dept
//SELECT UserID from eb_user where deptid in (...)
//
//--类型为job
//select UserID from eb_uservsjob where jobid in (...)
//
//--类型为workgroup
//select UserID from eb_uservsworkgroup where workgroupid in (...)
func findUserIdByAccMap(accmap map[string][]int) []int {
	for key, value := range accmap {
		switch key {
		case "role":
			return findUserIdByRole(value)
		case "dept":
			return findUserIdByDept(value)
		case "job":
			return findUserIdByJob(value)
		}
	}

	return nil
}

func findUserIdByRole(accObjIds []int) []int {
	if len(accObjIds) == 0 {
		return nil
	}
	engine, _ := conn.GetDB()
	userRoles := make([]sysmodel.Eb_uservsrole, 0)
	engine.Where("RoleID in (?)", arrayToString(accObjIds)).Find(userRoles)

	if len(userRoles) == 0 {
		return nil
	}

	userIdList := make([]int, 0)
	for i := range userRoles {
		userIdList = append(userIdList, userRoles[i].UserID)
	}
	return userIdList
}

func findUserIdByDept(accObjIds []int) []int {
	if len(accObjIds) == 0 {
		return nil
	}
	engine, _ := conn.GetDB()
	users := make([]eb.Eb_user, 0)
	engine.Where("deptid in (?)", arrayToString(accObjIds)).Find(users)

	if len(users) == 0 {
		return nil
	}

	userIdList := make([]int, 0)
	for i := range users {
		userIdList = append(userIdList, users[i].UserID)
	}
	return userIdList
}

func findUserIdByJob(accObjIds []int) []int {
	if len(accObjIds) == 0 {
		return nil
	}
	engine, _ := conn.GetDB()
	users := make([]sysmodel.Eb_uservsjob, 0)
	engine.Where("JobID in (?)", arrayToString(accObjIds)).Find(users)

	if len(users) == 0 {
		return nil
	}

	userIdList := make([]int, 0)
	for i := range users {
		userIdList = append(userIdList, commutil.ToInt(users[i].UserID))
	}
	return userIdList
}

func arrayToString(array []int) string {
	if len(array) == 0 {
		return ""
	}
	var arrayString string
	for i := range array {
		arrayString = commutil.AppendStr(arrayString, commutil.ToString(array[i]), ",")
	}
	arrayString = strings.TrimRight(arrayString, ",")
	return arrayString
}

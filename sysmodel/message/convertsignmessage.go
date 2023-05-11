package message

/**
 * @Author zhengjie
 * @Description 转批消息
 * @Date 11:59 2020/2/21
 **/
import (
	"github.com/luoliDark/base/db/conn"
	"github.com/luoliDark/base/sysmodel/wf"
	"github.com/luoliDark/base/util/commutil"
)

// 消息实体类
type ConvertSignMessage struct {
	BillId int
	StepId int
}

func (m ConvertSignMessage) Message() string {
	msg := commutil.AppendStr("单号:", commutil.ToString(m.BillId), ",您有新的转批审批待您审批")
	return msg
}

//查询Sys_wfstepaccessdynamic表中转批用户
func (m ConvertSignMessage) Receivers() []int {
	engine, _ := conn.GetDB()

	users := make([]wf.Sys_wfstepaccessdynamic, 0)
	engine.Where("IsConvert = 1", m.StepId).Find(&users)

	if len(users) == 0 {
		return nil
	}

	userIdList := make(map[string][]int)
	for i := range users {
		tempList := userIdList[users[i].AccType]
		if tempList == nil {
			tempList = make([]int, 0)
		}
		tempList = append(tempList, users[i].AccObjID)
		userIdList[users[i].AccType] = tempList
	}
	return findUserIdByAccMap(userIdList)
}

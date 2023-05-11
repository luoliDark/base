package message

/**
 * @Author zhengjie
 * @Description 抄送消息
 * @Date 11:57 2020/2/21
 **/
import (
	"github.com/luoliDark/base/db/conn"
	"github.com/luoliDark/base/sysmodel"
	"github.com/luoliDark/base/util/commutil"
)

// 消息实体类
type CCMessage struct {
	BillId int
	StepId int
}

func (m CCMessage) Message() string {
	msg := commutil.AppendStr("单号", commutil.ToString(m.BillId), ",有新的审批记录抄送给您")
	return msg
}

//查询Sys_WFCCUser表中抄送用户
func (m CCMessage) Receivers() []int {
	engine, _ := conn.GetDB()

	users := make([]sysmodel.Sys_wfccuser, 0)
	engine.Where("StepID = ?", m.StepId).Find(&users)

	if len(users) == 0 {
		return nil
	}

	userIdList := make([]int, 0)
	for i := range users {
		userIdList = append(userIdList, users[i].CCUserID)
	}
	return userIdList
}

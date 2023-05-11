package message

/**
 * @Author zhengjie
 * @Description  关注消息
 * @Date 12:17 2020/2/21
 **/
import (
	"github.com/luoliDark/base/db/conn"
	"github.com/luoliDark/base/sysmodel"
	"github.com/luoliDark/base/util/commutil"
)

// 消息实体类
type WaiteLookMessage struct {
	Pid    int
	BillId int
}

func (m WaiteLookMessage) Message() string {
	msg := commutil.AppendStr("您关注的单号", commutil.ToString(m.BillId), "有新的审批记录")
	return msg
}

//查询Sys_WFWaiteLook
func (m WaiteLookMessage) Receivers() []int {
	engine, _ := conn.GetDB()

	users := make([]sysmodel.Sys_wfwaitelook, 0)
	engine.Where("Pid = ? and BillID = ?", m.Pid, m.BillId).Find(&users)

	if len(users) == 0 {
		return nil
	}

	userIdList := make([]int, 0)
	for i := range users {
		userIdList = append(userIdList, users[i].LookUid)
	}
	return userIdList
}

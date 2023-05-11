package sendmsg

import (
	"paas/base/db/conn"
	"paas/base/sysmodel"
	"paas/base/sysmodel/eb"
	"paas/base/util/commutil"
	"testing"
)

func TestSendMain(t *testing.T) {

	msg := sysmodel.MsgEntity{}
	//查询当前节点
	eng, _ := conn.GetDB()

	msg.Pid = 50201
	msg.PrimaryKey = "0dd8a1d5b40144039223ea692c5b22f5"
	msg.EntId = 1
	msg.BillNo = "EN20220809007"
	msg.TemplateId = 1 //  1  是通用型模板
	msg.LinkUrl = GetMobileMsgUrl(50201, 0, "日常报销", "0dd8a1d5b40144039223ea692c5b22f5")

	msg.Body += "测试通知下"

	//-------------------------------------------------发送抄送信息--------------------------------------------
	msg.Title = "测试通知 给您！" + msg.BillNo + commutil.GetNowTime()

	users := make([]eb.Eb_user, 0)
	err := eng.Where("usercode=28").Find(&users)
	if err != nil {
		return
	}

	msg.ToUsers = users
	SendMain(msg, &sysmodel.SSOUser{UserID: "8"})
}

package sendmsg

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/luoliDark/base/db/conn"
	"github.com/luoliDark/base/db/dbhelper"
	"github.com/luoliDark/base/loghelper"
	"github.com/luoliDark/base/redishelper/rediscache"
	"github.com/luoliDark/base/sysmodel"
	"github.com/luoliDark/base/sysmodel/bbs"
	"github.com/luoliDark/base/util/commutil"
	"github.com/luoliDark/base/util/httputil"
)

//发送消息入口方法
func SendMain(msg sysmodel.MsgEntity, user *sysmodel.SSOUser) sysmodel.ResultBean {

	//插入DB
	if msg.MsgType == "cc" || msg.MsgType == "push" || msg.MsgType == "info" {
		_ = InsertMsgToDB(msg, user)
	}

	//获取该entid对应的msgserver
	entM := rediscache.GetHashMap("eb_enterprise_" + commutil.ToString(msg.EntId))
	msgUrl := entM["msgserverip"]
	if g.IsEmpty(msgUrl) {
		loghelper.ByError("发送消息失败", "没有维护msgUrl字段到eb_enterprise, msg=", user.UserID)
		return sysmodel.ResultBean{IsSuccess: false, ErrorMsg: "没有维护msgUrl字段到eb_enterprise"}
	}

	msgUrl = strings.TrimRight(msgUrl, "/")
	msgUrl += "/main/beginSend"
	result, err := httputil.PostJson(msgUrl, msg)

	msgtmp, _ := json.Marshal(msg)
	ms := string(msgtmp)
	if err != nil {
		loghelper.ByError("发送消息失败,请检查", fmt.Sprint(result, ",err:", err.Error(), ",msg:"+ms), user.UserID)
		return sysmodel.ResultBean{IsSuccess: false, ErrorMsg: err.Error()}
	} else {
		resultBean := sysmodel.JsonToBean(result)
		if !resultBean.IsSuccess {
			loghelper.ByError("发送消息失败,请检查", fmt.Sprint("result:", result, ",msg:"+ms), user.UserID)
			return sysmodel.ResultBean{IsSuccess: false, ErrorMsg: result}
		}
		//else {
		//	loghelper.ByInfo("发送消息成功", result + ms, user.UserID)
		//}
	}

	return sysmodel.ResultBean{IsSuccess: true}
}

//插入消息到DB记录
func InsertMsgToDB(entity sysmodel.MsgEntity, user *sysmodel.SSOUser) error {
	var bbsmsgbypid bbs.Bbsmsgbypid
	bbsmsgbypid.Pid = entity.Pid
	bbsmsgbypid.MsgId = commutil.GetUUID()
	bbsmsgbypid.MsgType = entity.MsgType
	bbsmsgbypid.BillNo = entity.BillNo
	bbsmsgbypid.BillId = entity.PrimaryKey
	bbsmsgbypid.Title = entity.Title
	bbsmsgbypid.Body = entity.Body
	bbsmsgbypid.EntId = commutil.ToString(entity.EntId)
	bbsmsgbypid.FormEntId = user.FormEntId
	bbsmsgbypid.IsReply = 0

	if g.IsEmpty(bbsmsgbypid.MsgType) {
		bbsmsgbypid.MsgType = "info"
	}

	bbsmsgbypid.SendDate = commutil.GetNowTime()
	bbsmsgbypid.SendUid = user.UserID

	if bbsmsgbypid.BillVer == 0 {
		m := rediscache.GetHashMap("sys_fpagever_" + commutil.ToString(bbsmsgbypid.FormEntId) + "_" + commutil.ToString(bbsmsgbypid.Pid))
		if len(m) > 0 {
			bbsmsgbypid.BillVer = commutil.ToInt(m["verid"])
			bbsmsgbypid.PName = m["pname"]
		} else {
			panic("redis未找到该单据信息")
		}
	}

	var formUid string
	//if g.IsEmpty(bbsmsgbypid.BillNo) {
	m := rediscache.GetHashMap("sys_fpage_" + commutil.ToString(bbsmsgbypid.Pid))
	if len(m) > 0 {

		tb := m["sqltablename"]

		sql := "select billno,totalmoney,userid,deptid,detailcostallname from " + tb + " where billid=?"

		re, _ := dbhelper.Query("", true, sql, bbsmsgbypid.BillId)
		if len(re) > 0 {
			//补齐流水号
			bbsmsgbypid.BillNo = re[0]["billno"]
			bbsmsgbypid.TotalMoney = re[0]["totalmoney"]
			formUid = re[0]["userid"]
			bbsmsgbypid.DetailCostAllName = re[0]["detailcostallname"]
		}

		bbsmsgbypid.PName = m["pname"]
	} else {
		panic("redis未找到该单据信息")
	}
	//}

	//发送人信息
	if g.IsEmpty(bbsmsgbypid.UserName) {
		sql := "select username,deptid from eb_user where userid=? and entid=?"
		row, _ := dbhelper.Query("", true, sql, formUid, bbsmsgbypid.EntId)
		if len(row) > 0 {
			bbsmsgbypid.UserName = row[0]["username"]
			bbsmsgbypid.SendDeptId = row[0]["deptid"]

			sql = "select deptname from eb_dept where deptid=? and entid=?"
			dname, _ := dbhelper.QueryFirstCol("", true, sql, bbsmsgbypid.SendDeptId, bbsmsgbypid.EntId)
			bbsmsgbypid.BDeptname = dname
		}
	}
	engine, _ := conn.GetDB()

	//插入消息表
	_, err := engine.InsertOne(bbsmsgbypid)
	if err != nil {
		return err
	}

	//插入消息接收人表

	var recivers = entity.ToUsers
	bbsbypidrecivers := make([]bbs.Bbsbypidreciver, 0)

	if !g.IsEmpty(recivers) {
		for _, reciver_user := range recivers {
			var bbsbypidreciver bbs.Bbsbypidreciver
			bbsbypidreciver.IsViewed = 0
			bbsbypidreciver.MsgId = bbsmsgbypid.MsgId
			bbsbypidreciver.ReciverUid = reciver_user.UserID
			bbsbypidreciver.ReciverTime = commutil.GetNowTime()
			bbsbypidreciver.MsgType = entity.MsgType
			bbsbypidreciver.ViewTime = commutil.GetNowTime()
			bbsbypidreciver.EntId = commutil.ToInt(entity.EntId)
			bbsbypidrecivers = append(bbsbypidrecivers, bbsbypidreciver)
		}
		if len(bbsbypidrecivers) > 0 {
			_, error := engine.Insert(bbsbypidrecivers)
			if error != nil {
				return error
			}

		}
	}

	return nil
}

var h5url = ""

func init() {
	initH5Url()
}

func initH5Url() {
	if g.IsEmpty(h5url) {
		m := rediscache.GetHashMap("sys_global_1")
		if len(m) > 0 {
			h5url = m["h5url"]
			fmt.Println("h5url", h5url)
		}
	}
}

func GetMobileMsgUrl(pid, ver int, pname, primarykey string) string {
	var url string
	if h5url == "" {
		initH5Url()
	}
	url = h5url + "?pid=" + commutil.ToString(pid) + "&verid=0" + "&isreadonly=1&primarykey=" + primarykey + "&pk=" + primarykey + "&title=" + pname

	return url
}

//Form/editform/50201/548/copy_50201D5OE1BML22BDI591
func GetPCMsgUrl(pid, ver int, pname, primarykey string) string {
	var url string
	url = fmt.Sprintf(`Form/editform/%v/%v/%v?pname=%v`,
		pid, ver, primarykey, pname)
	return url
}

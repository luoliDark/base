package loghelper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/luoliDark/base/confighelper"
	"github.com/rs/zerolog"
)

type LogBaseInfo struct {
	Id        int64
	Title     string `xorm:"Title" json:"title"`
	LogMsg    string `xorm:"LogMsg" json:"logMsg"`
	Uid       string `xorm:"Uid" json:"uid"`
	IP        string `xorm:"Ip" json:"ip"`
	FileName  string `xorm:"Filename" json:"fileName"`
	FileLine  int    `xorm:"fileline" json:"fileline"`
	Extend1   string `xorm:"Extend1" json:"extend1"`
	Extend2   string `xorm:"Extend2" json:"extend2"`
	CreatedAt int64  `xorm:"CreatedAt" json:"createdAt"`
	LogTime   string `xorm:"logtime" json:"logtime"`
	LogType   string `xorm:"LogType" json:"logType"`
}

// 程序开始执行前调用本方法，会返回开始时间，程序执行完后再调用 execEnd ,execEnd方法会自动判定是否超时
func BeginimeRecord() time.Time {
	return time.Now()
}

//记录结束时间
//userID 用户ID
//countType 统计类型 例：SQL执行，rest请求，事件执行
//beginExecTime 开始执行的时间
// execScript 执行的SQL或其它语句
func EndTimeRecord(userID string, countType string, beginExecTime time.Time, execScript ...interface{}) {

	//todo 时间比较如果超过5秒 需要收到到日志服务器

	diff := time.Now().Sub(beginExecTime)
	if isdev == "1" {

		if len(execScript) > 0 {
			fmt.Println("执行时长：", diff.Seconds(), " 秒", execScript[0])
		} else {
			fmt.Println("执行时长：", diff.Seconds(), " 秒", countType)
		}

	}
	if diff.Seconds() > 5 {
		//超过5秒
		byteSlice := make([]byte, 20)
		byteBuffer := bytes.NewBuffer(byteSlice)
		byteBuffer.WriteString(fmt.Sprint("超时了：", diff.Seconds(), " s(秒) "))
		if len(execScript) > 0 {
			for index := range execScript {
				b, err := json.Marshal(execScript[index])
				if err == nil {
					byteBuffer.Write(b)
				}
				byteBuffer.Write([]byte("---"))
			}
		}

		msg := byteBuffer.String()

		if isdev == "1" {
			fmt.Println("超时了：", diff.Seconds(), " s(秒) ", msg)
		}
		_, file, line, _ := runtime.Caller(1)
		SendLog(countType, msg, "", userID, "logTimeout", file, line)
	}
}

// ByInfo 普通日志记录
func ByInfo(title string, logMsg string, uid string) {
	_, file, line, _ := runtime.Caller(1)
	ip := GetIP()

	if isdev == "1" {
		fmt.Println("普通日志：", title, logMsg, file, line, ip)
	}

	SendLog(title, logMsg, ip, uid, "logInfo", file, line)
}

// ByDataCenterInfo 数据平台日志
func ByDataCenterInfo(title string, logMsg string, uid string) {
	_, file, line, _ := runtime.Caller(1)
	ip := GetIP()

	if isdev == "1" {
		fmt.Println("数据平台日志：", title, logMsg, file, line, ip)
	}

	SendLog(title, logMsg, ip, uid, "ByDataCenter", file, line)
}

// ByError用于记录所有错误日志（业务错误）
func ByError(title string, logMsg string, uid string) {
	_, file, line, _ := runtime.Caller(1)
	ip := GetIP()
	logMsg = logMsg + ";Stack=" + string(debug.Stack())
	if isdev == "1" {
		fmt.Println("一般错误："+title, logMsg)
		fmt.Println(file, line)
	}

	SendLog(title, logMsg, ip, uid, "logError", file, line)
}

// ByHighError用于记录所有严重错误日志（例：系统异常）
func ByHighError(title string, logMsg string, uid string) {
	_, file, line, _ := runtime.Caller(1)
	ip := GetIP()
	logMsg = logMsg + ";Stack=" + string(debug.Stack())
	if isdev == "1" {
		fmt.Println("严重错误："+title, logMsg, file, line, ip)
	}

	SendLog(title, logMsg, ip, uid, "logHighError", file, line)
}

// ByConfig  所有后台配置操作一律要记录日志
func ByConfig(title string, logMsg string, uid string) {
	_, file, line, _ := runtime.Caller(1)
	ip := GetIP()

	if isdev == "1" {
		fmt.Println(" 配置被修改了："+title, logMsg, file, line, ip)
	}

	SendLog(title, logMsg, ip, uid, "logConfig", file, line)
}

// ByTimeOut  记录所有超时的日志 例：ajax请求，事件执行效率，保存超时时间，提交超时时间 以超过5s做为记录依据
func ByTimeOut(title string, logMsg string, uid string, sqlScriptOrFnName string) {
	_, file, line, _ := runtime.Caller(1)
	ip := GetIP()

	if isdev == "1" {
		fmt.Println(" 运行超时了："+title, logMsg, "超时代码为：", sqlScriptOrFnName, file, line, ip)
	}
	SendLog(title, logMsg, ip, uid, "logTimeout", file, line)
}

// ByConfig  所有后台配置操作一律要记录日志
func ByRestCatchErr(title string, logMsg string, uid string) {
	_, file, line, _ := runtime.Caller(1)
	ip := GetIP()

	if isdev == "1" {
		fmt.Printf("%c[1;31;40m Rest请求错误 UserID="+uid+"时间="+time.Now().Format("2006-01-02 12:02")+" 错误代码："+logMsg+" %c[0m", 0x1B, 0x1B) //红色打印
	}

	SendLog(title, logMsg, ip, uid, "logHighError", file, line)
}

var logModel *LogModel
var isdev string

func init() {
	path := confighelper.GetIniConfig("logserver", "logPath")
	m, err := NewLogModel(path)
	if err != nil {
		fmt.Println(err)
	}

	logModel = m

	//是否开发模式
	isdev = confighelper.GetIniConfig("global", "isdeving")
}

func SendLog(title, logMsg, ip, uid, logType, filePath string, line int) error {
	log := LogBaseInfo{
		Title:     title,
		Uid:       uid,
		LogMsg:    logMsg,
		IP:        ip,
		FileName:  filePath,
		FileLine:  line,
		CreatedAt: time.Now().Unix(),
		LogTime:   time.Now().Format("2006-01-02 15:04:05"),
		LogType:   logType,
	}

	if logModel == nil {
		return errors.New("log path error")
	}

	if logType == "logError" || logType == "logHighError" || logType == "logTimeout" {
		logModel.Err(log)
	} else {
		logModel.Info(log)
	}

	return nil
}

type LogModel struct {
	timeStr string
	errLog  *zerolog.Logger
	infoLog *zerolog.Logger
	path    string
}

func (l *LogModel) Err(i interface{}) {
	l.checkTime()
	l.errLog.Err(errors.New("err")).Interface("logModel", i).Send()
}

func (l *LogModel) Info(i interface{}) {
	l.checkTime()
	l.infoLog.Info().Interface("logModel", i).Send()
}

func (l *LogModel) checkTime() {
	now := time.Now()
	timeStr := now.Format("2006-01-02")
	fmt.Println("timeStr:" + now.Format("2006-01-02 15:04:05"))
	if l.timeStr != timeStr {
		fmt.Println("change time:" + timeStr)
		errLogFile, _ := os.OpenFile(l.path+"/err_"+timeStr+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
		errLog := zerolog.New(errLogFile)

		infoLogFile, _ := os.OpenFile(l.path+"/info_"+timeStr+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)

		infoLog := zerolog.New(infoLogFile)

		l.timeStr = timeStr
		l.errLog = &errLog
		l.infoLog = &infoLog
	}
}

func NewLogModel(path string) (*LogModel, error) {
	l := &LogModel{}
	l.timeStr = time.Now().Format("2006-01-02")

	errLogFile, err := os.OpenFile(path+"/err_"+l.timeStr+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		return nil, err
	}
	errLog := zerolog.New(errLogFile)

	infoLogFile, err := os.OpenFile(path+"/info_"+l.timeStr+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		return nil, err
	}
	infoLog := zerolog.New(infoLogFile)

	l.errLog = &errLog
	l.infoLog = &infoLog
	l.path = path

	return l, nil
}

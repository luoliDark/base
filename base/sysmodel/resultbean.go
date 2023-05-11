//jsz by 2020.2.8 用于sevices层返回到前台的统一用的resultBean对象

package sysmodel

import (
	"base/base/util/commutil"
	"encoding/json"
)

type ResultBean struct {
	IsSuccess   bool        `xorm:"IsSuccess" json:"IsSuccess"`     //是否成功 系统错误或业务错误都返回false
	ErrorCode   string      `xorm:"ErrorCode" json:"ErrorCode"`     //错误代码
	ErrorMsg    string      `xorm:"ErrorMsg" json:"ErrorMsg"`       //报错详细信息 注：只针对于系统运行错误，业务代码错误一律只返回编码后再用多语言做处理
	ResultData  interface{} `xorm:"ResultData" json:"ResultData"`   //返回的json数据字符串
	ResultTotal int         `xorm:"ResultTotal" json:"ResultTotal"` //查询的总数
	PrimaryKey  string
}

// 处理成功
func (this *ResultBean) SetSuccess(resultData interface{}) *ResultBean {
	this.IsSuccess = true
	this.ResultData = resultData
	return this
}

// 处理成功
func (this *ResultBean) SetTotal(total int) {
	this.ResultTotal = total
}

// 处理失败
func (this *ResultBean) SetError(errCode string, errMsg string, resultData interface{}) *ResultBean {
	this.IsSuccess = false
	this.ErrorCode = errCode
	this.ErrorMsg = errMsg
	this.ResultData = resultData
	return this
}

// 返回内存地址的值
func (this *ResultBean) Get() ResultBean {
	return *this
}

// 将本对象转为json
func (this *ResultBean) ToJson() string {
	str, _ := json.Marshal(this)
	return commutil.ToString(str)
}

// json 反序列化为bean
func JsonToBean(beanStr string) (bean ResultBean) {
	bean = ResultBean{IsSuccess: false, ErrorMsg: "转换前"}
	err := json.Unmarshal([]byte(beanStr), &bean)
	if err != nil {
		bean.IsSuccess = false
		bean.ErrorMsg = err.Error()
	}
	return bean
}

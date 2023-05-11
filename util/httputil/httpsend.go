package httputil

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/luoliDark/base/loghelper"
)

const (
	CONN_TIME_OUT = time.Second * 10
)

// MIME
const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
)

/**
Get
*/
func Get(apiURL string, params url.Values) (resData string, err error) {
	parse, err := url.Parse(apiURL)
	if err != nil {
		return "", err
	}
	parse.RawQuery = params.Encode()
	urlPath := parse.String()
	resp, err := http.Get(urlPath)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

/**

 */
func HttpHeaderReq(apiURL string, params url.Values, method string, headParams map[string]string, body string, contentType string) (resData string, err error) {
	client := &http.Client{}
	apiURL = strings.TrimLeft(strings.TrimRight(apiURL, " "), " ")
	var request *http.Request
	if contentType == "form-data" {
		///设置请求
		requestout, err := http.NewRequest(method, apiURL, strings.NewReader(params.Encode()))
		request = requestout
		if err != nil {
			panic(err)
		}
	} else {
		///设置请求
		requestout, err := http.NewRequest(method, apiURL, strings.NewReader(body))
		request = requestout
		if err != nil {
			panic(err)
		}
	}
	///判断请求头参数是否为空
	if nil != headParams {
		for k, v := range headParams {
			request.Header.Add(k, v)
		}
	}
	resp, err := client.Do(request)
	if err != nil {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
		return "", err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

/**
POST
*/
func Post(apiURL string, params url.Values) (resData string, err error) {
	apiURL = strings.TrimLeft(strings.TrimRight(apiURL, " "), " ")
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func PostWithTimeout(apiURL string, params url.Values) (resData string, err error) {
	apiURL = strings.TrimLeft(strings.TrimRight(apiURL, " "), " ")
	client := &http.Client{
		Timeout: 600 * time.Second,
	}
	resp, err := client.PostForm(apiURL, params)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

/**
POST
*/
// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func PostJson(url string, data interface{}) (string, error) {
	url = strings.TrimLeft(strings.TrimRight(url, " "), " ")
	// 超时时间：5秒
	jsonStr, _ := json.Marshal(data)
	loghelper.ByInfo("发送参数", url+",请求参数:"+string(jsonStr), "")
	client := &http.Client{
		Timeout: 600 * time.Second,
	}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	//resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}

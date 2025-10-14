package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/luoliDark/base/confighelper"
	"github.com/luoliDark/base/loghelper"
	"github.com/luoliDark/base/util/commutil"

	"github.com/gogf/gf/os/gfile"
)

const (
	CONN_TIME_OUT = time.Second * 10
	ostype        = runtime.GOOS
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

/*
*
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
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func GetByUrl(url string) string {
	res, err := http.Get(url)
	if err != nil {
		return ""
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return ""
	}
	return string(robots)
}

/*
*
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

// 下载文件
func HttpGetFile(httpUrl, ex, savePath string) (string, error) {
	resp, err := http.Get(httpUrl)
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", fmt.Errorf("resp is nil")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("resp status is %s", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	uuid := commutil.GetUUID()
	year := commutil.ToString(time.Now().Year())
	month := commutil.ToString(time.Now().Month().String())
	day := commutil.ToString(time.Now().Day())

	dicPath := path.Join(savePath, year, month, day)
	filePath := commutil.AppendStr(strings.ReplaceAll(confighelper.LoadGoEnv(), "\\", "/"), dicPath)
	filePath = formatFilePath(filePath)

	err = gfile.Mkdir(filePath)
	if err != nil {
		return "", err
	}
	fileName := uuid + ex

	fmt.Println("保存的路径为：", filePath+fileName)
	f, _ := os.OpenFile(filePath+fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	_, err = f.Write(body)
	if err != nil {
		return "", err
	}

	return dicPath + fileName, nil
}

func formatFilePath(pathStr string) string {
	if ostype == "windows" {
		pathStr = strings.ReplaceAll(pathStr, "\\/", "\\")
		pathStr = strings.ReplaceAll(pathStr, "/", "\\")
	}
	pathStr = strings.ReplaceAll(pathStr, "\\", "/")
	dir := path.Dir(pathStr)
	_, errDir := os.Stat(dir)
	if errDir != nil {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err.Error())
		}
	}
	return pathStr
}

/*
*
发送请求,获取response对象，
返回的是文件对象
*/
func HttpHeaderReqGetResp(apiURL string, params url.Values, method string,
	headParams map[string]string, body string, contentType string) (response *http.Response, err error) {
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

	return resp, err
}

/*
*
发送文件上传请求
*/
func HttpHeaderReqMulti(apiURL string, params map[string]string, method string, headParams map[string]string,
	body string) (resData string, err error) {
	client := &http.Client{}
	apiURL = strings.TrimLeft(strings.TrimRight(apiURL, " "), " ")
	var request *http.Request

	payload := new(bytes.Buffer)
	w := multipart.NewWriter(payload)
	defer w.Close() //结束后，必须释放
	for k, v := range params {
		if k == "file" {
			file, errFile1 := os.Open(v)
			part1, errFile1 := w.CreateFormFile("file", filepath.Base(v))
			_, errFile1 = io.Copy(part1, file)
			if errFile1 != nil {
				panic("文档读取错误！" + errFile1.Error())
			}
		} else {
			w.WriteField(k, v)
		}
	}
	error := w.Close()
	if error != nil {
		fmt.Println(error)
		return
	}
	///设置请求
	requestout, err := http.NewRequest(method, apiURL, payload)
	request = requestout
	if err != nil {
		panic(err)
	}
	///判断请求头参数是否为空
	if nil != headParams {
		for k, v := range headParams {
			request.Header.Add(k, v)
		}
	}
	request.Header.Set("Content-Type", w.FormDataContentType())

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

/*
*
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

/*
*
发送post请求
新增head参数
*/
func PostHeaderContentType(apiURL string, params string, contentType string, headmap map[string]string) (resData string, err error) {
	apiURL = strings.TrimLeft(strings.TrimRight(apiURL, " "), " ")
	var client = &http.Client{}
	request, err := http.NewRequest("post", apiURL, strings.NewReader(params))
	//添加header选项
	for key, value := range headmap {
		request.Header.Add(key, value)
	}
	request.Header.Set("Content-Type", contentType)
	jsonStr2, _ := json.Marshal(headmap)
	loghelper.ByInfo("发送参数", apiURL+",请求参数:"+string(params)+",请求头："+string(jsonStr2)+",Content-Type:"+contentType, "")

	resp, err := client.Do(request)
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

/*
*
发送post请求
新增head参数
*/
func PostHeader(apiURL string, params string, headmap map[string]string) (resData string, err error) {
	apiURL = strings.TrimLeft(strings.TrimRight(apiURL, " "), " ")
	var client = &http.Client{}
	request, err := http.NewRequest("POST", apiURL, strings.NewReader(params))
	//添加header选项
	for key, value := range headmap {
		request.Header.Add(key, value)
	}
	jsonStr2, _ := json.Marshal(headmap)
	loghelper.ByInfo("发送参数", apiURL+",请求参数:"+string(params)+",请求头："+string(jsonStr2), "")

	resp, err := client.Do(request)
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

/*
*
发送post请求
新增head参数
*/
func HttpBytesHeader(httpMethod string, apiURL string, params []byte, headmap map[string]string) (resData string, err error) {
	apiURL = strings.TrimLeft(strings.TrimRight(apiURL, " "), " ")
	var client = &http.Client{}
	body := bytes.NewReader(params)
	request, err := http.NewRequest(httpMethod, apiURL, body)
	//添加header选项
	for key, value := range headmap {
		request.Header.Add(key, value)
	}
	jsonStr2, _ := json.Marshal(headmap)
	loghelper.ByInfo("发送参数", apiURL+",请求参数:"+string(params)+",请求头："+string(jsonStr2), "")

	resp, err := client.Do(request)
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
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}

/**
POST
*/
// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func PostJsonStr(url string, str string) (string, error) {
	url = strings.TrimLeft(strings.TrimRight(url, " "), " ")
	loghelper.ByInfo("发送参数", url+"发送参数为："+string(str), "")
	// 超时时间：5秒
	resp, err := http.Post(url, "application/json", bytes.NewBufferString(str))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}

/**
POST
*/
// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func PostMapJson(url string, data map[string]interface{}) (string, error) {
	url = strings.TrimLeft(strings.TrimRight(url, " "), " ")
	// 超时时间：5秒

	jsonStr, _ := json.Marshal(data)
	fmt.Println(string(jsonStr))
	loghelper.ByInfo("发送参数", url+"发送参数为："+string(jsonStr), "")

	client := &http.Client{
		Timeout: 600 * time.Second,
	}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer([]byte(jsonStr)))
	//resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}

/**
POST
*/
// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func PostJsonByHeader(url string, data string, headmap map[string]string) (string, error) {
	url = strings.TrimLeft(strings.TrimRight(url, " "), " ")
	// 超时时间：5秒
	client := &http.Client{
		Timeout: 600 * time.Second,
	}
	jsonStr, _ := json.Marshal(data)
	jsonStr2, _ := json.Marshal(headmap)
	reqID := commutil.GetUUID()
	loghelper.ByInfo("发送参数"+reqID, url+"发送参数为："+string(jsonStr)+"，请求头："+string(jsonStr2), "")

	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	//添加header选项
	for key, value := range headmap {
		request.Header.Add(key, value)
	}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}

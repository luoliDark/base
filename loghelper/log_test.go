package loghelper

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestByError(t *testing.T) {
	ByError(",", "", "")
}

func TestEndTimeRecord(t *testing.T) {
	var (
		str  string
		list []string
		m    map[string]string
	)

	str = "aaa"

	list = append(list, "aaa")
	list = append(list, "bbb")
	list = append(list, "ccc")
	list = append(list, str)

	m = make(map[string]string)
	m["aaa"] = "aaa"
	m["bbb"] = "bbb"
	m["ccc"] = "ccc"

	b, err := json.Marshal(m)

	fmt.Println(string(b))
	fmt.Println(err)

	EndTimeRecord("userId", "type", time.Now().Add(time.Second*50), str, list, m)
}

func TestSendLog(t *testing.T) {
	//SendLog("123", "msg", "127", "3", "logError")
	//SendLog("123", "msg", "127", "3", "logConfig")
}

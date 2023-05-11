package log

import "testing"

func TestSaveLog(t *testing.T) {

	l := LogByTimeOut{
		LogBaseInfo{
			Title:    "asdf",
			LogMsg:   "logmsg",
			Uid:      "asdf",
			IP:       "ip",
			FileName: "filename",
			Extend1:  "ex1",
			Extend2:  "ex2",
		},
	}

	l.SaveLog()
}

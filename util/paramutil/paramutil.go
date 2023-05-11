package paramutil

import (
	"base/db/conn"
	"base/redishelper/rediscache"
	"time"
)

var ROW, COL string
var ISNULL = "isnull"
var DBO = "dbo."
var NOLOCK = "(NOLOCK)"
var BeginDate time.Time
var LangCode = ""
var Haaaa = ""

func IPrintImpExpImpl() {
	if conn.DBType == "mysql" {
		ISNULL = "IFNULL"
		DBO = ""
		NOLOCK = ""
	} else if conn.DBType == "oracle" {
		ISNULL = "NVL"
		DBO = ""
		NOLOCK = ""
	}
	BeginDate = time.Now()
	ROW = rediscache.GetLanguageText("label", LangCode, "row") + ":"
	COL = rediscache.GetLanguageText("label", LangCode, "column") + ":"

}

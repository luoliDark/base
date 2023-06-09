package paramutil

import (
	"time"

	"github.com/luoliDark/base/db/conn"
	"github.com/luoliDark/base/redishelper/rediscache"
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

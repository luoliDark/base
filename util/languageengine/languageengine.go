package languageengine

import (
	"paas/base/db/dbhelper"
	"paas/base/util/commutil"
	"strings"
)

var I18NTABEL_SUFFIX = "_i18n"

func I18nTableIsExists(tableName string, columnName string) bool {
	var tableIsExist = false
	if commutil.IsNullOrEmpty(columnName) {
		columnName = "*"
	}

	_, error := dbhelper.Query("0", false, "select "+columnName+" from "+(tableName+I18NTABEL_SUFFIX)+" where 1=1 limit 1")
	tableIsExist = true
	if error != nil {
		tableIsExist = false
	}
	return tableIsExist
}

func HandleLanguageCode(languageCode string) string {
	if commutil.IsNullOrEmpty(languageCode) || languageCode == "\"null\"" || languageCode == "undefined" || languageCode == "null" {
		languageCode = ""
	}
	return languageCode
}

func GetDefaultLanguage() (result string) {
	data, _ := dbhelper.Query("0", false, "select langcode from sys_lang where isdf = 1")
	if data != nil && len(data) > 0 {
		result = strings.Trim(data[0]["langcode"], "")
	}
	return result
}

func FormatI18nSql(languageCode string, sql string) string {
	var dfLanguageCode = GetDefaultLanguage()

	if !commutil.IsNullOrEmpty(languageCode) && strings.ReplaceAll(languageCode, " ", "") != strings.ReplaceAll(dfLanguageCode, " ", "") &&
		strings.Index(sql, "#languagecode#") > -1 {
		sql = strings.ReplaceAll(sql, "#languagecode#", "'"+languageCode+"'")
	} else {
		sql = strings.ReplaceAll(sql, "_i18n", " ")
		sql = strings.ReplaceAll(sql, "i18n.languagecode=#languagecode#", " 1=1 ")

		sql = strings.ReplaceAll(sql, "languagecode=#languagecode#", " 1=1 ")
	}
	return sql
}

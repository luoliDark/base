package xormutil

import (
	"base/base/db/conn"
	"reflect"

	"github.com/xormplus/xorm"
)

func UpdateAllSession(entity interface{}, session xorm.Session) error {
	e := reflect.ValueOf(entity).Elem()

	fields := make([]string, 0)
	for i := 0; i < e.NumField(); i++ {
		fieldName := e.Type().Field(i).Name
		fields = append(fields, fieldName)
	}
	_, err := session.Cols(fields...).Update(entity)
	return err
}

func UpdateAllEngine(entity interface{}, engine xorm.Engine) error {
	e := reflect.ValueOf(entity).Elem()

	fields := make([]string, 0)
	for i := 0; i < e.NumField(); i++ {
		fieldName := e.Type().Field(i).Name
		fields = append(fields, fieldName)
	}
	_, err := engine.Cols(fields...).Update(entity)
	return err
}

func QueryTableAll(tableName string) ([]map[string]interface{}, error) {
	eng, err := conn.GetDB()
	if err != nil {
		return nil, err
	}
	return eng.SQL("select * from " + tableName).Query().List()
}

package log

import (
	"github.com/luoliDark/base/db/conn/logserverconn"
)

type Logger interface {
	SaveLog() error
}

type LogBaseInfo struct {
	Id        int64
	Title     string `xorm:"Title" json:"title"`
	LogMsg    string `xorm:"LogMsg" json:"logMsg"`
	Uid       string `xorm:"Uid" json:"uid"`
	IP        string `xorm:"Ip" json:"ip"`
	FileName  string `xorm:"Filename" json:"fileName"`
	Extend1   string `xorm:"Extend1" json:"extend1"`
	Extend2   string `xorm:"Extend2" json:"extend2"`
	CreatedAt int64  `xorm:"CreatedAt" json:"createdAt"`
}

func (l *LogBaseInfo) SaveLog() error {
	return nil
}

type LogByInfo struct {
	LogBaseInfo `xorm:"extends"`
}

func (*LogByInfo) TableName() string {
	return "log_by_info"
}

func (l *LogByInfo) SaveLog() error {
	engine, _ := logserverconn.GetDB()

	_, err := engine.Insert(l)

	return err
}

type LogByError struct {
	LogBaseInfo `xorm:"extends"`
}

func (*LogByError) TableName() string {
	return "log_by_error"
}

func (l *LogByError) SaveLog() error {
	engine, _ := logserverconn.GetDB()

	_, err := engine.Insert(l)

	return err
}

type LogByHighError struct {
	LogBaseInfo `xorm:"extends"`
}

func (*LogByHighError) TableName() string {
	return "log_by_higherror"
}

func (l *LogByHighError) SaveLog() error {
	engine, _ := logserverconn.GetDB()

	_, err := engine.Insert(l)

	return err
}

type LogByConfig struct {
	LogBaseInfo `xorm:"extends"`
}

func (*LogByConfig) TableName() string {
	return "log_by_config"
}

func (l *LogByConfig) SaveLog() error {
	engine, _ := logserverconn.GetDB()

	_, err := engine.Insert(l)

	return err
}

type LogByTimeOut struct {
	LogBaseInfo `xorm:"extends"`
}

func (*LogByTimeOut) TableName() string {
	return "log_by_timeout"
}

func (l *LogByTimeOut) SaveLog() error {
	engine, _ := logserverconn.GetDB()

	_, err := engine.Insert(l)

	return err
}

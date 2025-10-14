package dateutil

import (
	"EasyFinance2020/base/util/commutil"
	"errors"
	"strconv"
	"strings"
	"time"
)

// 结算日期 天模式
type SettleDayHandler struct{}

const (
	SettlePatternDay   = "day"
	SettlePatternMonth = "month"
	SettleDay          = "T"
	SettleWeek         = "W"
	SettleMonth        = "M"
)

func (sh *SettleDayHandler) calculate(payTime time.Time, cycle int, pattern string, isSum bool) ([]string, error) {
	var settleDateList []string
	if isSum {
		for i := 0; i < cycle; i++ {
			d, _ := time.ParseDuration("-24h")
			payTime = payTime.Add(d)
			format := payTime.Format(TimeLayoutDate)
			settleDateList = append(settleDateList, format)
		}
	} else {
		// 如果不是汇总则返回周期日期
		dateCycle := "-" + commutil.ToString(cycle*24) + "h"
		d, _ := time.ParseDuration(dateCycle)
		payTime = payTime.Add(d)
		format := payTime.Format(TimeLayoutDate)
		settleDateList = append(settleDateList, format)
	}
	var result []string
	if len(settleDateList) >= 2 {
		result = append(result, settleDateList[len(settleDateList)-1], settleDateList[0])
	} else if len(settleDateList) == 1 {
		result = append(result, settleDateList[0], settleDateList[0])
	}

	return result, nil
}

// 获取结算模式 周模式
type SettleWeekHandler struct{}

func (sh *SettleWeekHandler) calculate(payTime time.Time, cycle int, pattern string, isSum bool) ([]string, error) {
	payTime = payTime.AddDate(0, 0, -7*cycle)
	var result []string
	// 开始时间
	start, err := GetLastWeekMonday(payTime, "2006-01-02")
	if err != nil {
		return nil, err
	}
	end, err := GetLastWeekSunday(payTime, "2006-01-02")
	if err != nil {
		return nil, err
	}

	result = append(result, start, end)

	return result, nil
}

// 通过年月获取当前月第一天
func GetMonthFirstDateTimeWithFormat(yearMonth, format string) (string, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return "", err
	}
	theTime, _ := time.ParseInLocation(TimeLayoutDateTime, yearMonth+"-01 00:00:00", loc)
	newMonth := theTime.Month()
	year := theTime.Year()
	dateTime := time.Date(year, newMonth, 1, 0, 0, 0, 0, time.Local).Format(format)
	return dateTime, nil
}

// 通过年月获取当前月最后一天最后时刻
func GetMonthLastDateTimeWithFormat(yearMonth, format string) (string, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return "", err
	}
	theTime, _ := time.ParseInLocation(TimeLayoutDateTime, yearMonth+"-01 00:00:00", loc)
	newMonth := theTime.Month()
	year := theTime.Year()
	dateTime := time.Date(year, newMonth+1, 0, 23, 59, 59, 0, time.Local).Format(format)
	return dateTime, nil
}

// 获取一个月第一天和最后一天
func GetMonthFirstAndLastDayWithFormat(yearMonth, format string) (string, string, error) {
	firstDay, err := GetMonthFirstDateTimeWithFormat(yearMonth, format)
	if err != nil {
		return "", "", err
	}
	lastDay, err := GetMonthLastDateTimeWithFormat(yearMonth, format)
	if err != nil {
		return "", "", err
	}
	return firstDay, lastDay, nil
}

type SettleMonthHandler struct{}

func (sh *SettleMonthHandler) calculate(payTime time.Time, cycle int, pattern string, isSum bool) ([]string, error) {
	var result []string
	// 获取上个月的年份和月份
	yearMonth := payTime.AddDate(0, -1*cycle, 0).Format("2006-01")
	start, end, err := GetMonthFirstAndLastDayWithFormat(yearMonth, TimeLayoutDate)
	if err != nil {
		return nil, err
	}
	result = append(result, start, end)

	return result, nil
}

func GetSettleDate(settleCycle string) string {
	if strings.Contains(settleCycle, "sum") {
		return "SUM"
	}
	return "NORMAL"
}

type SettleDateModel struct {
	DataList      []string
	CycleModel    string
	MonthStrRange []string
	DyaStrRange   []string
}

func (sd *SettleDateModel) GetKey() string {
	var str string
	str = sd.CycleModel
	if sd.CycleModel == SettlePatternDay {
		str += "^" + strings.Join(sd.DyaStrRange, "^")
	} else if sd.CycleModel == SettlePatternMonth {
		str += "^" + strings.Join(sd.MonthStrRange, "^")
	}

	return str
}

func SettleDateWithSettleRange(payTimeStr, settleCycle string) (*SettleDateModel, error) {
	dateList, cycleModel, _, _, err := SettleDateV2(payTimeStr, settleCycle)
	if err != nil {
		return nil, err
	}
	if len(dateList) <= 1 {
		return nil, nil
	}
	result := &SettleDateModel{}
	result.DataList = dateList
	result.CycleModel = cycleModel
	startTime, err := time.Parse(TimeLayoutDate, dateList[0])
	endTime, err := time.Parse(TimeLayoutDate, dateList[len(dateList)-1])
	if err != nil {
		return nil, err
	}
	if result.CycleModel == SettlePatternMonth {
		result.MonthStrRange = append(result.MonthStrRange, startTime.Format("200601"), endTime.Format("200601"))
	} else if result.CycleModel == SettlePatternDay {
		result.DyaStrRange = append(result.DyaStrRange, startTime.Format("20060102"), endTime.Format("20060102"))
	}
	return result, nil
}

// SettleDate
//
//	Description: 获取结算时间，
//	param payTimeStr 支付时间 举例：2023-10-23
//	param settleCycle  模式  T = 日  W=周  M=月  SUM = 总和日期（目前只有T有）  举例：T+1（前一天） W+1（上个星期一，上个星期行日日）  M+1（上个月1号，上个月最后一天）
//	return []string  只返回开始日期和结束日期，如果为一天 则开始时间=结束时间
//	return error
func SettleDateV2(payTimeStr, settleCycle string) ([]string, string, string, string, error) {
	if settleCycle == "" {
		return nil, "", "", "", errors.New("结算周期为空")
	}
	cycleModel := ""
	payTime, err := time.Parse(TimeLayoutDate, payTimeStr)
	if err != nil {
		return nil, "", "", "", err
	}
	isSum := false
	var pattern string
	// 判断是否有sum  sum(T+3)
	if strings.Contains(settleCycle, "sum") {
		isSum = true
		sumStart := strings.Index(settleCycle, "sum(")
		sumEnd := strings.LastIndex(settleCycle, ")")
		if sumStart >= 0 && sumEnd > 0 {
			settleCycle = settleCycle[sumStart+4 : sumEnd]
		}
	}
	// 获取模式
	pattern = string(settleCycle[0])
	var handler settlePatternInterface
	cycle, err := strconv.Atoi(settleCycle[strings.Index(settleCycle, "+"):])
	if err != nil {
		return nil, "", "", "", err
	}
	pattern = strings.ToUpper(pattern)
	// 判断模式
	if pattern == SettleDay {
		// 天模式
		cycleModel = SettlePatternDay
		handler = &SettleDayHandler{}
	} else if pattern == SettleWeek {
		handler = &SettleWeekHandler{}
		cycleModel = "week"
	} else if pattern == SettleMonth {
		cycleModel = SettlePatternMonth
		handler = &SettleMonthHandler{}
	}
	var monthStr, dayStr string

	ll, err := handler.calculate(payTime, cycle, cycleModel, isSum)
	if len(ll) > 0 {
		tt, _ := time.Parse(TimeLayoutDate, ll[0])
		dayStr = tt.Format("20060102")
		monthStr = tt.Format("200601")
	}

	return ll, cycleModel, monthStr, dayStr, err
}

type settlePatternInterface interface {
	calculate(payTime time.Time, cycle int, pattern string, isSum bool) ([]string, error)
}

// SettleDate
//
//	Description: 获取结算时间，
//	param payTimeStr 支付时间 举例：2023-10-23
//	param settleCycle  模式  T = 日  W=周  M=月  SUM = 总和日期（目前只有T有）  举例：T+1（前一天） W+1（上个星期一，上个星期行日日）  M+1（上个月1号，上个月最后一天）
//	return []string  只返回开始日期和结束日期，如果为一天 则开始时间=结束时间
//	return error
func SettleDate(payTimeStr, settleCycle string) ([]string, error) {
	if settleCycle == "" {
		return nil, errors.New("结算周期为空")
	}

	payTime, err := time.Parse(TimeLayoutDate, payTimeStr)
	if err != nil {
		return nil, err
	}
	isSum := false
	var pattern string
	// 判断是否有sum  sum(T+3)
	if strings.Contains(settleCycle, "sum") {
		isSum = true
		sumStart := strings.Index(settleCycle, "sum(")
		sumEnd := strings.LastIndex(settleCycle, ")")
		if sumStart >= 0 && sumEnd > 0 {
			settleCycle = settleCycle[sumStart+4 : sumEnd]
		}
	}
	// 获取模式
	pattern = string(settleCycle[0])
	var handler settlePatternInterface
	cycle, err := strconv.Atoi(settleCycle[strings.Index(settleCycle, "+"):])
	if err != nil {
		return nil, err
	}
	pattern = strings.ToUpper(pattern)
	// 判断模式
	if pattern == SettleDay {
		// 天模式
		handler = &SettleDayHandler{}
	} else if pattern == SettleWeek {
		handler = &SettleWeekHandler{}
	} else if pattern == SettleMonth {
		handler = &SettleMonthHandler{}
	}

	return handler.calculate(payTime, cycle, pattern, isSum)
}

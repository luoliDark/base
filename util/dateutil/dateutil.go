package dateutil

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/luoliDark/base/util/commutil"

	"github.com/gogf/gf/frame/g"
)

var (
	TimeLayoutDate      = "2006-01-02"
	TimeLayoutDateTime  = "2006-01-02 15:04:05"
	TimeLayoutDateTimeA = "20060102150405"
)

func MonthStart() time.Time {
	y, m, _ := time.Now().Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
}

func TodayStart() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

func TodayEnd() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 23, 59, 59, 1e9-1, time.Local)
}

func NowUnix() int64 {
	return time.Now().Unix()
}

func NowDate() string {
	//默认查明天因为 当天如果12点时 会查不到票
	return time.Now().AddDate(0, 0, 1).Format(TimeLayoutDate)
}

func NowDateTime() string {
	return time.Now().Format(TimeLayoutDateTime)
}

func NowDateTimeCustom(timeLayout string) string {
	return time.Now().Format(timeLayout)
}

func ParseDate(dt string) (time.Time, error) {
	return time.Parse(TimeLayoutDate, dt)
}

func ParseDayToInt(dt string) (int, error) {
	t, err := time.Parse(TimeLayoutDate, dt)
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(t.Format("20060102"))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func ParseDateTime(dt string) (time.Time, error) {
	return time.Parse(TimeLayoutDateTime, dt)
}

func ParseStringTime(tm, lc string) (time.Time, error) {
	loc, err := time.LoadLocation(lc)
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(TimeLayoutDateTime, tm, loc)
}

// 获取两个月份之间的第一天和最后一天
func GetMonthBettenFirstAndLastDay(startmonth, endmonth string) (string, string, error) {
	firstDay, err := GetMonthFirstDateTime(startmonth)
	if err != nil {
		return "", "", err
	}
	lastDay, err := GetMonthLastDateTime(endmonth)
	if err != nil {
		return "", "", err
	}
	return firstDay, lastDay, nil
}

// 获取一个月第一天和最后一天
func GetMonthFirstAndLastDay(yearMonth string) (string, string, error) {
	firstDay, err := GetMonthFirstDateTime(yearMonth)
	if err != nil {
		return "", "", err
	}
	lastDay, err := GetMonthLastDateTime(yearMonth)
	if err != nil {
		return "", "", err
	}
	return firstDay, lastDay, nil
}

// 通过年月获取当前月第一天
func GetMonthFirstDateTime(yearMonth string) (string, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return "", err
	}
	theTime, _ := time.ParseInLocation(TimeLayoutDateTime, yearMonth+"-01 00:00:00", loc)
	newMonth := theTime.Month()
	year := theTime.Year()
	dateTime := time.Date(year, newMonth, 1, 0, 0, 0, 0, time.Local).Format(TimeLayoutDateTime)
	return dateTime, nil
}

// 通过年月获取当前月最后一天最后时刻
func GetMonthLastDateTime(yearMonth string) (string, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return "", err
	}
	theTime, _ := time.ParseInLocation(TimeLayoutDateTime, yearMonth+"-01 00:00:00", loc)
	newMonth := theTime.Month()
	year := theTime.Year()
	dateTime := time.Date(year, newMonth+1, 0, 23, 59, 59, 0, time.Local).Format(TimeLayoutDateTime)
	return dateTime, nil
}

func IsValidTime(t time.Time) bool {
	if t.IsZero() {
		return false
	}
	if t.Unix() <= 0 {
		return false
	}
	return true
}

// SinceForHuman 1小时前 -> 这样的展示方式
func SinceForHuman(t time.Time) string {
	duration := time.Since(t)
	hour := duration.Hours()
	minutes := duration.Minutes()
	seconds := duration.Seconds()

	unit := "秒"
	s := 0
	if hour > (365 * 24) {
		s = int(math.Floor(hour / 365))
		unit = "年"
	} else if hour > 30 {
		s = int(math.Floor(hour / 30))
		unit = "月"
	} else if hour > 1 {
		s = int(math.Floor(hour))
		unit = "小时"
	} else if minutes > 1 {
		s = int(math.Floor(minutes))
		unit = "分钟"
	} else if seconds > 0 {
		return "刚刚"
	}

	return strconv.Itoa(s) + unit + "前"
}

// 针对读出的日期是数字时直接转
func NumberToDate(number string) time.Time {
	excelTime := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	var days, _ = strconv.Atoi(number)
	return excelTime.Add(time.Second * time.Duration(days*86400))
}

func StrToDateDay(datestr string) string {

	datestr = strings.TrimLeft(datestr, " ")
	datestr = strings.TrimRight(datestr, " ")

	if !g.IsEmpty(datestr) {
		if strings.Index(datestr, "-") == -1 {
			//验证日期格式中是否含有年月日中文
			if strings.Index(datestr, "年") != -1 {
				//如果日期中有年月日中文，把他替换成"-"
				datestr = strings.Replace(datestr, "年", "-", -1)
				datestr = strings.Replace(datestr, "月", "-", -1)
				datestr = strings.Replace(datestr, "日", "", -1)
				//取年月日 去掉时分秒
				datestr = strings.Split(datestr, " ")[0]
				//验证月,日补0
				datestrList := strings.Split(datestr, "-")
				month := ""
				day := ""

				if len(datestrList) <= 1 {
					fmt.Println("非法日期格式")
					return ""
				}

				if len(datestrList[1]) == 1 {
					month = "0" + datestrList[1]
				} else {
					month = datestrList[1]
				}

				if len(datestrList[2]) == 1 {
					day = "0" + datestrList[2]
				} else {
					day = datestrList[2]
				}
				datestr = datestrList[0] + "-" + month + "-" + day
			} else {
				//取前8位
				if len(datestr) > 8 {
					datestr = datestr[0:8]
					year := datestr[0:4]
					month := datestr[4:6]
					day := datestr[6:8]
					datestr = year + "-" + month + "-" + day
				} else {
					//将数字转日期
					t := NumberToDate(datestr)
					datestr = t.Format(commutil.Time_Fomat03)
				}

			}
		} else {
			var datestrList []string
			//日期格式中包含"-",把年月日中文替换成空字符串
			if strings.Index(datestr, "年") != -1 {
				datestr = strings.Replace(datestr, "年", "", -1)
				datestr = strings.Replace(datestr, "月", "", -1)
				datestr = strings.Replace(datestr, "日", "", -1)
				datestr = strings.Split(datestr, " ")[0]
				datestrList = strings.Split(datestr, "-")
			} else {
				datestr = strings.Split(datestr, " ")[0]
				datestrList = strings.Split(datestr, "-")
			}
			month := ""
			day := ""

			if len(datestrList) <= 1 {
				fmt.Println("非法日期格式")
				return ""
			}

			if len(datestrList[1]) == 1 {
				month = "0" + datestrList[1]
			} else {
				month = datestrList[1]
			}

			if len(datestrList) == 3 {
				if len(datestrList[2]) == 1 {
					day = "0" + datestrList[2]
				} else {
					day = datestrList[2]
				}
			} else {
				day = "01"
			}

			datestr = datestrList[0] + "-" + month + "-" + day
		}
	}
	return datestr
}

// 获取事件列表起始区间
func SearchInterval(timeList []string) (string, string) {
	if len(timeList) <= 0 {
		return "", ""
	}
	//var startTimeStr,endTimeStr string
	var startTime, endTime time.Time
	index := 0
	for _, timeStr := range timeList {
		if commutil.IsNullOrEmpty(timeStr) {
			continue
		}
		timeStr := timeStr[0:10]
		if commutil.IsNullOrEmpty(timeStr) {
			continue
		}
		parse, _ := time.Parse(TimeLayoutDate, timeStr)
		if index == 0 {
			//startTimeStr =timeStr
			//endTimeStr = timeStr
			startTime = parse
			endTime = parse
			index++
			continue
		}

		if parse.Before(startTime) {
			startTime = parse
			//startTimeStr = timeStr
		}

		if parse.After(endTime) {
			endTime = parse
			//endTimeStr = timeStr
		}

	}
	return startTime.Format(TimeLayoutDate) + " 00:00:00", endTime.Format(TimeLayoutDate) + " 23:59:59"
}

// 判断时间天数是否相等
func JudgeDayIsSame(times ...string) bool {
	sample, _ := time.Parse(TimeLayoutDate, times[0])
	for i := 0; i < len(times); i++ {
		s := times[i][0:10]
		parse, err := time.Parse(TimeLayoutDate, s)
		if err != nil {
			return false
		}
		if i == 0 {
			continue
		}
		equal := sample.Equal(parse)
		if !equal {
			return equal
		}
	}
	return true
}

func CustomTodayStart(datetime, layout string) time.Time {
	//y, m, d := time.Now().Date()
	parse, _ := time.Parse(layout, datetime)
	year, month, day := parse.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func CustomTodayEnd(datetime, layout string) time.Time {
	parse, _ := time.Parse(layout, datetime)
	year, month, day := parse.Date()
	return time.Date(year, month, day, 23, 59, 59, 1e9-1, time.Local)
}

func GetEveryDayList(year, month int, formatStr string) (list []string, err error) {
	start, end, err := GetMonthStartAndEnd(year, month)
	if err != nil {
		return nil, err
	}
	list, err = GetEveryDayStrByStartEndTime(start, end, formatStr)
	return
}

func GetMonthStartAndEnd(yearInt, monthInt int) (int64, int64, error) {
	// 数字月份必须前置补零
	year := strconv.Itoa(yearInt)
	mouth := strconv.Itoa(int(monthInt))

	if len(mouth) == 1 {
		mouth = "0" + mouth
	}

	timeLayout := "2006-01-02 15:04:05"
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return 0, 0, err
	}
	theTime, _ := time.ParseInLocation(timeLayout, year+"-"+mouth+"-01 00:00:00", loc)
	newMonth := theTime.Month()

	t1 := time.Date(yearInt, newMonth, 1, 0, 0, 0, 0, time.Local)
	t2 := time.Date(yearInt, newMonth+1, 0, 23, 59, 59, 0, time.Local)

	fmt.Println(t1.Unix(), "---", t2.Unix())

	return t1.Unix(), t2.Unix(), nil
}

func GetEveryDayStrByStartEndTime(start, end int64, formatStr string) (list []string, err error) {
	var (
		startTime time.Time
		endTime   time.Time
	)

	startTime = time.Unix(start, 0)
	endTime = time.Unix(end, 0)

	if err != nil {
		return
	}

	for {
		if startTime.After(endTime) {
			break
		}
		str := startTime.Format(formatStr)
		list = append(list, str)
		startTime = startTime.AddDate(0, 0, 1)
	}

	return
}

// 获取某一天的0点时间
func GetZeroTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// 获取本周周一的日期
func GetMondayOfWeek(t time.Time, fmtStr string) (dayStr string) {
	dayObj := GetZeroTime(t)
	if t.Weekday() == time.Monday {
		//修改hour、min、sec = 0后格式化
		dayStr = dayObj.Format(fmtStr)
	} else {
		offset := int(time.Monday - t.Weekday())
		if offset > 0 {
			offset = -6
		}
		dayStr = dayObj.AddDate(0, 0, offset).Format(fmtStr)
	}
	return
}

// 获取上周周一日期
func GetLastWeekMonday(t time.Time, fmtStr string) (day string, err error) {
	monday := GetMondayOfWeek(t, fmtStr)
	dayObj, err := time.Parse(fmtStr, monday)
	if err != nil {
		return
	}
	day = dayObj.AddDate(0, 0, -7).Format(fmtStr)
	return
}

// 获取上周周日日期
func GetLastWeekSunday(t time.Time, fmtStr string) (day string, err error) {
	monday := GetMondayOfWeek(t, fmtStr)
	dayObj, err := time.Parse(fmtStr, monday)
	if err != nil {
		return
	}
	day = dayObj.AddDate(0, 0, -1).Format(fmtStr)
	return
}

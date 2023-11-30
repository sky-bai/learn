package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	keyTime := GetIccIdKeyByDate(Timestamp())
	fmt.Println("keyTime=", keyTime)
	//"cfFlowIccid_89860621260096013412_20231019"
	//SCAN 0 MATCH cfFlowIccid_89860619130109955004_20231019_ruichiman_0_gps
	test()
	// 869497052085063
	//89860621260042715391

	//"cfFlowIccidSoredSet_20231019_gps"
	//ZRANGEBYSCORE cfFlowIccidSoredSet_20231019_gps (1697698256647715000 (1697698433678587000

	fmt.Println("111", time.Now().UnixNano())
}

func test() {
	ti := time.Now().Format("2006_01_02")

	fmt.Println(ti + "_in_pgs")
	fmt.Println(ti + "_out_gps")

}

// 获取每天的时间戳
func GetIccIdKeyByDate(ts int64) string {
	time := Date("Ymd", ts)
	return time
}

// Timestamp / 时间戳（ms）
func Timestamp() int64 {
	return time.Now().UnixNano() / 1e6
	//return time.Now().Unix()
}

func Date(f string, timestamp int64) string {
	if len(IntToStr(int(timestamp))) == 10 {
		return format(f, time.Unix(timestamp, 0).Local())
	} else { // 添加毫秒判断
		return format(f, time.UnixMilli(timestamp).Local())
	}

}

func IntToStr(num int) string {
	res := strconv.Itoa(num)
	return res
}

func format(f string, t time.Time) string {
	pattern, err := getPattern(f)

	if err != nil {
		return parse(f, t)
	}

	return t.Format(pattern.layout)
}

// getPattern gets the matched pattern by the given php style date/time format string
func getPattern(format string) (pattern, error) {
	for _, p := range _defaultPatterns {
		if p.format == format {
			return p, nil
		}
	}

	return pattern{}, errors.New("No pattern found")
}

func parse(format string, t time.Time) string {
	result := ""
	for _, s := range format {
		result += recognize(string(s), t)
	}
	return result
}

// pattern stores the mapping of golang datetime layout and php datetime format
type pattern struct {
	regexp string // golang regular expression
	layout string // golang datetime layout
	format string // php datetime format
}

// patterns is an array of pattern
type patterns []pattern

var _defaultPatterns patterns

// recognize the character in the php date/time format string
func recognize(c string, t time.Time) string {
	switch c {
	// Day
	case "d": // Day of the month, 2 digits with leading zeros
		return fmt.Sprintf("%02d", t.Day())
	case "D": // A textual representation of a day, three letters
		return t.Format("Mon")
	case "j": // Day of the month without leading zeros
		return fmt.Sprintf("%d", t.Day())
	case "l": // A full textual representation of the day of the week
		return t.Weekday().String()
	case "w": // Numeric representation of the day of the week
		return fmt.Sprintf("%d", t.Weekday())
	case "z": // The day of the year (starting from 0)
		return fmt.Sprintf("%v", t.YearDay()-1)

	// Week
	case "W": // ISO-8601 week number of year, weeks starting on Monday
		_, w := t.ISOWeek()
		return fmt.Sprintf("%d", w)

	// Month
	case "F": // A full textual representation of a month
		return t.Month().String()
	case "m": // Numeric representation of a month, with leading zeros
		return fmt.Sprintf("%02d", t.Month())
	case "M": // A short textual representation of a month, three letters
		return t.Format("Jan")
	case "n": // Numeric representation of a month, without leading zeros
		return fmt.Sprintf("%d", t.Month())
	case "t": // Number of days in the given month
		return LastDateOfMonth(t).Format("2")
	// Year
	case "L": // Whether it's a leap year
		if IsLeapYear(t) {
			return "1"
		}
		return "0"
	case "o":
		fallthrough
	case "Y": // A full numeric representation of a year, 4 digits
		return fmt.Sprintf("%v", t.Year())
	case "y": // A two digit representation of a year
		return t.Format("06")

	// Time
	case "a": // Lowercase Ante meridiem and Post meridiem
		return t.Format("pm")
	case "A": // Uppercase Ante meridiem and Post meridiem
		return strings.ToUpper(t.Format("pm"))
	case "g": // 12-hour format of an hour without leading zeros
		return t.Format("3")
	case "G": // 24-hour format of an hour without leading zeros
		return fmt.Sprintf("%d", t.Hour())
	case "h": // 12-hour format of an hour with leading zeros
		return t.Format("03")
	case "H": // 24-hour format of an hour with leading zeros
		return fmt.Sprintf("%02d", t.Hour())
	case "i": // Minutes with leading zeros
		return fmt.Sprintf("%02d", t.Minute())
	case "s": // Seconds, with leading zeros
		return fmt.Sprintf("%02d", t.Second())
	case "e": // Timezone identifier
		fallthrough
	case "T": // Timezone abbreviation
		return t.Format("MST")
	case "O": // Difference to Greenwich time (GMT) in hours
		return t.Format("-0700")
	case "P": // Difference to Greenwich time (GMT) with colon between hours and minutes
		return t.Format("-07:00")
	case "U": // Seconds since the Unix Epoch (January 1 1970 00:00:00 GMT)
		return fmt.Sprintf("%v", t.Unix())

	default:
		return c
	}
}

func LastDateOfMonth(t time.Time) time.Time {
	t2 := FirstDateOfNextMonth(t)
	return time.Unix(t2.Unix()-86400, 0)
}

func FirstDateOfNextMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	if month == time.December {
		year++
		month = time.January
	} else {
		month++
	}
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}

func IsLeapYear(t time.Time) bool {
	t2 := time.Date(t.Year(), time.December, 31, 0, 0, 0, 0, time.UTC)
	return t2.YearDay() == 366
}

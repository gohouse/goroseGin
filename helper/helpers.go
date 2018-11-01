package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/gohouse/gorose/utils"
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"time"
	"golang.org/x/text/transform"
	"strings"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"path"
)

func MapValueToString(res map[string]interface{}) map[string]interface{} {
	if len(res) > 0 {
		for k, v := range res {
			switch v.(type) {
			case nil:
				res[k] = ""
			default:
				res[k] = utils.ParseStr(v)
				if res[k] == "NULL" {
					res[k] = ""
				}
			}
		}
		return res
	}
	return nil
}

func MapValueNilToNull(res map[string]interface{}) map[string]interface{} {
	if len(res) > 0 {
		for k, v := range res {
			switch v.(type) {
			case nil:
				res[k] = ""
			}
		}
	}

	return res
}

func SliceMapValueToString(res []map[string]interface{}) []map[string]interface{} {
	if len(res) > 0 {
		for index, item := range res {
			res[index] = MapValueToString(item)
			//if len(item) > 0 {
			//	for k, v := range item {
			//		switch v.(type) {
			//		case nil:
			//			res[index][k] = ""
			//		default:
			//			res[index][k] = utils.ParseStr(v)
			//		}
			//	}
			//}
		}
	}
	return res
}

func FileDownload(c *gin.Context, myFile string) utils.ApiReturn {
	if myFile == "" {
		return utils.FailReturn("file is empty")
	}

	file, err := os.Open(myFile)
	if err != nil {
		return utils.FailReturn(err.Error())
	}
	defer file.Close()

	// 截取文件名
	fileName := path.Base(myFile)

	w := c.Writer
	w.Header().Add("Content-type", "application/octet-stream")
	c.Header("Content-Description", "File Transfer")
	w.Header().Add("Content-Disposition", "attachment; filename= filename=\""+fileName+"\"")

	io.Copy(w, file)

	return utils.SuccessReturn()
}

type DateTime struct {
	LastMonthStart string
	LastMonthEnd   string
	ThisMonthStart string
	TodayStart     string
	YesterdayStart string
	ThisWeekStart  string
	LastWeekStart  string
	LastWeekEnd    string
	Now            string
}

func GetDate() DateTime {
	const DATE_FORMAT = "2006-01-02"
	const DATETIME_FORMAT = "2006-01-02 15:04:05"
	// 现在
	now := time.Now()
	// 年月日
	year, month, day := now.Date()
	// 今天开始
	today := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	// 本月开始
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	// 昨天开始
	yeaterday := today.AddDate(0, 0, -1)
	// 周计算
	weekDay := now.Weekday()
	//fmt.Println(weekDay)
	// 本周开始
	thisWeekStart := today.AddDate(0, 0, -int(weekDay)+1)
	// 上周开始
	lastWeekStart := thisWeekStart.AddDate(0, 0, -7)
	// 上周结束
	//lastWeekEnd := thisWeekStart.AddDate(0, 0, -1)
	lastWeekEnd := thisWeekStart.Add(-1 * time.Second)
	//fmt.Println(lastWeekEnd)
	//time.Weekday
	return DateTime{
		LastMonthStart: thisMonth.AddDate(0, -1, 0).Format(DATE_FORMAT),
		LastMonthEnd:   thisMonth.AddDate(0, 0, -1).Format(DATE_FORMAT),
		ThisMonthStart: thisMonth.Format(DATE_FORMAT),
		TodayStart:     today.Format(DATE_FORMAT + " 00:00:00"),
		YesterdayStart: yeaterday.Format(DATE_FORMAT + " 00:00:00"),
		ThisWeekStart:  thisWeekStart.Format(DATE_FORMAT + " 00:00:00"),
		LastWeekStart:  lastWeekStart.Format(DATE_FORMAT + " 00:00:00"),
		LastWeekEnd:    lastWeekEnd.Format(DATETIME_FORMAT),
		Now:            now.Format(DATETIME_FORMAT),
	}
}

type DayDatetime struct {
	DateStart string
	DateEnd   string
}

func GetDateStartAndEndByDateTime(datetime string) DayDatetime {
	var day DayDatetime
	var datetime_format = "2006-01-02 15:04:05"
	// ===============
	res2, _ := time.Parse(datetime_format, datetime)
	y, m, d := res2.Date()
	dayStart := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	dayEnd := time.Date(y, m, d, 23, 59, 59, 999, time.Local)
	day.DateStart = dayStart.Format(datetime_format)
	day.DateEnd = dayEnd.Format(datetime_format)

	return day
}

func UTF82GBK(src string) (string, error) {
	reader := transform.NewReader(strings.NewReader(src), simplifiedchinese.GBK.NewEncoder())
	if buf, err := ioutil.ReadAll(reader); err != nil {
		return "", err
	} else {
		return string(buf), nil
	}
}

func Md5(str interface{}) string {
	data := []byte(fmt.Sprintf("%v", str))
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制

	return md5str
}

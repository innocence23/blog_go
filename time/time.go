package main

import (
	"fmt"
	"time"
)

func main() {

	// 默认UTC    国籍标准时间
	// GMT   	格林尼治标准时间     和utc 相等吧
	// CST       北京时间 哈哈哈哈  China Standard Time，  和上面查8小时


	//当前时间
	t := time.Now()
	fmt.Println(t)

	// 时间转换字符处
	fmt.Println(t.Format("2016-01-02 15:04:05"))

	//时间戳 转换字符串 毫秒数
	var _startDate int64 = time.Now().Unix()
	var startDate string = time.Unix(_startDate, 0).Format("2006-01-02 15:04:05")
	fmt.Println(startDate)


	//字符串 转换成时间格式

	layout := "2006-01-02 15:04:05"
	t, _ = time.Parse(layout, "2013-10-05 18:30:50")
	fmt.Println(t.Year())
	fmt.Println(t)  // 2013-10-05 18:30:50 +0000 UTC

	// 当前时区时间格式化转换
	strTime := "2018-03-24T20:01:00+08:00"
	//tim, _ := time.ParseInLocation("2006-01-02T15:04:05+08:00", strTime, time.Local)
	loc, _ := time.LoadLocation("PRC") // 2018-03-24 20:01:00 +0800 CST
	loc, _ = time.LoadLocation("UTC") // 2018-03-24 20:01:00 +0000 UTC
	//todo
	//todo  只是时区的差别
	tim, _ := time.ParseInLocation("2006-01-02T15:04:05+08:00", strTime, loc)
	fmt.Println(tim.Year())
	fmt.Println(tim) //2018-03-24 20:01:00 +0800 CST
	fmt.Println(time.Local)


	//时区问题
	// func (t Time) In(loc *Location) Time 当前时间对应指定时区的时间
	loc, _ = time.LoadLocation("America/Los_Angeles")
	fmt.Println(time.Now().In(loc))
}


/**

可以看看

https://blog.csdn.net/qq_26981997/article/details/53454606

 */
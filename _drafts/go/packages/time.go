package main

import (
	"fmt"
	"time"
)

func main()  {

	var t time.Time
	t = time.Date(2019, 10, 8, 10, 5, 11, 0, time.UTC)
	now := time.Now()
	fmt.Println("Now()", time.Now())
	fmt.Println("Now()", now)
	fmt.Println("Date()", t)

	fmt.Println()
	fmt.Println("Now().Unix()",time.Now().Unix())
	fmt.Println("Now().UnixNano()", time.Now().UnixNano())

	fmt.Println()
	fmt.Println("Local()", t.Local())
	fmt.Println("Location()", t.Location())

	fmt.Println()
	fmt.Println("Year()", t.Year())
	year, month, day := t.Date()
	fmt.Println("Date()", year, month, day)
	fmt.Printf("Date() format: %d-%d-%d\n",year, month, day)
	fmt.Printf("Date() format2: %d-%02d-%02d\n",year, month, day)
	fmt.Println("Month()", t.Month())
	fmt.Println("Weekday()", t.Weekday())
	year, week := t.ISOWeek()
	fmt.Println("ISOWeek()", year, week)
	fmt.Println("Day()", t.Day())
	hour, min, sec := t.Clock()
	fmt.Println("Clock()", hour, min, sec)
	fmt.Printf("Clock() format: %d-%d-%d\n", hour, min, sec)
	fmt.Printf("Clock() format2: %02d-%02d-%02d\n", hour, min, sec)
	fmt.Println("Hour()", t.Hour())
	fmt.Println("Minute()", t.Minute())
	fmt.Println("Second()", t.Second())

	fmt.Println("")
	// 必须是这个时间 2006-01-02 15:04:05 golang的诞生日，否则会出错
	fmt.Println("Format()", t.Format("2006-01-02 15:04:05"))
	fmt.Println("Formant() err", t.Format("2007-01-02 12:02:03"))

	fmt.Println("")
	duration, _ := time.ParseDuration("-1m")
	t1 := t.Add(duration)
	fmt.Println("ParseDuration(\"-1m\")", t1)

	duration2, _ := time.ParseDuration("-1h")
	t2 := t.Add(duration2)
	fmt.Println("ParseDuration(\"-1h\")", t2)

	duration3, _ := time.ParseDuration("-24h")
	t3 := t.Add(duration3)
	fmt.Println("ParseDuration(\"-24h\")", t3)

	duration4, _ := time.ParseDuration("1m")
	t4 := t.Add(duration4)
	fmt.Println("ParseDuration(\"1m\")", t4)

	duration5, _ := time.ParseDuration("1h")
	t5 := t.Add(duration5)
	fmt.Println("ParseDuration(\"1h\")", t5)

	duration6, _ := time.ParseDuration("24h")
	t6 := t.Add(duration6)
	fmt.Println("ParseDuration(\"24h\")", t6)

	duration7, _ := time.ParseDuration("1h")
	t7 := t.Add(duration7).AddDate(1,1,0)
	fmt.Println("ParseDuration(\"1h\").AddDae(1, 1, 0)", t7)

	fmt.Println("AddDate(1, 0, 0)", t.AddDate(1, 0, 0))
	fmt.Println("AddDate(1, 1, 0)", t.AddDate(1, 1, 0))
	fmt.Println("AddDate(1, 1, 1)", t.AddDate(1, 1, 1))



	fmt.Println()
	fmt.Println(t.Equal(time.Now()))
	fmt.Println(t.Before(time.Now()))
	fmt.Println(t.After(time.Now()))

}
/*
Package time: https://golang.google.cn/pkg/time

output:

Now() 2020-01-14 18:11:38.5806099 +0800 CST m=+0.002000701
Now() 2020-01-14 18:11:38.5806099 +0800 CST m=+0.002000701
Date() 2019-10-08 10:05:11 +0000 UTC

Now().Unix() 1578996698
Now().UnixNano() 1578996698593639700

Local() 2019-10-08 18:05:11 +0800 CST
Location() UTC

Year() 2019
Date() 2019 October 8
Date() format: 2019-10-8
Date() format2: 2019-10-08
Month() October
Weekday() Tuesday
ISOWeek() 2019 41
Day() 8
Clock() 10 5 11
Clock() format: 10-5-11
Clock() format2: 10-05-11
Hour() 10
Minute() 5
Second() 11

Format() 2019-10-08 10:05:11
Formant() err 8007-10-08 108:08:10

ParseDuration("-1m") 2019-10-08 10:04:11 +0000 UTC
ParseDuration("-1h") 2019-10-08 09:05:11 +0000 UTC
ParseDuration("-24h") 2019-10-07 10:05:11 +0000 UTC
ParseDuration("1m") 2019-10-08 10:06:11 +0000 UTC
ParseDuration("1h") 2019-10-08 11:05:11 +0000 UTC
ParseDuration("24h") 2019-10-09 10:05:11 +0000 UTC
ParseDuration("1h").AddDae(1, 1, 0) 2020-11-08 11:05:11 +0000 UTC
AddDate(1, 0, 0) 2020-10-08 10:05:11 +0000 UTC
AddDate(1, 1, 0) 2020-11-08 10:05:11 +0000 UTC
AddDate(1, 1, 1) 2020-11-09 10:05:11 +0000 UTC

false
true
false

*/
package time

import (
	"fmt"
	"time"
)

func TimeExample() {
	t := time.Now()

	// 日期格式
	fmt.Println("日期格式:", t.Format("2006-01-02 15:05:01"))
	fmt.Println("日期格式2:", t.Format("2006/01/02 15:05:01"))

	/*
		格式化时间使用  T.Format()
		其中2006-01-02 15:04:05是格式化的标准格式
		可以理解为占位符 年(2006) 月(01/1) 日(02/2) 时(15/3) 分(04/4) 秒(05/5)
		每个占位符中间可以使用符号比如"-","/"等，也可以什么都不加直接连接在一起
	*/
	fmt.Printf("%v\n", t.Format("2006-01-02 15:04:05"))
	fmt.Printf("%v\n", t.Format("2006-01-02"))
	fmt.Printf("%v\n", t.Format("2006-1-2")) // 占位符月日不带0，格式化后的小于10的月日也不带0
	fmt.Printf("%v\n", t.Format("01"))

	// 具体时间
	fmt.Println("日期年:", t.Year())
	fmt.Println("日期年:", t.Month())
	fmt.Println("月转数字格式:", int(t.Month()))
	fmt.Println("日期日", t.Day())
	fmt.Println("日期小时:", t.Hour())
	fmt.Println("日期分钟:", t.Minute())
	fmt.Println("日期秒:", t.Second())
	fmt.Println("日期周:", t.Weekday())
	fmt.Println("日期周数字:", int(t.Weekday()))

	// 一天后
	fmt.Printf("%v\n", time.Now().Add(time.Hour*24).Format("2006/01/02 15:03:04"))
	// 一天前
	fmt.Printf("%v\n", time.Now().Add(-time.Hour*24).Format("2006/01/02 15:03:04"))

	// 日期转时间戳
	Loc, _ := time.LoadLocation("Asia/Shanghai")
	t1, err := time.ParseInLocation("2006-01-02", "2022-11-03", Loc)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", t1.Unix())
	}

	// 时间: 代码测试性能
	TNanoStart := time.Now().UnixNano() //纳秒
	time.Sleep(2 * time.Second)
	TNanoEnd := time.Now().UnixNano()
	fmt.Println("时间间隔纳秒:", TNanoEnd-TNanoStart)

	// 定义时间格式: 日期转换
	temp := "2006-01-02 15:04:05"
	timestamp := time.Now().Unix()
	fmt.Println("当前时间戳:", timestamp)
	dates := time.Unix(timestamp, 0)
	fmt.Println("当前时间:", dates.Format(temp))

}

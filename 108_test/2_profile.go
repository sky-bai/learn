package main

import (
	"backend/lib/crontab"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

/**
 * ## 使用方法
 * Field name   | Mandatory? | Allowed values  | Allowed special characters
 * ----------   | ---------- | --------------  | --------------------------
 * Minutes      | Yes        | 0-59            | * / , -
 * Hours        | Yes        | 0-23            | * / , -
 * Day of month | Yes        | 1-31            | * / , - ?
 * Month        | Yes        | 1-12 or JAN-DEC | * / , -
 * Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
 *
 * 标准的cron是没有秒的，可以使用
 * @every 5s 来实现每5s，同理：@every 1h10m18s
 */

func main() {
	c := crontab.NewAndRun()

	fmt.Println("cron start...")

	// 每半个小时获取tcp应用的profile文件
	c.AddFunc("0 */30 * * * *", func() {
		GetProfile()
	})

	// 每半个小时获取tcp应用的profile文件
	c.AddFunc("@every 15s", func() {
		fmt.Println("每10s获取一次profile文件")
		GetProfile()
		fmt.Println("获取profile文件完成")
	})

	// 等待中断信号
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// 等待定时器完成任务
	ctx := c.Stop()
	<-ctx.Done()
	fmt.Println("cron stop...")

}

//

func GetProfile() {

	now := time.Now().Format("2006_01_02_15_04_05")
	fmt.Println("now:", now)

	data, err := exec.Command("wget", "http://127.0.0.1:6060/debug/pprof/profile").Output()
	if err != nil {
		fmt.Println(err)
	}
	// 保存为文件
	fileName := "profile" + now
	// 获取到当前目录位置
	dir, _ := os.Getwd()
	f, err := os.Create(dir + "/" + fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}

}

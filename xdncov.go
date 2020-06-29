package main

import (
	"math/rand"
	"time"

	"github.com/robfig/cron/v3"
)

type Job struct{}

// 定时任务计划
/*
- spec，传入 cron 时间设置
- job，对应执行的任务
*/
func StartJob(spec string, job Job) {
	c := cron.New()

	// 这里可以添加多个任务，多次 c.AddJob()
	c.AddJob(spec, job)

	// Run()启动执行任务
	// Start()会开启一个新的goroutine，如不设置time.Sleep()
	// 则该goroutine将被直接关闭
	c.Run()
	// 退出时关闭计划任务
	defer c.Stop()
}

// implement Run() interface to start rsync job
func (this Job) Run() {
	studentConfigSlice := CollectConfigs("./configs")

	for _, eachConfig := range studentConfigSlice {
		// 为可能的并发做准备
		// go signIn(eachConfig)

		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		SignIn(eachConfig)
	}
}

func main() {

	// 首次运行程序则立即执行
	studentConfigSlice := CollectConfigs("./configs")
	for _, eachConfig := range studentConfigSlice {
		// 为可能的并发做准备
		// go signIn(eachConfig)

		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		SignIn(eachConfig)
	}

	job1 := Job{}

	// 检查cron书写格式 https://crontab.guru/
	StartJob("3,7 0,12 * * *", job1)
}

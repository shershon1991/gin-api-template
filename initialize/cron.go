package initialize

import (
	"github.com/robfig/cron/v3"
	"shershon1991/gin-api-template/crontab"
	"shershon1991/gin-api-template/global"
)

// 定时任务管理
func initCron() {
	if !global.GvaConfig.Cron.Enable {
		return
	}
	c := cron.New(cron.WithSeconds())
	addJob(c)
	addFunc(c)
	c.Start()

}

// 添加Job任务
func addJob(c *cron.Cron) {
	_, _ = c.AddJob("@every 1s", crontab.DemoCron{})
}

// 添加Func任务
func addFunc(c *cron.Cron) {

}

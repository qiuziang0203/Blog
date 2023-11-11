package timedtask

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func Init() {
	// 新建一个定时任务对象,根据cron表达式进行时间调度,cron可以精确到秒,大部分表达式格式也是从秒开始
	// 默认从分开始进行时间调度
	// cronTab := cron.New()
	// 精确到秒
	cronTab := cron.New(cron.WithSeconds())
	// 定义定时器调用的任务函数
	task := func() {
		fmt.Println(time.Now())
		LikeBlog()
		LikeUser()
		FavorUser()
		FavorBlog()
	}
	// 定时任务,cron表达式,每3秒一次
	spec := "*/3 * * * * ?"
	// 添加定时任务
	cronTab.AddFunc(spec, task)
	// 启动定时器
	cronTab.Start()
}

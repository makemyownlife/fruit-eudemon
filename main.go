package main

import (
	"fruit-eudemon/meican"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	var ch chan int
	logrus.Info("hello, 新鲜水果开始抓取程序")
	//定时任务
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for range ticker.C {
			meican.DoTask()
		}
		ch <- 1
	}()
	<-ch
}

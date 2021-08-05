package main

import (
	"fruit-eudemon/meican"
	"fruit-eudemon/rlog"
	"time"
)

func main() {
	var ch chan int
	rlog.Info("hello, 新鲜水果开始抓取程序", nil)
	//定时任务
	ticker := time.NewTicker(time.Second * 15)
	go func() {
		for range ticker.C {
			meican.DoTask()
		}
		ch <- 1
	}()
	<-ch
}

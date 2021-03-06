package main

import (
	"bird/config"
	"bird/pkg/spider"
	"bird/pkg/useGPU"
	"bird/pkg/workPool"
	"fmt"
	"sync"
	"time"
)

type TypeFunc func(*config.Global)

// 实现多线程从任务池拿去任务，多线程爬虫下载图片，多线程任务调度获取GPU资源，保障GPU满负荷运行
func main() {
	var global = config.NewGlobal()
	new := time.Now()
	fmt.Println(new.Format("2006-01-02 15:04:05"))
	start := time.Now().UnixNano()
	var wg sync.WaitGroup
	// // 从任务池获取任务
	// go workPool.GetSpiderWork()
	// // 使用爬虫下载图片
	// go spider.GetWorkFromGSE()
	// // 使用gpu设备处理
	// go useGPU.GetOCRInfo()
	opts := []TypeFunc{workPool.GetSpiderWork, spider.GetWorkFromChan, useGPU.GetOCRInfo}
	for _, opt := range opts {
		wg.Add(1)
		go func(opt TypeFunc) {
			fmt.Println("启动线程")
			opt(global)
			wg.Done()
		}(opt)
	}
	wg.Wait()
	end := time.Now().UnixNano()
	fmt.Printf("cost is :%d \n", (end-start)/1000)
}

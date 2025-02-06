package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	go func() {
		if err := GetServer().Run(); err != nil {
			panic(err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个 5 秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的 Ctrl+C 就是触发系统 SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify 把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给 quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV) // 此处不会阻塞
	<-quit                                                                // 阻塞在此，当接收到上述两种信号时才会往下执行
	shutdownGraceful()
}

// shutdownGraceful 优雅关闭
// 这里主要是用来关闭自身的资源
func shutdownGraceful(fs ...func(chan struct{}) error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	var wg sync.WaitGroup
	for _, f := range fs {
		wg.Add(1)
		go func(f func(chan struct{}) error) {
			defer wg.Done()

			c := make(chan struct{}, 1)
			go f(c)
			select {
			case <-c:
			case <-ctx.Done():
			}
		}(f)
	}

	wg.Wait()
}

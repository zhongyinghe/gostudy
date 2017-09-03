package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "监控一")
	go watch(ctx, "监控二")
	go watch(ctx, "监控三")

	time.Sleep(time.Second * 10)
	fmt.Println("通知监控结束...")
	cancel()
	time.Sleep(time.Second * 5)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(time.Second * 2)
		}
	}
}

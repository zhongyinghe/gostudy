package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(time.Second * 2)
			}
		}
	}(ctx)

	time.Sleep(time.Second * 10)
	fmt.Println("通知监控停止...")
	cancel()
	time.Sleep(time.Second * 5)
}

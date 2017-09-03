package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	key := "name"
	val := "监控中心"
	ctx, cancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx, key, val)
	go watch(valueCtx, key)
	time.Sleep(time.Second * 10)
	fmt.Println("可以了, 通知监控停止...")
	cancel()
	time.Sleep(time.Second * 5)
}

func watch(ctx context.Context, key string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), "监控退出了，监控停止...")
			return
		default:
			fmt.Println(ctx.Value(key), "goroutine正在监控中...")
			time.Sleep(time.Second * 1)
		}
	}
}

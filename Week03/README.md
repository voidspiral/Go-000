## Q&A

基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出

课后补了signal包和errgoup包的用法

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctx)

	fmt.Println(os.Getgid())
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Duration(10) * time.Second)
		fmt.Fprintln(writer, "Hello World.")
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	//  添加服务启动函数
	group.Go(func() error {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Print("http server error: %v", err.Error())
		}
		return err
	})

	// 添加信号检测函数
	group.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		select {
		case <-sigs:
			return errors.New("notify kill signals.")
		}
	})
	//其他服务
	group.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("cancel")
				return ctx.Err()
			default:
				fmt.Println("do something")
				time.Sleep(1 * time.Second)
			}
		}
	})
	
	if err := group.Wait(); err != nil {
		fmt.Println("cancel request!")
		cancel()
		fmt.Println("begin shutdown server")
		server.Shutdown(ctx)
	}
}

```




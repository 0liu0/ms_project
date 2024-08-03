package common

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(r *gin.Engine, servName string, port string) {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	//保证下面的优雅启停
	go func() {
		log.Printf("web server [%s] running in %s \n", servName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal)
	//SIGINT 用户发送INTR字符(Ctrl+C)触发
	//SIGTERM 结束程序(可以被捕获、阻塞或忽略)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutting Down project web server [%s]...", servName)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("web server [%s] Shutdown, cause by : %val \n", servName, err.Error())
	}
	select {
	case <-ctx.Done():
		log.Println("关闭超时")
	}
	log.Printf("web server [%s] stop success... \n", servName)
}

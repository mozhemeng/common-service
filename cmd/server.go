package cmd

import (
	"common_service/global"
	"common_service/internal/routers"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title 通用服务
// @version 1.1
// @description 通用服务框架

// @securitydefinitions.apikey JWT
// @in header
// @name Authorization
func runServer() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	srv := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("srv.ListenAndServe error: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 等待5秒内server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}


var serverCmd = &cobra.Command{
	Use: "server",
	Aliases: []string{"api", "service", "run"},
	Short: "run api server",
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

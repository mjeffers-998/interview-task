package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/mjeffers-998/interview-task/controllers"
	"github.com/mjeffers-998/interview-task/routes"
	"github.com/mjeffers-998/interview-task/storage"
)

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Println("Exiting")
		os.Exit(0)
	}()
	db, err := storage.NewDB()
	if err != nil {
		panic(err)
	}
	controllers.SetController(db)
	r := gin.Default()
	routes.CreateRoutes(r)
	r.Run("localhost:9007")

}

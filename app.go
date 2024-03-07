package main

import (
	"context"
	"fmt"
	"os"
	"server/database"
	"server/handlers"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("config.env")

	GIN_MODE := os.Getenv("GIN_MODE")
	GIN_ADDR := os.Getenv("GIN_ADDR")

	gin.SetMode(GIN_MODE)

	db := database.NewDb()

	defer db.Conn.Close(context.Background())

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./server-vue/dist", false)))
	router.POST("/api/auth", handlers.CheckUserPassword(db))

	s, err := gocron.NewScheduler()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	j, err := s.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(
			func() {
				fmt.Println("Working...")
			},
		),
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(j.ID())

	s.Start()

	router.Run(GIN_ADDR)

	err = s.Shutdown()
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"github.com/gammazero/workerpool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Printf("-------------------start--------------------------------------------------------------------------")
	// if want sync, use only one worker
	// if concurrency, use > 1 worker
	wp := workerpool.New(1)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		for i := 0; i < 100; i++ {
			// always use new variable for one job, if you don't ware race conditions data
			a := i
			wp.Submit(func() {
				fmt.Println("Job : ", a, &a)
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//fmt.Printf("-------------------stopWait--------------------------------------------------------------------------")
	//wp.StopWait()
	//fmt.Printf("-------------------end--------------------------------------------------------------------------")
}

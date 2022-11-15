package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/alitto/pond"
)

func GroupWork(pool *pond.WorkerPool) {
	// Create a task group associated to a context
	group, ctx := pool.GroupContext(context.Background())

	var urls = []string{
		"https://www.golang.org/",
		"https://www.google.com/",
		"https://www.github.com/",
	}

	// Submit tasks to fetch each URL
	for _, url := range urls {
		url := url
		group.Submit(func() error {
			fmt.Println("start job")
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}

	// Wait for all HTTP requests to complete.
	err := group.Wait()
	if err != nil {
		fmt.Printf("Failed to fetch URLs: %v", err)
	} else {
		fmt.Println("Successfully fetched all URLs")
	}
}

func SimpleWork(pool *pond.WorkerPool) {
	// Submit 1000 tasks
	for i := 0; i < 1000; i++ {
		n := i
		pool.Submit(func() {
			fmt.Printf("Running task #%d\n", n)
		})
	}
}

func main() {
	// Create a worker pool
	pool := pond.New(1, 1000)
	defer pool.StopAndWait()

	r := gin.Default()
	r.GET("/groupWork", func(c *gin.Context) {
		go GroupWork(pool)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/simpleWork", func(c *gin.Context) {
		SimpleWork(pool)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

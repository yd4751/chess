package main

import (
	"fmt"
	"time"

	"github.com/gochenzl/chess/util/redis"
)

func main() {
	server := redis.NewServer(":6379", redis.NewMemoryStore())
	if err := server.Run(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("redis server start")
	for {
		time.Sleep(time.Second * 10)
	}
}

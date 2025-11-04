package main

import (
	"test_bi/mods"

	"github.com/gin-gonic/gin"
)

// var (
// 	cache *redis.Client
// )

// func initRedis() {
// 	cache = redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "",
// 		DB:       0,
// 	})
// 	_, err := cache.Ping(context.Background()).Result()
// 	if err != nil {
// 		log.Fatal("Redis连接失败:", err)
// 	}
// }

func main() {
	mods.InitDB()
	// initRedis()
	r := gin.Default()
	mods.SetupRoutes(r)
	r.Run(":8080")
}

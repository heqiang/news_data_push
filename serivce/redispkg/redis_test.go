package redispkg

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

func TestAdd(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	resInt, err := client.SAdd("news:md5Set", "test").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resInt)

}

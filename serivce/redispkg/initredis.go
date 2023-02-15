package redispkg

import (
	"data_push/config"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func InitRedis(redisConfig *config.RedisConfig) error {
	config.RedisDb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port), // redis地址
		Password: redisConfig.Password,                                     // redis密码，没有则留空
		DB:       redisConfig.Db,                                           // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err := config.RedisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// Add redis 集合中添加key
// 返回值true 说明不存在redis,否则存在
func Add(key string) bool {
	addInt, err := config.RedisDb.SAdd(config.Config.RedisConfig.SetKey, key).Result()
	if err != nil {
		zap.L().Error("redis 新闻指纹添加失败")
		return false
	}
	if addInt == 1 {
		return true
	}
	return false
}

package utils

import (
	"douban-webend/config"
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
	"time"
)

// 长连接
var redisClient *redis.Client

func ConnectRedis() {
	redisClient = makeAConnection(0)
	pong, err := redisClient.Ping().Result()

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("redis connect!")
	log.Println("redis ping result:", pong)
}

// makeAConnection 建立一个连接
// 参数
// 		- ttl 这段时间后关闭连接 ttl < 0 则不关闭连接
func makeAConnection(ttl time.Duration) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisAddrInner, // 使用内网
		Password: config.Config.RedisPassword,
		DB:       0,
	})
	if ttl > 0 {
		go func(t time.Duration, c *redis.Client) {
			<-time.NewTimer(t).C
			err := client.Close()
			if err != nil {
				log.Println("没有正常关闭Redis连接!")
				log.Panicln(err)
			}
			log.Printf("关闭连接 %v", client)
		}(ttl, client)
	}
	return client
}

// RedisSetString 写入字符串
// exp <= 0 说明没有过期时间
func RedisSetString(key, value string, exp time.Duration) error {
	err := redisClient.Set(key, value, exp).Err()
	if err != nil {
		return ServerInternalError
	}
	return nil
}

func RedisGetString(key string) (string, error) {
	result, err := redisClient.Get(key).Result()
	if err != nil {
		return "", ServerInternalError
	}
	return result, nil
}

func RedisSetStruct(key string, stc interface{}, exp time.Duration) error {
	b, err := json.Marshal(stc)
	if err != nil {
		return ServerInternalError
	}
	err = redisClient.Set(key, b, exp).Err()
	if err != nil {
		return ServerInternalError
	}
	return nil
}

// RedisGetStruct stc 结构体指针
func RedisGetStruct(key string, stc interface{}) error {
	result, err := redisClient.Get(key).Result()
	if err != nil {
		return ServerInternalError
	}
	err = json.Unmarshal([]byte(result), stc)
	if err != nil {
		return ServerInternalError
	}
	return nil
}

// TODO 添加更多

package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

//redis
//用作缓存系统、计数场景、list简单队列、排行榜
//支持RDB和AOF持久化，master/slave模式
//go get -u github.com/go-redis/redis

// 声明一个全局的rdb变量
var redisdb *redis.Client

// 初始化连接
func initClient() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 未设置密码
		DB:       0,  // 默认DB
	})
	_, err = redisdb.Ping().Result()
	if err != nil {
		return err
	}
	return
}

//zadd添加数据
func redisAdd() {
	zsetKey := "language_rank"
	languages := []*redis.Z{
		&redis.Z{Score: 91.0, Member: "Golang"},
		&redis.Z{Score: 98.0, Member: "Java"},
		&redis.Z{Score: 94.0, Member: "Python"},
	}
	// ZADD
	num, err := redisdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Println("添加失败", err)
		return
	}
	fmt.Println(num)

	// 把Golang加10
	newScore, err := redisdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Println("增分失败", err)
		return
	}
	fmt.Println(newScore)

	// 取分数最高的2种语言
	ret, err := redisdb.ZRevRangeWithScores(zsetKey, 0, 1).Result()
	if err != nil {
		fmt.Println("获取失败", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	//通过得分区间获取
	op := &redis.ZRangeBy{
		Min: "91",
		Max: "98",
	}
	ret, err = redisdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Println("获取失败", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func main() {
	err := initClient()
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	fmt.Println("连接成功")
	redisAdd()
}

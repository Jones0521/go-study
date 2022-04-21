package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client // 声明一个全局的redis连接对象
)

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.1.42:6379",
		Password: "PXEBA5ByKaURx", // Password
		// 下面为默认配置
		DB:       2,  // use default Db
		PoolSize: 10, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err = rdb.Ping(ctx).Result()
	return err
}

func demo1() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	//v, err := rdb.Set(ctx, "name", "wanderings", 50*time.Second).Result()
	// 设置值 类型的命令后面一般用 Err()
	err := rdb.Set(ctx, "name", "wanderings", 50*time.Second).Err()
	fmt.Println(err)
	// 获取值 类型的命令一般后面用 Result()
	v, err := rdb.Get(ctx, "name").Result()
	fmt.Println(v, err)
}

func redisExample2() {
	// key
	zsetKey := "language:rank"
	// value
	languages := []*redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// ZADD
	err := rdb.ZAdd(ctx, zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success")

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func main() {
	if err := initClient(); err != nil {
		fmt.Println("init redis clinet failed, err:", err)
		return
	}
	fmt.Println("连接redis成功")
	demo1()
}

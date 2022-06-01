package main

import "go.uber.org/zap"

// zap demo
// 要用zap 来记录日志, 就需要生成logger对象
func main() {
	// 获取Logger 对象
	zap.NewExample()
	zap.NewDevelopment()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	// 记录日志
	var uid int64 = 18923312
	isLogin := true
	name := "jade"
	data := []int{1, 2}
	logger.Info(
		"日志信息",
		zap.Int64("uid", uid),
		zap.Bool("isLogin", isLogin),
		zap.String("name", name),
		zap.Any("data", data),
		zap.Ints("data", data),
	)
}

package main

import (
	"fmt"
	"myWebCli/controller/run"
	"myWebCli/controller/validator"
	"myWebCli/db/mysql"
	"myWebCli/db/redis"
	"myWebCli/log"
	"myWebCli/routes"
	"myWebCli/settings"
	"myWebCli/utils/snowflake"

	"go.uber.org/zap"
)

func main() {
	// 1. 加载配置 viper config.yaml
	err := settings.Init("config.yaml")
	if err != nil {
		fmt.Println("加载配置失败！")
		return
	}
	fmt.Println("加载配置成功！")
	zap.L().Debug("加载配置成功！")
	defer zap.L().Sync()

	// 2. 加载日志组件 lumberjack zap库
	err = log.Init(settings.Conf.Log, settings.Conf.Mode)
	if err != nil {
		fmt.Println("初始化自定义日志失败！err=", err)
		return
	}
	fmt.Println("初始化日志成功！")
	zap.L().Debug("初始化日志成功")
	defer zap.L().Sync()

	// 3. 加载mysql  sqlx库
	err = mysql.Init(settings.Conf.MySQL)
	if err != nil {
		fmt.Println("初始化MySQL数据库失败！err=", err)
		zap.L().Debug("初始化MySQL数据库失败！")
		return
	}
	zap.L().Debug("初始化MySQL数据库成功！")
	defer mysql.Close()

	// 4. 加载redis
	err = redis.Init(settings.Conf.Redis)
	if err != nil {
		fmt.Println("初始化Redis数据库失败！err=", err)
		zap.L().Debug("初始化Redis数据库失败！")
		return
	}
	zap.L().Debug("初始化Redis数据库成功！")
	defer redis.Close()

	// 5. 加载雪花id算法
	err = snowflake.Init(settings.Conf.StartTime, int64(settings.Conf.MachineId))
	if err != nil {
		fmt.Println("初始化雪花id算法失败！err=", err)
		zap.L().Debug("初始化雪花id算法失败！")
		return
	}
	zap.L().Debug("初始化雪花id算法成功！")

	// 6. 初始化校验器
	// InitTrans初始化校验翻译器
	if err := validator.InitTrans("zh"); err != nil {
		fmt.Printf("init validator translator failed, err=%v\n", err)
	}
	// 5. 加载路由
	r := routes.SetUp(settings.Conf.Mode)

	// 6. 优雅启动
	run.ForeverElegant(r)
}

package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type AppConfig struct {
	*App   `mapstructure:"app"`
	*Log   `mapstructure:"log"`
	*MySQL `mapstructure:"mysql"`
	*Redis `mapstructure:"redis"`
}

type App struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineId int    `mapstructure:"machine_id"`
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
}

type Log struct {
	Level      string `mapstructure:"debug"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackUps int    `mapstructure:"max_backups"`
}

type MySQL struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	PassWord     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"maxopenconns"`
	MaxIdleConns int    `mapstructure:"maxidleconns"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
}

var Conf = new(AppConfig)

func Init(fileName string) (err error) {
	viper.SetConfigFile(fileName) // 配置文件名称
	viper.AddConfigPath(".")      // 配置文件相对于main.go文件的路径
	err = viper.ReadInConfig()    // 读取配置
	if err != nil {
		fmt.Println("读取配置错误！err=", err)
		return err
	}
	if err := viper.Unmarshal(Conf); err != nil { // 解析配置
		fmt.Println("解析配置错误！err=", err)
		return err
	}
	fmt.Println(Conf)
	viper.WatchConfig() // 监听配置文件热更改
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了！")
		if err := viper.Unmarshal(Conf); err != nil { // 解析配置
			fmt.Println("解析配置错误！err=", err)
		}
	})
	return
}

package common

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("获取项目路径error: %s \n", err))
	}
	viper.AddConfigPath(path)          //设置读取的文件路径
	viper.SetConfigName("application") //设置读取的文件名
	viper.SetConfigType("yaml")        //设置文件的类型
	//尝试进行配置读取
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置文件读取失败: %s \n", err))
	}
}

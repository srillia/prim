package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {

	runEnv := os.Getenv("RUN_ENV")

	switch runEnv {
	case "local":
		viper.SetConfigName("config/local")
	case "dev":
		viper.SetConfigName("config/dev")
	case "test":
		viper.SetConfigName("config/test")
	case "gray":
		viper.SetConfigName("config/gray")
	case "prod":
		viper.SetConfigName("config/prod")
	default:
		viper.SetConfigName("config/local")
	}
	viper.AddConfigPath("../") // 添加搜索路径

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config redis:", viper.Get("redis"))

}

package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"os"
	"prim/helper"
	"prim/lib/mongolib"
	"prim/lib/redislib"
	"prim/models"
)

var (
	configPath string

	ServerIp string
	RpcPort  string
	AccPort  string
	HttpPort string
)

func SetConfigPath(path string) {
	configPath = path
	return
}

func InitConfig() {

	runEnv := os.Getenv("RUN_ENV")

	//path := configPath

	switch runEnv {
	case "local":
		viper.SetConfigName("/local")
	case "dev":
		viper.SetConfigName("/dev")
	case "test":
		viper.SetConfigName("/test")
	case "gray":
		viper.SetConfigName("/gray")
	case "prod":
		viper.SetConfigName("/prod")
	default:
		viper.SetConfigName("/local")
	}

	if configPath == "" {
		dir, _ := os.Getwd()
		viper.AddConfigPath(dir + "/config") // 添加搜索路径
	} else {
		viper.AddConfigPath(configPath + "/config") // 添加搜索路径
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config redis:", viper.Get("redis"))

	initRpcServerConfig()

}

func initRpcServerConfig() {
	ServerIp = helper.GetServerIp()
	RpcPort = viper.GetString("app.rpcPort")
	AccPort = viper.GetString("app.webSocketPort")
	HttpPort = viper.GetString("app.httpPort")
}

func GetServer() (server *models.Server) {
	server = models.NewServer(ServerIp, RpcPort)

	return
}

func IsLocal(server *models.Server) (isLocal bool) {
	if server.Ip == ServerIp && server.Port == RpcPort {
		isLocal = true
	}

	return
}

// 初始化日志
func InitFile() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	logFile := viper.GetString("app.logFile")
	f, _ := os.Create(logFile)
	gin.DefaultWriter = io.MultiWriter(f)
}

func InitRedis() {
	redislib.NewClient()
}

func InitMongo() {
	mongolib.NewClient()
}

package main

import (
	"cloudPlatformDemo/middleware"
	"cloudPlatformDemo/routers"
	"cloudPlatformDemo/utils"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	//"net"
	"os"
	"path"
	"path/filepath"
	//"time"
)

var ENV = "dev"

func initConfig() map[string]*map[string]string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		middleware.Log.Error(err)
		os.Exit(1)
	}

	config, err := ini.Load(path.Join(dir, "../config/server.ini"))
	if err != nil {
		middleware.Log.Error(err)
		os.Exit(1)
	}
	sqlConfig, err := ini.Load(path.Join(dir, "../config/database.ini"))
	if err != nil {
		middleware.Log.Error(err)
		os.Exit(1)
	}

	configMap := make(map[string]*map[string]string)
	httpConfigMap := make(map[string]string, 5)
	tcpConfigMap := make(map[string]string, 5)
	mysqlConfigMap := make(map[string]string, 10)

	httpConfigMap["ip"] = config.Section(ENV).Key("http.ip").String()
	httpConfigMap["port"] = config.Section(ENV).Key("http.port").String()

	tcpConfigMap["ip"] = config.Section(ENV).Key("tcp.ip").String()
	tcpConfigMap["port"] = config.Section(ENV).Key("tcp.port").String()

	mysqlConfigMap["driver"] = sqlConfig.Section(ENV).Key("ccsResource.driver").String()
	mysqlConfigMap["host"] = sqlConfig.Section(ENV).Key("ccsResource.host").String()
	mysqlConfigMap["port"] = sqlConfig.Section(ENV).Key("ccsResource.port").String()
	mysqlConfigMap["database"] = sqlConfig.Section(ENV).Key("ccsResource.database").String()
	mysqlConfigMap["username"] = sqlConfig.Section(ENV).Key("ccsResource.username").String()
	mysqlConfigMap["password"] = sqlConfig.Section(ENV).Key("ccsResource.password").String()
	mysqlConfigMap["charset"] = sqlConfig.Section(ENV).Key("ccsResource.charset").String()

	configMap["http"] = &httpConfigMap
	configMap["ip"] = &tcpConfigMap
	configMap["mysql"] = &mysqlConfigMap

	return configMap
}

func httpServer(httpConfig map[string]string) {
	utils.Include(routers.ResourceRouters)//注册每个app的路由
	router := utils.Init()
	address := httpConfig["ip"] + ":" + httpConfig["port"]
	if err := router.Run(address); err != nil {
		middleware.Log.Error("http server open failed\n", err)
	}
}

func main() {
	//读取配置文件
	config := initConfig()
	//初始化数据库引擎
	utils.GetDb(*config["mysql"])
	//创建HTTP服务器
	httpServer(*config["http"])
	//创建TCP服务器
	/*address := tcpConfig["ip"] + ":" + tcpConfig["port"]
	log.Info("tcp server address:", address)
	listenFd, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("tcp server open failed", err)
		os.Exit(1)
	}
	defer listenFd.Close()

	for {
		acceptFd, err := listenFd.Accept()
		if err != nil {
			continue
		}
		acceptFd.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
		go handleTcpRequest()
	}
	*/
	//创建HTTP客户端
}
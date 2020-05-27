package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"time"
)

var Log = NewLogger()

func NewLogger() *logrus.Logger{
	now := time.Now()
	logFilePath := ""
	if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
		logFilePath = path.Join(dir, "../log")
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	//创建日志文件
	logFileName := now.Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	//打开日志文件，用于输出日志
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err:", err)
	}

	logger := logrus.New()
	logger.Out = src//设置日志输出文件地址
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{})

	return logger
}

func LoggerToFile() gin.HandlerFunc {
	logger := NewLogger()
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()//走中间件
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)//获取执行时间
		reqMethod := c.Request.Method //请求方式
		reqUri := c.Request.RequestURI//请求路由
		statusCode := c.Writer.Status() //响应状态码
		clientIp := c.ClientIP() //请求IP

		logger.Infof("| %3d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIp, reqMethod, reqUri)
	}
}

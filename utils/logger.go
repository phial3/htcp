package utils

import (
	"os"
)
import (
	"github.com/sirupsen/logrus"
)

// logrus提供了New()函数来创建一个logrus的实例。
// 项目中，可以创建任意数量的logrus实例。
var Logger = logrus.New()

func init() {
	// 设置日志格式为json格式
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetReportCaller(true)

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	Logger.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	Logger.SetLevel(logrus.DebugLevel)
}

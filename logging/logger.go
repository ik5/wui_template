package logging

import (
	"log/syslog"

	"github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
	"github.com/spf13/viper"
)

// Logger is the system logger
var Logger *logrus.Logger

// InitLog initialize logging facility
func InitLog(socketType, address, tag string, priority syslog.Priority) {
	Logger = logrus.New()
	Logger.Formatter = &logrus.JSONFormatter{PrettyPrint: false}
	logLevel, err := logrus.ParseLevel(viper.GetString("log_level"))
	if err != nil {
		Logger.Level = logrus.TraceLevel
	} else {
		Logger.Level = logLevel
	}
	Logger.SetReportCaller(true)

	if !viper.GetBool("use_syslog") {
		return
	}
	hook, err := logrus_syslog.NewSyslogHook(socketType, address, priority, tag)
	if err != nil {
		Logger.Errorf("Unable to use syslog: %s", err)
	} else {
		Logger.AddHook(hook)
	}
}

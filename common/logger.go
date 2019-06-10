package common

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

//use logrus hook rotatelogs
var log = Log

//Log ...
var Log = logrus.New()

//InitLogger ...
func initLogger(l *LogConfigure) {
	baseLogPath := path.Join(l.Logdir,
		l.LogFileName)
	writer, err := rotatelogs.New(
		baseLogPath+"-%Y%m%d%H.log",
		rotatelogs.WithLinkName(baseLogPath),                                 // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Duration(l.MaxSaveHour)*time.Hour),        // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(l.RotationHour)*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		Log.Errorf("config local file system logger error. %v", errors.WithStack(err))
	}

	switch level := l.LogLevel; level {
	/*
		如果日志级别不是debug就不要打印日志到控制台了
	*/

	case "debug":
		Log.SetLevel(logrus.DebugLevel)
		Log.SetOutput(os.Stderr)
	case "verbose":
		Log.SetLevel(logrus.DebugLevel)
	case "info":
		setNull()
		Log.SetLevel(logrus.InfoLevel)
	case "warn":
		setNull()
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		setNull()
		Log.SetLevel(logrus.ErrorLevel)
	default:
		setNull()
		Log.SetLevel(logrus.InfoLevel)
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})

	Log.AddHook(lfHook)
}

func setNull() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	writer := bufio.NewWriter(src)
	Log.SetOutput(writer)
}

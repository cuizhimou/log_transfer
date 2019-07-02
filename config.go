package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

type LogConfig struct {
	KafkaAddr  string
	EsAddr     string
	LogPath    string
	Loglevel   string
	KafkaTopic string
}

var (
	logConfig *LogConfig
)

func initConfig(confType string, filename string) (err error) {
	//初始化config对象
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}
	logConfig = &LogConfig{}

	logConfig.Loglevel = conf.String("logs::log_level")
	if len(logConfig.Loglevel) == 0 {
		logConfig.Loglevel = "debug"
	}

	logConfig.LogPath = conf.String("logs::log_path")
	if len(logConfig.LogPath) == 0 {
		logConfig.LogPath = "./logs"
	}

	logConfig.KafkaAddr = conf.String("kafka::server_addr")
	if len(logConfig.KafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka address ")
		return
	}

	logConfig.EsAddr = conf.String("es::addr")
	if len(logConfig.EsAddr) == 0 {
		err = fmt.Errorf("invalid es address ")
		return
	}

	logConfig.KafkaTopic = conf.String("kafka::topic")
	if len(logConfig.KafkaTopic) == 0 {
		err = fmt.Errorf("invalid kafka topic ")
		return
	}
	return
}

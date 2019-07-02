package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {
	err := initConfig("ini", "./conf/log_transfer.conf")
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(logConfig)

	err = initLogger(logConfig.LogPath, logConfig.Loglevel)
	if err != nil {
		panic(err)
		return
	}
	logs.Debug("init logger succ")
	err = initKafka(logConfig.KafkaAddr, logConfig.KafkaTopic)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}

	logs.Debug("init kafka succ")

	err = initEs(logConfig.EsAddr)
	if err != nil {
		logs.Error("init es failed, err:%v", err)
		return
	}
	logs.Debug("init es succ")

	err = run()
	if err != nil {
		logs.Error("run failed, err:%v", err)
		return
	}

	logs.Warn("warning, log_transfer is exited")
}

package main

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"strings"
	"sync"
)

//var (
//	wg sync.WaitGroup
//)

type KafkaClient struct {
	client sarama.Consumer
	add    string
	topic  string
	wg 	   sync.WaitGroup
}

var (
	kafkaClient *KafkaClient
)

func initKafka(add string, topic string) (err error) {
	kafkaClient = &KafkaClient{}
	consumer, err := sarama.NewConsumer(strings.Split(add, ","), nil)
	if err != nil {
		logs.Error("failed to start consumer: %s", err)
		return
	}

	kafkaClient.client = consumer
	kafkaClient.add = add
	kafkaClient.topic = topic
	return
	//
	//partitionList, err := consumer.Partitions(topic)
	//if err != nil {
	//	logs.Error("failed to get the list of partitions: ", err)
	//	return
	//}
	//
	//
	//for partition := range partitionList {
	//	pc, errRet := consumer.ConsumePartition("nginx_log", int32(partition), sarama.OffsetNewest)
	//	if errRet != nil {
	//		err = errRet
	//		logs.Error("fail to start consumer for partition %d:%s\n", partition, err)
	//		return
	//	}
	//	defer pc.AsyncClose()
	//	go func(sarama.PartitionConsumer) {
	//		for msg := range pc.Messages() {
	//			logs.Debug("Parttion:%d, Offset:%d, key:%s, value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	//			//fmt.Println()
	//		}
	//		//wg.Done()
	//	}(pc)
	//}
	////等待线程退出
	////wg.Wait()
	////time.Sleep(time.Hour)
	////consumer.Close()
	//return
}

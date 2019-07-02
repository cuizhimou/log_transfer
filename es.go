package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/olivere/elastic"
	"context"
)

type LogMessage struct {
	App string
	Topic string
	Message string
}

var(
	esClient *elastic.Client
)

func initEs(addr string)(err error) {
	//ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(addr))
	if err != nil {
		logs.Debug("connect es error", err)
		return
	}
	esClient=client
	return
	//fmt.Println("conn es succ")
	//
	//tweet :=LogMessage{App:"olivere",Message:"Take Five"}
	//_,err =client.Index().Index("teitter").Type("tweet").Id("1").BodyJson(tweet).Do(ctx)
	//if err !=nil{
	//	panic(err)
	//	return
	//}
	//
	//fmt.Println("insert succ")
}

func sendToEs(topic string,data []byte)(err error)  {
	ctx := context.Background()
	msg :=&LogMessage{}
	msg.Topic=topic
	msg.Message=string(data)
	_,err =esClient.Index().Index(topic).Type(topic).BodyJson(msg).Do(ctx)
	if err !=nil{
		panic(err)
		return
	}
	return
}
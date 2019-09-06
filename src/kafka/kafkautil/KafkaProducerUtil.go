package kafkautil

//
// Copyright (c) 2018-2028 Corp-ci All Rights Reserved
// Package: kafkautil
// Version: 1.0
//
// Created by Smile-Sun on 2019/9/3 16:02
//
import (
	"github.com/Shopify/sarama"
	_ "github.com/bsm/sarama-cluster"
	"log"
	"time"
)

//
// Send data asynchronously to the Kafka message pair column, where Kafka's topic is a custom topic
//
func AsynchronousKafkaProducer(msg string, topic string, address ...string) {

	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Producer.Timeout = 5 * time.Second
	kafkaClient, err := sarama.NewClient(address, kafkaConfig)
	if err != nil {
		log.Panicf("sarama.NewClient [%s]", err)
	}
	asyncProducer, err := sarama.NewAsyncProducerFromClient(kafkaClient)
	if err != nil {
		log.Panicf("sarama.NewAsyncProducerFromClient [%s]", err)
	}
	defer asyncProducer.AsyncClose()

	go func(producer sarama.AsyncProducer) {
		for {
			select {
			case err := <-producer.Errors():
				if err != nil {
					log.Panicf("producer.Errors [%s]", err)
				}
			case <-producer.Successes():
				//value, _ := suc.Value.Encode()
				//fmt.Println("offsetCfg:", suc.Offset, " partitions:", suc.Partition, " metadata:", suc.Metadata, " value:", string(value))
			}
		}
	}(asyncProducer)
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
	}
	asyncProducer.Input() <- message
	time.Sleep(1 * time.Second)

}

package main

import (
	_ "github.com/bsm/sarama-cluster"
	"kafka/kafkautil"
)

/**
 * Copyright (c) 2018-2028 Corp-ci All Rights Reserved
 *
 * Project: corp-go
 * Package: main
 * Version: 1.0
 *
 * Created by Smile-Sun on 2019/9/3 15:00
 */
func main() {
	//kafkautil.AsynchronousKafkaProducer("src/config/kafka-server.properties" , "hadoop101:9092")
	kafkautil.AsynchronousKafkaProducer(" this is my third kafka producer ... ", "flink_topic", "hadoop101:9092")
}

package main

import (
	"file/fileutil"
	"fmt"
)

/**
 * Copyright (c) 2018-2028 Corp-ci All Rights Reserved
 *
 * Project: corp-go
 * Package: main
 * Version: 1.0
 *
 * Created by Smile-Sun on 2019/9/3 15:02
 */

func main() {
	config := fileutil.InitConfig("src/config/kafka-server.properties")
	fmt.Println(config["kafka.brokers"])
}

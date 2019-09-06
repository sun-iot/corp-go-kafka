package fileutil

//
// Copyright (c) 2018-2028 Corp-ci All Rights Reserved
//
// Project:
// Package: fileutil
// Version: 1.0
//
// Created by Smile-Sun on 2019/9/3 15:17
//
import (
	"bufio"
	"io"
	"os"
	"strings"
)

//
// Pass the path of the configuration file to get the KV key-value pair of the configuration file.
//
func InitConfig(path string) map[string]string {
	config := make(map[string]string)
	// open the file
	file, e := os.Open(path)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	// Create an output stream to the buffer stream of the file
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		str := strings.TrimSpace(string(line))
		index := strings.Index(str, "=")
		if index <= 0 {
			continue
		}
		// Take the key to the left of the equal sign
		key := strings.TrimSpace(str[:index])
		if len(key) <= 0 {
			continue
		}
		// Take the value to the right of the equal sign
		value := strings.TrimSpace(str[index+1:])
		if len(value) <= 0 {
			continue
		}
		config[key] = value
	}
	return config
}

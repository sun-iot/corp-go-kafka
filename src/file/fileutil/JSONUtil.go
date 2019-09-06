package fileutil

//
// Copyright (c) 2018-2028 Corp-ci All Rights Reserved
//
// Package: kafkautil
// Version: 1.0
//
// Created by Smile-Sun on 2019/9/4 13:33
//

import (
	"encoding/json"
	"errors"
	"fmt"
)

//
// The incoming JSON string is matched according to the field information queried by the database. If the matching is
// passed, the correct code is returned, otherwise an error is returned.
//

func CheckOutJSON(targetJSON []byte, reference map[string]string) (str string, err error) {
	var tempJSON interface{}
	err = nil
	// str = "{\"state\":1,\"msg\":\"succeed\"}"
	json.Unmarshal(targetJSON, &tempJSON)
	mapJSON := tempJSON.(map[string]interface{})
	for k, _ := range mapJSON {
		values := mapJSON[k]
		valueInterface := values.([]interface{})
		for _, vv := range valueInterface {
			vvMap := vv.(map[string]interface{})
			for _k, _v := range vvMap {
				switch _v.(type) {
				case string:
					if reference[_k] != "string" {
						str = fmt.Sprintf("{\"state\":0,\"data\": { \"%s\" : {\"get\" :\"%s\" , \"need\":\"%s\"} } ,\"msg\":\"failed\"}", _k, "string", reference[_k])
						err = errors.New(" string type match error ")
					}
					break
				case int:
					break
				case bool:
					if reference[_k] != "bool" {
						str = fmt.Sprintf("{\"state\":0,\"data\": { \"%s\" : {\"get\" :\"%s\" , \"need\":\"%s\"} } ,\"msg\":\"failed\"}", _k, "bool", reference[_k])
						err = errors.New(" bool type match error ")
					}
					break
				case int64:
				case int32:
				case float64:
					if reference[_k] != "long" && reference[_k] != "int" && reference[_k] != "float" {
						str = fmt.Sprintf("{\"state\":0,\"data\": { \"%s\" : {\"get\" :\"%s\" , \"need\":\"%s\"} } ,\"msg\":\"failed\"}", _k, "long", reference[_k])
						err = errors.New(" long/int/float type match error ")
					}
				case []interface{}:
					break
				default:
					str = "{\"state\":0,\"data\": \"null\" ,\"msg\":\"failed\"}"
					err = errors.New(" no type found ")
					break
				}
			}
		}
	}
	if len(str) <= 0 {
		str = "{\"state\":1,\"data\": \"ok\" ,\"msg\":\"failed\"}"
	}
	return str, err
}

package mysqlutil

//
// Copyright (c) 2018-2028 Corp-ci All Rights Reserved
//
// Package: kafkautil
// Version: 1.0
//
// Created by Smile-Sun on 2019/9/3 17:40
//

import (
	"database/sql"
	"file/fileutil"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

//
// Returns a connection to a database, but this connection is not secure. In a concurrent environment, such a connection
// can cause our database to crash.
//
// We set up 2048 of the largest database connections and 1024 of the largest data idle connections.
//
//
//

func GetDB() *sql.DB {
	mysqlConfig := fileutil.InitConfig("src/config/mysql-server.properties")
	db, e := sql.Open(mysqlConfig["mysql.driver"], mysqlConfig["mysql.user"]+":"+
		mysqlConfig["mysql.password"]+"@tcp("+mysqlConfig["mysql.url"]+":"+mysqlConfig["mysql.port"]+")/"+
		mysqlConfig["mysql.db"]+"?charset=utf8")
	if e != nil {
		log.Fatalf("sql.Open : [%s]", e)
	} else {
		db.SetMaxOpenConns(2048)
		db.SetMaxIdleConns(1024)
		db.Ping()
	}
	return db
}

// According to the incoming sql query statement, the query result is automatically returned as a map object is not
// repeated
//

func GetContent(sqlTest string) map[string]string {
	db := GetDB()
	rows, e := db.Query(sqlTest)
	if e != nil {
		log.Panicf("GetContent:db.Query [%s]", e)
	}
	return GetNameType(rows)
}

// According to the incoming sql query statement, the query result is automatically returned as a map object is not
// repeated

func GetNameType(rows *sql.Rows) map[string]string {
	config := make(map[string]string)
	for rows.Next() {
		var _name string
		var _type string
		rows.Scan(&_name, &_type)
		config[_name] = strings.TrimLeft(strings.TrimLeft(_type, "ci"), "_")
	}
	return config
}

// Customized method according to requirements.

func GetNameTypeTableId(rows *sql.Rows) (map[string]string, float64) {
	config := make(map[string]string)
	var tableid float64
	for rows.Next() {
		var _name string
		var _type string
		rows.Scan(&_name, &_type, &tableid)
		config[_name] = strings.TrimLeft(strings.TrimLeft(_type, "ci"), "_")
	}
	return config, tableid
}

// Customized method according to requirements.

func GetResutlWithTableId(sqlText string, args ...interface{}) (map[string]string, float64) {
	db := GetDB()
	rows, e := db.Query(sqlText, args[0], args[1], args[2], args[3])
	if e != nil {
		log.Panicf("stmt.Query [%s]", e)
	}
	db.Close()
	return GetNameTypeTableId(rows)
}

//
// The following method is used for testing. It is not recommended.
//

// Get one or more rows based on the incoming sql query

func GetRowsQuery(sql string) *sql.Rows {
	db := GetDB()
	rows, e := db.Query(sql)
	if e != nil {
		log.Panicf("GetRowsQuery:db.Query [%s]", e)
	}
	db.Close()
	return rows
}

// According to the incoming sql query statement and formal parameters, and convert the result into a map object return.

func GetResutl(sqlText string, args ...interface{}) map[string]string {
	db := GetDB()
	rows, e := db.Query(sqlText, args[0], args[1], args[2], args[3])
	if e != nil {
		log.Panicf("stmt.Query [%s]", e)
	}
	return GetNameType(rows)
}

// Perform analysis and processing on rows, return to map.

func GetRows(rows *sql.Rows) map[string]string {
	// 初始化
	config := make(map[string]string)
	// 得到每一列的列名
	columns, err := rows.Columns()
	if err != nil {
		log.Panicf("rows.Columns [%s]", err)
	}
	// 创建一个数组切片，用来临时存储我们的取到的值
	values := make([]sql.RawBytes, len(columns))
	// 真实参与扫描后得到的值
	scanArgs := make([]interface{}, len(values))
	// 把所有需要扫描的得到值得项拿到手
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		e := rows.Scan(scanArgs...)
		if e != nil {
			panic(err.Error())
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ":", value)
			config[columns[i]] = value
		}
	}
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
	return config
}

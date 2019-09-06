package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mysql/mysqlutil"
)

/**
 * Copyright (c) 2018-2028 Corp-ci All Rights Reserved
 *
 * Project:
 * Package: main
 * Version: 1.0
 *
 * Created by Smile-Sun on 2019/9/3 16:17
 */

func main() {

	sqlText := `
		SELECT ptf.name,ptf.type,ptf.datatbl_id
		FROM
		platform_table_fields AS ptf
		INNER JOIN 
		platform_data_tables AS pdt
		ON 
		ptf.datatbl_id=pdt.id
		INNER JOIN 
		platform_dataware_houses AS pdh
		ON
		pdt.dwh_id=pdh.id
		WHERE
		pdh.appkey= ?
		AND 
		pdh.channel= ?
		AND
		pdh.name=  ?
		AND
		pdt.name= ?
	`
	sqlResult, tableId := mysqlutil.GetResutlWithTableId(sqlText, "corp-ci-1", 1, "corp-ci-1", "corp-ci-1")
	fmt.Println(sqlResult["is"], tableId)
}

package urlutil

import (
	"file/fileutil"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"kafka/kafkautil"
	"log"
	"mysql/mysqlutil"
	"net/http"
	"strconv"
	"strings"
)

func HttpRequest(web http.ResponseWriter, req *http.Request) {
	// 解析参数，默认是不会解析的
	req.ParseForm()
	values := req.Form
	urlHeader := make(map[string]string)
	for k, v := range values {
		urlHeader[k] = strings.Join(v, "")
	}
	BodyJSON(req, urlHeader, web)
}

/**
对request中的body进行解析
*/
func BodyJSON(req *http.Request, urlMap map[string]string, web http.ResponseWriter) {
	byteJSON, e := ioutil.ReadAll(req.Body)
	if e != nil {
		log.Printf("ioutil.ReadAll [%s]", e)
	}

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
	//
	appkey := strings.Trim(urlMap["appkey"], "$")
	channel := strings.Trim(urlMap["channel"], "$")
	dataware := strings.Trim(urlMap["dataware"], "$")
	table := strings.Trim(urlMap["table"], "$")
	sqlResult, tableId := mysqlutil.GetResutlWithTableId(sqlText, appkey, channel, dataware, table)
	str, e := fileutil.CheckOutJSON(byteJSON, sqlResult)
	if e != nil {
		fmt.Fprintf(web, str)
		log.Printf("fileutil.CheckOutJSON [%s]", e)
	} else {
		// Splicing the topics to be sent to Kafka
		topic := fmt.Sprintf("local_%s_%s_%s", appkey, channel, strconv.FormatFloat(tableId, 'f', 0, 64))
		kafkautil.AsynchronousKafkaProducer(string(byteJSON), topic, "hadoop101:9092")
		fmt.Fprintf(web, str)
	}

}

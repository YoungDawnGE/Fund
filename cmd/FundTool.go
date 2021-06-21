package cmd

import (
	"fmt"
	"net/http"
	"time"
)

//"http://fund.10jqka.com.cn/001410/json/jsondwjz.json"

/**
下载fund的数据
 */
func DownLoadFundJson(code string)  {
	uri := "http://fund.10jqka.com.cn/"+code+"/json/jsondwjz.json"


	//Header := map[string][]string{
	//			"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"},
	//}
	//request := http.Request{
	//	Header:     Header,
	//	Method:     "GET",
	//	RequestURI: uri,
	//}


	client := &http.Client{Timeout: 5 * time.Second}
	response, _ := client.Get(uri)
	fmt.Println(response)
}

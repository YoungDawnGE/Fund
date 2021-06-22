package cmd

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

//"http://fund.10jqka.com.cn/001410/json/jsondwjz.json"
//下载fund的数据
func DownLoadFundTxt(code string) {
	//常量定义
	uri := "http://fund.10jqka.com.cn/" + code + "/json/jsondwjz.json"
	destFilename := "./raw_data/" + code + ".txt"
	bufSize := 10 * 1024

	//发送请求
	client := &http.Client{Timeout: 5 * time.Second}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	resp, _ := client.Do(req)

	//处理body
	fundRawBody := resp.Body
	reader := bufio.NewReaderSize(fundRawBody, bufSize)
	defer fundRawBody.Close()
	file, err := os.Create(destFilename)
	if err != nil {
		log.Println("create file failed :", destFilename)
		panic(err)
	}
	writer := bufio.NewWriter(file)

	//从body复制数据
	c, err := io.Copy(writer, reader)
	if err != nil {
		log.Println("copy file error :", err)
		return
	}
	log.Println("download", code, "ok :", c, "chars")
}

//将fund数据转化为.json格式
func GenFundJson() {

}

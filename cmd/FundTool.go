package cmd

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//"http://fund.10jqka.com.cn/001410/json/jsondwjz.json"
//下载fund的数据为txt格式
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
		log.Println("create file txt failed :", destFilename)
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	//从body复制数据
	c, err := io.Copy(writer, reader)
	if err != nil {
		log.Println("copy file error :", err)
		return
	}
	log.Println("download", code, "ok :", c, "chars")
}

//将下载好的fund.txt数据转化为.json格式
func GenFundJson(code string) {
	//变量名
	destFilename := "./json_data/" + code + ".json"

	//读取文件
	data, err := ioutil.ReadFile("./raw_data/" + code + ".txt")
	if err != nil {
		log.Println("File reading error", err)
		return
	}

	//var dwjz_006113=[["20180808","1.0000"]
	//取等于号的位置
	equalIndex := strings.Index(string(data), "=")
	//取data的[]byte
	data = data[equalIndex+1:]

	//创建./json_data/code.json
	file, err := os.Create(destFilename)
	if err != nil {
		log.Println("create file json failed :", destFilename)
		panic(err)
	}
	defer file.Close()

	//写入文件
	writer := bufio.NewWriter(file)
	if _, err = writer.Write(data); err != nil {
		log.Println("write to file json " + destFilename + " failed")
		return
	}
	defer writer.Flush()
	log.Println("write to file json " + destFilename + " success")
}

//读取文件 把json文件转化为日期数组和value数组
func JsonToDataArray(code string) ([]string, []float64) {
	sourceFilename := "./json_data/" + code + ".json"
	file, err := os.Open(sourceFilename)
	if err != nil {
		log.Println("open json file " + sourceFilename + " failed")
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("File reading error", err)
		panic(err)
	}

	tempArr := make([][]string, 10)
	_ = json.Unmarshal(data, &tempArr)

	//创建数组
	date := make([]string, len(tempArr))       //日期
	dataValue := make([]float64, len(tempArr)) //数据

	for i, item := range tempArr {
		date[i] = item[0]
		dataValue[i], _ = strconv.ParseFloat(item[1], 64)
	}
	log.Println("Generate date&data array success")
	return date, dataValue
}

//计算第N天第涨幅，今日涨幅=(今日净值-昨日净值)/昨日净值
func JsonToRateArray(code string) ([]string, []float64) {
	date, data := JsonToDataArray(code)
	rate := make([]float64, len(data))
	rate[0] = 0 //第一台你的

	for i := 1; i < len(data); i++ {
		rate[i] = (data[i] - data[i-1]) / data[i-1]
	}
	return date, rate
}

//计算N天的涨幅
func JsonToRateArrayNDay(code string, n int) ([]string, []float64) {
	date, data := JsonToDataArray(code)
	rate := make([]float64, len(data))
	for i := 0; i < n; i++ {
		rate[i] = 0
	}

	for i := n; i < len(data); i++ {
		rate[i] = (data[i] - data[i-n]) / data[i-n]
	}
	return date, rate
}

//判断是不是fund的码
func IsCode(code string) bool {
	isCode, _ := regexp.MatchString("\\b\\d{6}\\b", code)
	return isCode
}

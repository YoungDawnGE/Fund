package main

import (
	"fmt"
	"gyc.com/Fund/cmd"
)

func main() {
	//cmd.DownLoadFundTxt("008087")
	//cmd.GenFundJson("006113")
	_, data := cmd.JsonToDataArray("006113")
	fmt.Println(data)
}

package main

import "fmt"
import "deltatch/common"
import "gostudy"
import "gostudy/amusing"

func main() {

	//basestudy()
	//primeNumber()
	//utf8Demo()

	gostudy.GoroutineEntry()

	mat := new(amusing.MatrixConceal)
	mat.Init()
	//fmt.Printf("矩阵藏宝图是：%v\r\n", mat.Result)
}

func utf8Demo() {
	text := "hello,GO语言。"
	buffer := common.TextToUtf8Bytes(text) //make([]byte, 0)

	fmt.Printf("字符串UTF8编码形式（类转码）：%v\r\n", buffer)
	buffer = []byte(text)
	fmt.Printf("字符串UTF8编码形式（硬转码）：%v\r\n", buffer)

	text = common.BytesToUft8Text(buffer)
	fmt.Printf("字符串还原后：%v\r\n", text)
}

func primeNumber() {
	limit := 1000
	pnums := amusing.PrimeNumber(limit)
	fmt.Printf("%v以内的所有素数：%v", limit, pnums)
}

func basestudy() {
	fmt.Println("Hello World,您好，世界!")
	gostudy.Variable()
	gostudy.CatchExceptioin()

	line := common.ReadLine("请输入一行字符：")
	fmt.Println(line)
	line = common.ReadLine("即将终止运行...")
}

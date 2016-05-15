package main

import "fmt"
import "deltatch/common"

import "deltatch/cnregions"
import "gostudy"
import "gostudy/amusing"

func main() {

	//basestudy()
	//primeNumber()
	//utf8Demo()

	provider := cnregions.NewNbsDsProvider()
	provs, err := provider.GetVillage("360721104")
	if err != nil {
		fmt.Printf("爬取失败，详细：%+v", err)
		return
	}
	fmt.Println("从国家统计局获取的省份数据：\r\n")
	for _, n := range provs {
		fmt.Printf("%+v\r\n", *n)
	}

	//gostudy.GoroutineEntry()

	//mat := new(amusing.MatrixConceal)
	//mat.Init()
	//fmt.Printf("矩阵藏宝图是：%v\r\n", mat.Result)
}

func luckssim() {

	ws := 0
	var ps float32 = 0.00
	bs := 0

	for i := 0; i < 5; i++ {
		w, p, b := amusing.LuckyStarSim(10, 50, 6, 10)
		ws += w
		ps += p
		bs += b
	}
	fmt.Printf("%v,%v,%v", ws, ps/5, bs/5)

}

func jsonDemo() {
	gostudy.GetBaidu()
	ip := "223.5.5.5"
	tip := gostudy.GetIPFromTaobao(ip)
	fmt.Printf("%v的IP信息是：%+v。\r\n", ip, tip)

	ip = "114.114.114.114"
	jip := gostudy.GetIPFromTaobaoUseSimpjson(ip)
	fmt.Printf("%v的IP信息是：%+v。\r\n", ip, jip)
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

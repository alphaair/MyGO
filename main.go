package main

import "fmt"
//import "deltatch/common"
import "gostudy/amusing"

func main() {
	fmt.Println("Hello World,您好，世界!")
	//gostudy.Variable()

	//line := common.ReadLine("请输入一行字符：")
    //fmt.Println(line)
    //line = common.ReadLine("即将终止运行...")
    
    limit := 1000
    pnums := amusing.PrimeNumber(limit)
    fmt.Printf("%v以内的所有素数：%v",limit,pnums)
}

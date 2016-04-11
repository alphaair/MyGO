package main

import "fmt"
import "deltatch/common"
import "gostudy"
import "gostudy/amusing"

func main() {
     
    //basestudy()
    //primeNumber()
    
    mat := new(amusing.MatrixConceal)
    mat.Init()
    fmt.Printf("矩阵藏宝图是：%v",mat.Result)
}

func primeNumber() {
    limit := 1000
    pnums := amusing.PrimeNumber(limit)
    fmt.Printf("%v以内的所有素数：%v",limit,pnums)
}

func basestudy() {
    fmt.Println("Hello World,您好，世界!")
	gostudy.Variable()
    gostudy.CatchExceptioin()

	line := common.ReadLine("请输入一行字符：")
    fmt.Println(line)
    line = common.ReadLine("即将终止运行...")
}

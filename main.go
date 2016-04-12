package main

import "fmt"
import "unicode/utf8"
import "deltatch/common"
import "gostudy"
import "gostudy/amusing"

func main() {
     
    //basestudy()
    //primeNumber()
    
    text := "hello,GO语言。"
    buffer := make([]byte,0)
    
    for len(text)>0 {
        bf := make([]byte,3)
        
        r, size := utf8.DecodeRuneInString(text)
        text=text[size:]
        
        size = utf8.EncodeRune(bf,r)
        common.Append(buffer,size)
    }
    
    fmt.Printf("字符串UTF8编码形式：%v\r\n", buffer)
    buffer = []byte(text)
    fmt.Printf("字符串UTF8编码形式（硬转码）：%v\r\n", buffer)
   
    mat := new(amusing.MatrixConceal)
    mat.Init()
    fmt.Printf("矩阵藏宝图是：%v\r\n",mat.Result)
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

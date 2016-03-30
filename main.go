package main

import "fmt"
import "deltatch/common"

func main() {
	fmt.Println("Hello World,您好，世界!")
	//gostudy.Variable()

	line := common.ReadLine("请输入一行字符：")
	//fmt.Printf("keyborad input:%v", line)
    fmt.Println(line)
    line = common.ReadLine("即将终止运行...")
}

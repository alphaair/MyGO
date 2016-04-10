// Copyright alphaair 2016
// 这是一个开源学习代码文件

/*
Pakcage gostudy 是一个GO基础学习包，供应学习之用。
*/
package gostudy

import "fmt"

// Variable 函数是，GO变量语法
// value.无返回值
func Variable() {

	var num0 int = 3
	num1 := 3

	fmt.Printf("n1=%d,n2=%d", num0, num1)
}

// ThrowException 抛异常学习
func ThrowException() {
    panic("这是一个异常。")
    panic(404)
    //panic(Error("这是一个系统错误。"))
}

// CatchExceptioin 异常捕捉
func CatchExceptioin() {
    
    defer func(){
        if err := recover(); err != nil {
            fmt.Printf("捕获到异常:%v",err)
        }
    }()
    ThrowException()
    
    //err := recover()
    //fmt.Printf("捕获到异常:%v",err)
}
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

// Points 指针操作
func Points() {

	i := 765
	fmt.Printf("变量i原始值是：%v\r\n", i)

	p := &i //定义指针，指向变量i内存地址
	fmt.Printf("指定p指定的内存地址，即变量i的真实内存地址：%v\r\n", p)

	*p = -765 //利用指定赋值
	fmt.Printf("变量i通过指针修改后的值：%v\r\n", i)

}

// ThrowException 抛异常学习
func ThrowException() {
	panic("这是一个异常。")
	panic(404)
	//panic(Error("这是一个系统错误。"))
}

// CatchExceptioin 异常捕捉
func CatchExceptioin() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("捕获到异常:%v", err)
		}
	}()
	ThrowException()

	//err := recover()
	//fmt.Printf("捕获到异常:%v",err)
}

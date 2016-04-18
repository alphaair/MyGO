// Copyright alphaair 2016
// 这是GO多线程学习代码

package gostudy

import "fmt"
import "time"
import "math/rand"

// GoroutineEntry GO多线程学习入口点
func GoroutineEntry() {
	fmt.Println("Hello go thread...")
	for i := 1; i < 10; i++ {
		go thredDo(i)
	}
}

func thredDo(id int) {
	wt := rand.Int63n(11)
	fmt.Printf("我是标识:%v，请等待：%vS.", id, wt)
	wt *= int64(time.Millisecond)
	time.Sleep(time.Duration(wt))
}

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

	channelSync()
}

func thredDo(id int) {
	wt := rand.Int63n(11)
	fmt.Printf("我是标识:%v，请等待：%vS.\r\n", id, wt)
	wt *= int64(time.Millisecond)
	time.Sleep(time.Duration(wt))
}

// channelSync同步操作
func channelSync() {
	ch := make(chan int)

	go func() {
		//先等待另一线程启动
		time.Sleep(100 * time.Millisecond)
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Printf("交换值%v已经放入Channel.\r\n", i)
		}
	}()

	read := func() {
		for i := 0; i < 11; i++ {
			v := <-ch
			fmt.Printf("已经从Channel获取值：%v。\r\n", v)
		}
	}

	go read()

	time.Sleep(20 * time.Second)
}

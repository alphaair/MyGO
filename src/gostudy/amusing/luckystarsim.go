// Copyright alphaair 2016
// “幸运星座”游戏模拟推演示

package amusing

import (
	"fmt"
	"math/rand"
	"time"
)

// newNum生成一个指定范围内的随机数
func newNum(min int, max int) int {

	src := rand.NewSource(time.Now().UnixNano() / 1000 / 1000)
	r := rand.New(src)
	n := r.Intn(max + 1)

	for n < min {
		n = r.Intn(max + 1)
	}

	return n
}

// LuckyStarSim 幸运星座模拟游戏
func LuckyStarSim(integral, persions, period int) {

	fmt.Printf("游戏参数，积分数：%v，允许人数：%v，模拟期数：%v。\r\n", integral, persions, period)

	balance := 0
	wines := 0

	for i := 1; i <= period; i++ {

		num := newNum(1, 12)
		isw := true
		fmt.Printf("%3.f <%2.f> [", float32(i), float32(num))
		time.Sleep(200)

		for j := 0; j < persions; j++ {

			balance += integral
			pnum := newNum(1, 12)
			time.Sleep(200)

			if pnum == num {
				balance -= integral * 12
				isw = false
			}

			if j > 0 {
				fmt.Print(",")
			}
			fmt.Printf("%2.f", float32(pnum))
		}

		if isw {
			wines++
			fmt.Print("] W\r\n")
		} else {
			fmt.Print("] L\r\n")
		}

	}

	fmt.Printf("累计赢期数：%v\r\n", wines)
	fmt.Printf("赢比例：%%%5.2f \r\n", float32(wines)/float32(period)*float32(100))
	fmt.Printf("积分结余：%v \r\n", balance)
}

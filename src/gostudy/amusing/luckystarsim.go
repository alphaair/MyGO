// Copyright alphaair 2016
// “幸运星座”游戏模拟推演示

package amusing

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
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

// randomOrgApi 从random.Org生成随机数
func randomOrgApi(num int) []int {

	url := "https://www.random.org/integers/?num={num}&min=1&max=12&col=1&base=10&format=plain&rnd=new"
	url = strings.Replace(url, "{num}", strconv.Itoa(num), 1)

	rsp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求random.org随机数失败。")
		return nil
	}

	buffer, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("读取random.org随机数失败。")
		return nil
	}

	text := string(buffer)

	result := make([]int, num)
	ncodes := strings.Split(text, "\n")

	for i := 0; i < num; i++ {
		result[i], _ = strconv.Atoi(ncodes[i])
	}

	return result
}

// LuckyStarSim 幸运星座模拟游戏
func LuckyStarSim(integral, persions, multiple, period int) (wines int, proportion float32, balance int) {

	fname := strconv.Itoa(integral) + "_" + strconv.Itoa(persions) + "_" + strconv.Itoa(period) + ".log"
	file, _ := os.OpenFile(fname, os.O_CREATE|os.O_APPEND, os.ModeAppend)

	text := fmt.Sprintf("=========模拟参数，最多容纳人数：%v，赔率：%v，模拟期数：%v=========\r\n", persions, multiple, period)
	fmt.Print(text)
	file.WriteString(text)

	balance = 0
	wines = 0

	for i := 1; i <= period; i++ {

		//num := newNum(1, 12)
		nums := randomOrgApi(persions + 1)
		num := nums[0]
		isw := true

		text = fmt.Sprintf("%3.f <%2.f> [", float32(i), float32(num))
		fmt.Print(text)
		file.WriteString(text)

		//time.Sleep(200)

		for j := 1; j <= persions; j++ {

			balance += integral
			pnum := nums[j]
			//time.Sleep(200)

			if pnum == num {
				balance -= integral * multiple
				isw = false
			}

			if j > 1 {
				fmt.Print(",")
				file.WriteString(",")
			}
			text = fmt.Sprintf("%2.f", float32(pnum))
			fmt.Print(text)
			file.WriteString(text)

		}

		if isw {
			wines++
			fmt.Print("] W\r\n")
			file.WriteString("] W\r\n")
		} else {
			fmt.Print("] L\r\n")
			file.WriteString("] L\r\n")
		}

	}

	fmt.Print("********************************\r\n")
	file.WriteString("********************************\r\n")
	text = fmt.Sprintf("* 庄家赢期数：%v\r\n", wines)
	fmt.Print(text)
	file.WriteString(text)

	proportion = float32(wines) / float32(period) * float32(100)
	text = fmt.Sprintf("* 庄家赢概率：%%%5.2f \r\n", proportion)
	fmt.Print(text)
	file.WriteString(text)

	text = fmt.Sprintf("* 积分结余：%v \r\n", balance)
	fmt.Print(text)
	file.WriteString(text)

	fmt.Print("********************************\r\n")
	file.WriteString("********************************\r\n")

	defer file.Close()

	return
}

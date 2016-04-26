// copyright alphaair 2016
// 这是关于HTTP操作的演示

package gostudy

import "fmt"
import "io/ioutil"
import "net/http"
import "encoding/json"
import "github.com/bitly/go-simplejson"

// GetBaidu 方法请求百度热度主页
func GetBaidu() {

	rsp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Printf("访问百度失败，详细信息：%v\r\n。", err)
		return
	}

	//读取内容
	buffer, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("内容解析失败。")
		return
	}
	text := string(buffer)
	fmt.Printf("百度首页内容是：%v\r\n。", text)

	defer rsp.Body.Close()
}

// GetIPFromTaobao 通过淘宝IP开放数据库查询IP地址信息
func GetIPFromTaobao(ip string) Taobaoip {
	url := "http://ip.taobao.com//service/getIpInfo.php?ip=" + ip

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Alphaiar-GO-Robot")
	req.Header.Add("Sign", "Alphaair")

	client := &http.Client{}
	rsp, _ := client.Do(req)

	//buffer := []byte(`{"Code": 5, "Data": {"Country":"USA","Area":"Unkown","Owner":"Google"}}`)
	buffer, _ := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()

	fmt.Println(string(buffer))
	var tip Taobaoip
	json.Unmarshal(buffer, &tip)

	return tip
}

// GetIPFromTaobaoUseSimpjson 查询IP信息，并使用SimpleJSON解析
func GetIPFromTaobaoUseSimpjson(ip string) *simplejson.Json {
	url := "http://ip.taobao.com//service/getIpInfo.php?ip=" + ip

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Alphaiar-GO-Robot")
	req.Header.Add("Sign", "Alphaair")

	client := &http.Client{}
	rsp, _ := client.Do(req)

	jsdata, _ := simplejson.NewFromReader(rsp.Body)
	defer rsp.Body.Close()

	// var tip Taobaoip
	// json.Unmarshal(buffer, &tip)

	return jsdata
}

// Ipdetail 是淘宝开放IP数据库的一个查询结果明细
type Ipdetail struct {
	Atcountry string
	Atarea    string
}

// Taobaoip 是淘宝开放IP数据库的一个查询结果
type Taobaoip struct {
	ReCode int
	Data   Ipdetail
}

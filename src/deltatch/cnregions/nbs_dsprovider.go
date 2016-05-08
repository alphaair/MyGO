// Copyright alphaair 2016
// 国家统计局数据源提供程序

package cnregions

import (
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/iconv.v1"
	"net/http"
	"regexp"
)

// NbsDsProvider 国家统计局数据源提供程序
type NbsDsProvider struct {
	Name string
}

// NewNbsDsProvider 初始化一个新的提供程序
func NewNbsDsProvider() *NbsDsProvider {
	return &NbsDsProvider{"国家统计局数据源提供程序"}
}

func (slef *NbsDsProvider) requestHtml(url string) (*iconv.Reader, error) {

	cd, err := iconv.Open("GB2312", "uft8")
	if err != nil {
		return nil, err
	}

	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	nr := iconv.NewReader(cd, rsp.Body, 512)
	defer rsp.Body.Close()

	return nr, nil
}

// extractNum 从text提取出连续的数字
func (slef *NbsDsProvider) extractNum(text string) string {
	reg := regexp.MustCompile(`\d+`)
	return reg.FindString(text)
}

// GetProvinces 获取所有省级行政区域节点
func (self *NbsDsProvider) GetProvinces() ([]*RegionNode, error) {

	const url string = "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2014/index.html"
	rd, err := self.requestHtml(url)
	if err != nil {
		return nil, err
	}
	//doc, err := goquery.NewDocument(url)
	doc, err := goquery.NewDocumentFromReader(rd)
	if err != nil {
		return nil, err
	}

	provs := make([]*RegionNode, 31)
	doc.Find("tr.provincetr td a").Each(func(i int, pe *goquery.Selection) {

		href, _ := pe.Attr("href")

		regnode := new(RegionNode)
		regnode.Category = Province
		regnode.PrevCode = "0"
		regnode.Code = self.extractNum(href)
		regnode.Name = pe.Text()

		//provs = append(provs, regnode)
		provs[i] = regnode
	})

	return provs, nil
}

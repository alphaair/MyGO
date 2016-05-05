// Copyright alphaair 2016
// 国家统计局数据源提供程序

package cnregions

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

// NbsDsProvider国家统计局数据源提供程序
type NbsDsProvider struct {
	Name string
}

// NewNbsDsProvider 初始化一个新的提供程序
func NewNbsDsProvider() *NbsDsProvider {
	return &NbsDsProvider{"国家统计局数据源提供程序"}
}

//func requestHtml(url string) *

// extractNum 从text提取出连续的数字
func extractNum(text string) string {
	reg := regexp.MustCompile(`\d+`)
	return reg.FindString(text)
}

// GetProvinces 获取所有省级行政区域节点
func (slef *NbsDsProvider) GetProvinces() []*RegionNode {

	const url string = "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2014/index.html"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil
	}

	provs := make([]*RegionNode, 31)
	doc.Find("tr.provincetr td a").Each(func(i int, pe *goquery.Selection) {

		href, _ := pe.Attr("href")

		regnode := new(RegionNode)
		regnode.Category = Province
		regnode.PrevCode = "0"
		regnode.Code = extractNum(href)
		regnode.Name = pe.Text()

		//provs = append(provs, regnode)
		provs[i] = regnode
	})

	return provs
}

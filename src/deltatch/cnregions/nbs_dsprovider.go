// Copyright alphaair 2016
// 国家统计局数据源提供程序

package cnregions

import (
	dio "deltatch/io"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/guotie/gogb2312"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// NbsDsProvider 国家统计局数据源提供程序
type NbsDsProvider struct {
	Name string
}

// NewNbsDsProvider 初始化一个新的提供程序
func NewNbsDsProvider() *NbsDsProvider {
	return &NbsDsProvider{"国家统计局数据源提供程序"}
}

func (slef *NbsDsProvider) requestHtml(url string) (io.Reader, error) {

	rsp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("HTTP请求发生异常，详细信息：" + err.Error())
	}

	defer rsp.Body.Close()
	buffer, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.New("读取请求数据出错，详细信息：" + err.Error())
	}

	buffer, err, _, _ = gogb2312.ConvertGB2312(buffer)
	if err != nil {
		return nil, errors.New("GB2312编码转换失败，详细信息：" + err.Error())
	}

	rd := dio.NewMemoryStream(256)
	rd.Write(buffer)

	return rd, nil
}

// extractNum 从text提取出连续的数字
func (self *NbsDsProvider) extractNum(text string) string {
	i := strings.LastIndex(text, "/")
	if i > -1 {
		text = text[i+1:]
	}

	reg := regexp.MustCompile(`\d+`)
	return reg.FindString(text)
}

// GetProvinces 获取所有省级行政区域节点
func (self *NbsDsProvider) GetProvinces() ([]*RegionNode, error) {

	const url string = "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2014/index.html"
	rd, err := self.requestHtml(url)
	if err != nil {
		return nil, errors.New("请求失败，详细信息：" + err.Error())
	}

	//doc, err := goquery.NewDocument(url)
	doc, err := goquery.NewDocumentFromReader(rd)
	if err != nil {
		return nil, errors.New("HTML文档渲染失败：" + err.Error())
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

// GetCitys 获取指定省份的地级市
func (self *NbsDsProvider) GetCitys(code string) ([]*RegionNode, error) {

	url := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2014/{code}.html"
	url = strings.Replace(url, "{code}", code, 1)

	rd, err := self.requestHtml(url)
	if err != nil {
		return nil, errors.New("请求失败，详细信息：" + err.Error())
	}

	//doc, err := goquery.NewDocument(url)
	doc, err := goquery.NewDocumentFromReader(rd)
	if err != nil {
		return nil, errors.New("HTML文档渲染失败：" + err.Error())
	}

	citys := make([]*RegionNode, 0)
	doc.Find("tr.citytr").Each(func(i int, pe *goquery.Selection) {

		hfs := pe.Find("a")
		if hfs.Length() == 0 {
			//无下级时，即没有a标签
			hfs = pe.Find("td")
		}

		regnode := new(RegionNode)
		regnode.Category = City
		regnode.PrevCode = code
		regnode.Code, _ = hfs.First().Attr("href")
		regnode.Code = self.extractNum(regnode.Code)
		regnode.Name = hfs.Last().Text()

		citys = append(citys, regnode)
		//provs[i] = regnode
	})

	return citys, nil
}

// GetCountys 获取指定市级份的县级市
func (self *NbsDsProvider) GetCountys(code string) ([]*RegionNode, error) {

	url := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2014/{pcode}/{ccode}.html"
	url = strings.Replace(url, "{pcode}", code[0:2], 1)
	url = strings.Replace(url, "{ccode}", code, 1)

	rd, err := self.requestHtml(url)
	if err != nil {
		return nil, errors.New("请求失败，详细信息：" + err.Error())
	}

	doc, err := goquery.NewDocumentFromReader(rd)
	if err != nil {
		return nil, errors.New("HTML文档渲染失败：" + err.Error())
	}

	trs := doc.Find("tr.countytr")
	if trs.Length() == 0 {
		//部地级市直管镇，如东莞，中山
		trs = doc.Find("tr.towntr")
	}

	countys := make([]*RegionNode, trs.Length())
	trs.Each(func(i int, pe *goquery.Selection) {

		regnode := new(RegionNode)
		regnode.Category = County
		regnode.PrevCode = code

		hfs := pe.Find("a")
		if hfs.Length() == 0 {
			//无下级时，即没有a标签
			hfs = pe.Find("td")
			regnode.Code = hfs.First().Text()
		} else {
			regnode.Code, _ = hfs.First().Attr("href")
			regnode.Code = self.extractNum(regnode.Code)
		}
		regnode.Name = hfs.Last().Text()

		countys[i] = regnode
	})

	return countys, nil
}

// GetTowns 获取指定县级份的镇（乡）社区
func (self *NbsDsProvider) GetTowns(code string) ([]*RegionNode, error) {

	url := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2014/{pcode}/{ccode}/{ctcode}.html"
	url = strings.Replace(url, "{pcode}", code[0:2], 1)
	url = strings.Replace(url, "{ccode}", code[2:4], 1)
	url = strings.Replace(url, "{ctcode}", code, 1)

	rd, err := self.requestHtml(url)
	if err != nil {
		return nil, errors.New("请求失败，详细信息：" + err.Error())
	}

	doc, err := goquery.NewDocumentFromReader(rd)
	if err != nil {
		return nil, errors.New("HTML文档渲染失败：" + err.Error())
	}

	trs := doc.Find("tr.towntr")
	if trs.Length() == 0 {
		trs = doc.Find("tr.villagetr")
	}

	countys := make([]*RegionNode, trs.Length())
	trs.Each(func(i int, pe *goquery.Selection) {

		regnode := new(RegionNode)
		regnode.Category = Town
		regnode.PrevCode = code

		hfs := pe.Find("a")
		if hfs.Length() == 0 {
			//无下级时，即没有a标签
			hfs = pe.Find("td")
			regnode.Code = hfs.First().Text()
		} else {
			regnode.Code, _ = hfs.First().Attr("href")
			regnode.Code = self.extractNum(regnode.Code)
		}

		regnode.Name = hfs.Last().Text()

		countys[i] = regnode
	})

	return countys, nil
}

// GetTowns 获取指定镇级份的村、居委会
func (self *NbsDsProvider) GetVillage(code string) ([]*RegionNode, error) {

	url := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2014/{pcode}/{ccode}/{ctcode}/{tcode}.html"
	url = strings.Replace(url, "{pcode}", code[0:2], 1)
	url = strings.Replace(url, "{ccode}", code[2:4], 1)
	url = strings.Replace(url, "{ctcode}", code[4:6], 1)
	url = strings.Replace(url, "{tcode}", code, 1)

	fmt.Println(url)
	rd, err := self.requestHtml(url)
	if err != nil {
		return nil, errors.New("请求失败，详细信息：" + err.Error())
	}

	doc, err := goquery.NewDocumentFromReader(rd)
	if err != nil {
		return nil, errors.New("HTML文档渲染失败：" + err.Error())
	}

	trs := doc.Find("tr.villagetr")

	countys := make([]*RegionNode, trs.Length())
	trs.Each(func(i int, pe *goquery.Selection) {

		regnode := new(RegionNode)
		regnode.Category = Village
		regnode.PrevCode = code

		hfs := pe.Find("td")
		regnode.Code = hfs.First().Text()
		regnode.Name = hfs.Last().Text()

		countys[i] = regnode
	})

	return countys, nil
}

// Copyright alphaair 2016
// 国家统计局数据源提供程序

package cnregions

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/guotie/gogb2312"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

type MemoryStream struct {
	//暂存数据流
	buffer   []byte
	Length   int
	Position int
}

// NewMemoryStream 以指定大小初始化一个内存数据流
func NewMemoryStream(size int) *MemoryStream {
	stream := new(MemoryStream)
	stream.buffer = make([]byte, size)

	return stream
}

// Read 读取内存数据流的数据到p
func (self *MemoryStream) Read(p []byte) (n int, err error) {

	if self.Length == 0 || self.Position+1 >= self.Length {
		return 0, io.EOF
	}

	cp := 0
	for i, _ := range p {
		self.Position++
		if self.Position >= self.Length {
			break
		}
		p[i] = self.buffer[self.Position]
		cp++
	}

	return cp, nil
}

// Write 将数据p写入到内存流中
func (self *MemoryStream) Write(p []byte) (n int, err error) {

	self.buffer = p
	self.Length = len(p)
	self.Position = -1

	return len(p), nil
}

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

	rd := NewMemoryStream(256)
	rd.Write(buffer)

	return rd, nil
}

// extractNum 从text提取出连续的数字
func (self *NbsDsProvider) extractNum(text string) string {
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

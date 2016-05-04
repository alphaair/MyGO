// Copyright alphaair 2015
// 当前定义了行政区域数据提供程序接口

package cnregions

// IDsProvider 数据源提供程序，数据爬虫接口必须实现本接口。
type IDsProvider interface {
	GetProvinces() []RegionNode
	GetCitys(code string) []RegionNode
	GetCountys(code string) []RegionNode
	GetTowns(code string) []RegionNode
	GetVillages(code string) []RegionNode
}

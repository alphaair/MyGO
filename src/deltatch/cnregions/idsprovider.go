// Copyright alphaair 2015
// 当前定义了行政区域数据提供程序接口

package cnregions

// IDsProvider 数据源提供程序，数据爬虫接口必须实现本接口。
type IDsProvider interface {
	// GetProvinces 获取所有的省级行政区域
	GetProvinces() []RegionNode
	// GetCitys 获取code指示的省份下辖的所有市级行政区域
	GetCitys(code string) []RegionNode
	// GetCountrys 获取code指示的市下辖的所有县级行政区域
	GetCountys(code string) []RegionNode
	// GetCountrys 获取code指示的县下辖的所有镇级行政区域
	GetTowns(code string) []RegionNode
	// GetCountrys 获取code指示的镇下辖的所有村级行政区域
	GetVillages(code string) []RegionNode
}

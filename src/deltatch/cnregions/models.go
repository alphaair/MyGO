// Copyright alphaair 2016
// 定义关于中国行政区域数据模型

package cnregions

// RegoinCategory 表示行政区域级别
type RegoinCategory int

// 行级区域级别定义
const (
	// Province 表示省级（自治区、直辖市）级
	Province RegoinCategory = iota
	// City 表示中国家地级市级
	City
	// County 表示区（县、县级市）级
	County
	// Town 表示镇（乡、街道）级
	Town
	// Village 表示行政村（社区、居委会）级
	Village
)

// RegionNode 表示一级行政区域基本信息
type RegionNode struct {
	PrevCode string
	Code     string
	Name     string
	Category RegoinCategory
}

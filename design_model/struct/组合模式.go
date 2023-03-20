package main

import "fmt"

// 组合是一种结构型设计模式，你可以使用它将对象组合成树状结构，并且能像使用独立对象一样使用它们。
// 对于绝大多数需要生成树状结构的问题来说，组合都是非常受欢迎的解决方案。组合最主要的功能是在整个树状结构上递归调用方法并对结果进行汇总。

// 一般来说一个地区统计人口或经济总量，总是通过行政区划一层层上报汇总得出结果，区镇是最低一级行政区划，
// 需要落实统计人口及经济总量的工作，再上一级行政区划需要将所辖区镇的数据汇总统计，以此类推每一级行政区划都需要统计人口与经济总量，
// 就像一个倒过来的树状结构，各级行政区划统一的组件接口是统计人口与经济总量，区镇相当于最底层的叶子节点，中间级别行政区划相当于组合节点；
// 下面代码以苏州市为例；

type (
	// Region 行政区，作为组合模式component接口
	Region interface {
		Name() string    // 名称
		Population() int //人口
		GDP() float64    // gdp
	}
	// town 区镇，组合模式中相当于叶子节点
	town struct {
		name       string
		population int
		gdp        float64
	}
)

// NewTown 创建区镇，根据名称、人口、GDP
func NewTown(name string, population int, gdp float64) *town {
	return &town{
		name:       name,
		population: population,
		gdp:        gdp,
	}
}

func (c *town) Name() string {
	return c.name
}

func (c *town) Population() int {
	return c.population
}
func (c *town) GDP() float64 {
	return c.gdp
}

// cities 市，包括县市或者地市，组合模式中相当于composite
type cities struct {
	name    string
	regions map[string]Region
}

// NewCities 创建一个市
func NewCities(name string) *cities {
	return &cities{
		name:    name,
		regions: make(map[string]Region),
	}
}

func (c *cities) Name() string {
	return c.name
}

func (c *cities) Population() int {
	sum := 0
	for _, r := range c.regions {
		sum += r.Population()
	}
	return sum
}

func (c *cities) GDP() float64 {
	sum := 0.0
	for _, r := range c.regions {
		sum += r.GDP()
	}
	return sum
}

// Add 添加多个行政区
func (c *cities) Add(regions ...Region) {
	for _, r := range regions {
		c.regions[r.Name()] = r
	}
}

// Remove 递归删除行政区
func (c *cities) Remove(name string) {
	for n, r := range c.regions {
		if n == name {
			delete(c.regions, name)
			return
		}
		if city, ok := r.(*cities); ok {
			city.Remove(name)
		}
	}
}

func (c *cities) Regions() map[string]Region {
	return c.regions
}

func main() {
	gusu := NewTown("姑苏区", 100, 2000.00)
	fmt.Println(ShowRegionInfo(gusu))
	wuzhong := NewTown("吴中区", 150, 2600.00)
	fmt.Println(ShowRegionInfo(wuzhong))
	huqiu := NewTown("虎丘区", 80, 1800.00)
	fmt.Println(ShowRegionInfo(huqiu))

	kunshan := NewCities("昆山区")
	kunshan.Add(NewTown("玉山镇", 60, 1200.00),
		NewTown("周庄镇", 68, 1900.00),
		NewTown("花桥镇", 78, 2200.00))
	fmt.Println(ShowRegionInfo(kunshan))

	changshu := NewCities("常熟区")
	changshu.Add(NewTown("沙家浜镇", 55, 1100.00),
		NewTown("古里镇", 59, 1300.00),
		NewTown("辛庄镇", 68, 2100.00))
	fmt.Println(ShowRegionInfo(changshu))

	suzhou := NewCities("苏州市")
	suzhou.Add(gusu, wuzhong, huqiu, kunshan, changshu)
	fmt.Println(ShowRegionInfo(suzhou))

	// 省
	suzhouPro := NewCities("苏州省")
	suzhouPro.Add(suzhou)
	fmt.Println(ShowRegionInfo(suzhouPro))

	// 全国
	guoPro := NewCities("全国")
	guoPro.Add(suzhouPro)
	fmt.Println(ShowRegionInfo(guoPro))
}

func ShowRegionInfo(region Region) string {
	return fmt.Sprintf("%s, 人口:%d万, GDP:%.2f亿", region.Name(), region.Population(), region.GDP())
}

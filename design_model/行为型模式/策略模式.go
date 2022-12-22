package main

import "fmt"

// 策略模式是一种行为设计模式，它能让你定义一系列算法，并将每种算法分别放入独立的类中，以使算法的对象能够相互替换。
// 原始对象被称为上下文，它包含指向策略对象的引用并将执行行为的任务分派给策略对象。
// 为了改变上下文完成其工作的方式，其他对象可以使用另一个对象来替换当前链接的策略对象。
// 策略模式是最常用的设计模式，也是比较简单的设计模式，是以多态替换条件表达式重构方法的具体实现，是面向接口编程原则的最直接体现；

// case:北京是一个四季分明的城市，每个季节天气情况都有明显特点；
// 我们定义一个显示天气情况的季节接口，具体的四季实现，都会保存一个城市和天气情况的映射表，城市对象会包含季节接口，随着四季的变化.
// 天气情况也随之变化；
type (
	Season interface {
		ShowWeather(city string) string // 显示指定城市的天气情况
	}
	spring struct {
		weathers map[string]string // 存储不同城市春天气候
	}
	summer struct {
		weathers map[string]string // 存储不同城市夏天气候
	}
	autumn struct {
		weathers map[string]string // 存储不同城市秋天气候
	}
	winter struct {
		weathers map[string]string // 存储不同城市冬天气候
	}
	// City 城市
	City struct {
		name    string
		feature string
		season  Season
	}
)

func NewSpring() *spring {
	return &spring{
		weathers: map[string]string{"北京": "干燥多风", "昆明": "清凉舒适"},
	}
}

func (s *spring) ShowWeather(city string) string {
	return fmt.Sprintf("%s的春天，%s;", city, s.weathers[city])
}

func NewSummer() *summer {
	return &summer{
		weathers: map[string]string{"北京": "高温多雨", "昆明": "清凉舒适"},
	}
}

func (s *summer) ShowWeather(city string) string {
	return fmt.Sprintf("%s的夏天，%s;", city, s.weathers[city])
}

func NewAutumn() *autumn {
	return &autumn{
		weathers: map[string]string{"北京": "凉爽舒适", "昆明": "清凉舒适"},
	}
}

func (a *autumn) ShowWeather(city string) string {
	return fmt.Sprintf("%s的秋天，%s;", city, a.weathers[city])
}

func NewWinter() *winter {
	return &winter{
		weathers: map[string]string{"北京": "干燥寒冷", "昆明": "清凉舒适"},
	}
}

func (w *winter) ShowWeather(city string) string {
	return fmt.Sprintf("%s的冬天，%s;", city, w.weathers[city])
}

// NewCity 根据名称及季候特征创建城市
func NewCity(name, feature string) *City {
	return &City{
		name:    name,
		feature: feature,
	}
}

// SetSeason 设置不同季节，类似天气在不同季节的不同策略
func (c *City) SetSeason(season Season) {
	c.season = season
}

// String 显示城市的气候信息
func (c *City) String() string {
	return fmt.Sprintf("%s%s，%s", c.name, c.feature, c.season.ShowWeather(c.name))
}

func main() {
	Beijing := NewCity("北京", "四季分明")

	Beijing.SetSeason(NewSpring())
	fmt.Println(Beijing.String())

	Beijing.SetSeason(NewSummer())
	fmt.Println(Beijing)

	Beijing.SetSeason(NewAutumn())
	fmt.Println(Beijing)

	Beijing.SetSeason(NewWinter())
	fmt.Println(Beijing)
}

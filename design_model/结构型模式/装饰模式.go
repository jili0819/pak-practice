package main

import "fmt"

// 装饰是一种结构设计模式，允许你通过将对象放入特殊封装对象中来为原对象增加新的行为。
// 由于目标对象和装饰器遵循同一接口，因此你可用装饰来对对象进行无限次的封装。结果对象将获得所有封装器叠加而来的行为。

// 地铁进站的过程一般情况下只需要买票，检票进站，但是如果你是带行李，就需要进行安全检查，
// 如果是疫情时期，就需要进行疫情防护检查，比如戴口罩、测量体温等，这里买票进站相当于通用进站流程，安检及防疫检查就相当于加强的修饰行为。
// 通过添加相同的方法，实现层层方法调用

type (
	// Station 车站，修饰器模式统一接口
	Station interface {
		Enter() string // 进站
	}

	// subwayStation 地铁站
	subwayStation struct {
		name string
	}
	// securityCheckDecorator 进站安检修饰器
	securityCheckDecorator struct {
		station Station
	}
	// epidemicProtectionDecorator 进站疫情防护修饰器
	epidemicProtectionDecorator struct {
		station Station
	}
)

// NewSubwayStation 创建指定站名地铁站
func NewSubwayStation(name string) *subwayStation {
	return &subwayStation{
		name: name,
	}
}

// Enter 进地铁站
func (s *subwayStation) Enter() string {
	return fmt.Sprintf("买票进入%s地铁站。", s.name)
}

func NewSecurityCheckDecorator(station Station) *securityCheckDecorator {
	return &securityCheckDecorator{
		station: station,
	}
}

func (s *securityCheckDecorator) Enter() string {
	return "行李通过安检=》" + s.station.Enter()
}

func NewEpidemicProtectionDecorator(station Station) *epidemicProtectionDecorator {
	return &epidemicProtectionDecorator{
		station: station,
	}
}

func (e *epidemicProtectionDecorator) Enter() string {
	return "测量体温，佩戴口罩=》" + e.station.Enter()
}

func main() {
	xierqiStation := NewSubwayStation("西二旗")
	fmt.Println(EnhanceEnterStationProcess(xierqiStation, false, false).Enter())
	fmt.Println(EnhanceEnterStationProcess(xierqiStation, true, false).Enter())
	fmt.Println(EnhanceEnterStationProcess(xierqiStation, true, true).Enter())
}

// EnhanceEnterStationProcess 根据是否有行李，是否处于疫情，增加进站流程
func EnhanceEnterStationProcess(station Station, hasLuggage bool, hasEpidemic bool) Station {
	if hasLuggage {
		station = NewSecurityCheckDecorator(station)
	}
	if hasEpidemic {
		station = NewEpidemicProtectionDecorator(station)
	}
	return station
}

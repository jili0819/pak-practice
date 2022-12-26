package main

import (
	"bytes"
	"fmt"
)

// 享元是一种结构型设计模式，它允许你在消耗少量内存的情况下支持大量对象。
// 模式通过共享多个对象的部分状态来实现上述功能。换句话来说，享元会将不同对象的相同数据进行缓存以节省内存。

// 北京出租车调度系统，需要每隔一分钟记录一下全市出租车的位置信息，假设为了提高系统响应速度，近一天的数据需要存储在内存中，
// 每个位置信息包括出租车辆信息及位置信息，位置信息在系统中就是一个(x,y)坐标，车辆信息包括车的号牌，颜色，品牌和所属公司，
// 在调度系统存储的出租车行驶轨迹中，位置是实时在变化的，但车辆信息就可以通过享元模式共用一个对象引用，来减少内存消耗

// Taxi 出租车，享元对象，保存不变的内在属性信息
type Taxi struct {
	licensePlate string // 车牌
	color        string // 颜色
	brand        string // 汽车品牌
	company      string // 所属公司
}

// LocateFor 获取定位信息
func (t *Taxi) LocateFor(monitorMap string, x, y int) string {
	return fmt.Sprintf("%s,对于车牌号%s,%s,%s品牌,所属%s公司,定位(%d,%d)", monitorMap,
		t.licensePlate, t.color, t.brand, t.company, x, y)
}

// taxiFactoryInstance 出租车工厂单例
var taxiFactoryInstance = &TaxiFactory{
	taxis: make(map[string]*Taxi),
}

// GetTaxiFactory 获取出租车工厂单例
func GetTaxiFactory() *TaxiFactory {
	return taxiFactoryInstance
}

// TaxiFactory 出租车工厂类
type TaxiFactory struct {
	taxis map[string]*Taxi // key为车牌号
}

// getTaxi 获取出租车
func (f *TaxiFactory) AddTaxi(licensePlate, color, brand, company string) *Taxi {
	if _, ok := f.taxis[licensePlate]; !ok {
		f.taxis[licensePlate] = &Taxi{
			licensePlate: licensePlate,
			color:        color,
			brand:        brand,
			company:      company,
		}
	}
	return f.taxis[licensePlate]
}

// getTaxi 获取出租车
func (f *TaxiFactory) getTaxi(licensePlate string) *Taxi {
	if _, ok := f.taxis[licensePlate]; !ok {
		return nil
	}
	return f.taxis[licensePlate]
}

// TaxiPosition 出租车位置信息 x,y为外在数据信息，taxi为内在数据信息（享元对象）
type TaxiPosition struct {
	x int
	y int
}

func NewTaxiPosition(x, y int) *TaxiPosition {
	return &TaxiPosition{
		x: x,
		y: y,
	}
}

// TaxiDispatcher 出租车调度系统
type TaxiDispatcher struct {
	name   string
	traces map[string][]*TaxiPosition // 存储出租车当天轨迹信息，key为车牌号
}

func NewTaxiDispatcher(name string) *TaxiDispatcher {
	return &TaxiDispatcher{
		name:   name,
		traces: make(map[string][]*TaxiPosition),
	}
}

// AddCar 添加车辆
func (t *TaxiDispatcher) AddCar(licensePlate, color, brand, company string) {
	GetTaxiFactory().AddTaxi(licensePlate, color, brand, company)
}

// AddTrace 添加轨迹
func (t *TaxiDispatcher) AddTrace(licensePlate string, x, y int) {
	t.traces[licensePlate] = append(t.traces[licensePlate], NewTaxiPosition(x, y))
}

// ShowTraces 显示轨迹
func (t *TaxiDispatcher) ShowTraces(licensePlate string) string {
	bytesBuf := bytes.Buffer{}
	taxi := GetTaxiFactory().getTaxi(licensePlate)
	for _, v := range t.traces[licensePlate] {
		bytesBuf.WriteString(taxi.LocateFor(t.name, v.x, v.y))
		bytesBuf.WriteByte('\n')
	}
	return bytesBuf.String()
}

func main() {
	dispatcher := NewTaxiDispatcher("北京市出租车调度系统")
	dispatcher.AddCar("京B.123456", "黄色", "北京现代", "北汽")
	dispatcher.AddTrace("京B.123456", 10, 20)
	dispatcher.AddTrace("京B.123456", 20, 30)
	dispatcher.AddTrace("京B.123456", 30, 40)
	dispatcher.AddTrace("京B.123456", 40, 50)

	dispatcher.AddCar("京B.567890", "红色", "一汽大众", "首汽")
	dispatcher.AddTrace("京B.567890", 20, 40)
	dispatcher.AddTrace("京B.567890", 50, 50)

	fmt.Println(dispatcher.ShowTraces("京B.123456"))
	fmt.Println(dispatcher.ShowTraces("京B.567890"))
}

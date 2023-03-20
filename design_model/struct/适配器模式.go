package main

import "fmt"

// 适配器模式是一种结构型设计模式，它能使接口不兼容的对象能够相互合作。
// 适配器可担任两个对象间的封装器，它会接收对于一个对象的调用， 并将其转换为另一个对象可识别的格式和接口。

// 通过充电宝给不同充电接口的手机充电是一个非常符合适配器模式特征的生活示例；
// 一般充电宝提供USB电源输出接口，手机充电输入接口则分为两类一是苹果手机的lightning接口，另一类是安卓手机的typeC接口，
// 这两类接口都需要通过适配电源线连接充电宝的USB接口，这里USB接口就相当于充电宝的通用接口，lightning或typeC接口要想充电需要通过充电线适配。

type (
	// HuaweiPlug 华为手机充电插槽接口
	HuaweiPlug interface {
		ConnectTypeC() string
	}

	// HuaweiPhone 华为系列手机
	HuaweiPhone struct {
		model string
	}

	// ApplePlug 苹果手机充电插槽
	ApplePlug interface {
		ConnectLightning() string
	}

	// IPhone 苹果系列手机
	IPhone struct {
		model string
	}
)

// NewHuaweiPhone 华为手机创建方法
func NewHuaweiPhone(model string) *HuaweiPhone {
	return &HuaweiPhone{
		model: model,
	}
}

// ConnectTypeC 华为手机TypeC充电插槽
func (h *HuaweiPhone) ConnectTypeC() string {
	return fmt.Sprintf("%v connect typeC plug", h.model)
}

// NewIPhone 苹果手机创建方法
func NewIPhone(model string) *IPhone {
	return &IPhone{
		model: model,
	}
}

// ConnectLightning 苹果手机Lightning充电插槽
func (i *IPhone) ConnectLightning() string {
	return fmt.Sprintf("%v connect lightning plug", i.model)
}

// CommonPlug 通用的USB电源插槽
type CommonPlug interface {
	ConnectUSB() string
}

// HuaweiPhonePlugAdapter 华为TypeC充电插槽适配通用USB充电插槽
type HuaweiPhonePlugAdapter struct {
	huaweiPhone HuaweiPlug
}

// NewHuaweiPhonePlugAdapter 创建华为手机适配USB充电插槽适配器
func NewHuaweiPhonePlugAdapter(huaweiPhone HuaweiPlug) *HuaweiPhonePlugAdapter {
	return &HuaweiPhonePlugAdapter{
		huaweiPhone: huaweiPhone,
	}
}

// ConnectUSB 链接USB
func (h *HuaweiPhonePlugAdapter) ConnectUSB() string {
	return fmt.Sprintf("%v adapt to usb ", h.huaweiPhone.ConnectTypeC())
}

// ApplePhonePlugAdapter 苹果Lightning充电插槽适配通用USB充电插槽
type ApplePhonePlugAdapter struct {
	iPhone ApplePlug
}

// NewApplePhonePlugAdapter 创建苹果手机适配USB充电插槽适配器
func NewApplePhonePlugAdapter(iPhone ApplePlug) *ApplePhonePlugAdapter {
	return &ApplePhonePlugAdapter{
		iPhone: iPhone,
	}
}

// ConnectUSB 链接USB
func (a *ApplePhonePlugAdapter) ConnectUSB() string {
	return fmt.Sprintf("%v adapt to usb ", a.iPhone.ConnectLightning())
}

// PowerBank 充电宝
type PowerBank struct {
	brand string
}

// Charge 支持通用USB接口充电
func (p *PowerBank) Charge(plug CommonPlug) string {
	return fmt.Sprintf("%v power bank connect usb plug, start charge for %v", p.brand, plug.ConnectUSB())
}

func main() {
	huaweiMate40Pro := NewHuaweiPhone("华为 mate40 pro")
	iphone13MaxPro := NewIPhone("苹果 iphone13 pro max")

	powerBank := &PowerBank{"飞利浦"}
	fmt.Println(powerBank.Charge(NewHuaweiPhonePlugAdapter(huaweiMate40Pro)))
	fmt.Println(powerBank.Charge(NewApplePhonePlugAdapter(iphone13MaxPro)))
}

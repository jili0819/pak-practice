package main

import (
	"bytes"
	"fmt"
)

// 代理是一种结构型设计模式，让你能提供真实服务对象的替代品给客户端使用。
// 代理接收客户端的请求并进行一些处理 （访问控制和缓存等）， 然后再将请求传递给服务对象。
// 代理对象拥有和服务对象相同的接口，这使得当其被传递给客户端时可与真实对象互换。
// 修饰与代理是非常相似的设计模式，都是基于组合设计原则，也就是说一个对象应该将部分工作委派给另一个对象。
// 但两者之间不同点我认为是，修饰器模式总是要执行服务对象，对于执行之前或执行之后结果进行加强，服务对象基本是客户端创建好再嵌套外层的修饰对象；
// 而代理模式不一定执行服务对象，有可能通过缓存，延迟加载等没 有访问服务对象，同时服务对象什么时候创建也是由代理类决定的。

// 房屋中介代理帮助房东卖房子，这个过程就是一个代理模式的过程，中介会收集尽量多的卖房信息，并通过各种渠道发布，
// 同时中介会随时带客户看房，并初步商讨价格，如果达成初步购买意向，才会约房东讨论房屋价格，最后签约卖房；
// 房屋中介与房东都实现卖房接口，中介会提前坐一些前期工作，如果都没问题，才会约房东执行真正的签约卖房流程。

type (
	// HouseSeller 房屋出售者
	HouseSeller interface {
		SellHouse(address string, buyer string) string
	}

	// houseProxy 房产中介代理
	houseProxy struct {
		houseSeller HouseSeller
	}
)

func NewHouseProxy(houseSeller HouseSeller) *houseProxy {
	return &houseProxy{
		houseSeller: houseSeller,
	}
}

// SellHouse 中介卖房，看房->初步谈价->最终和房东签协议
func (h *houseProxy) SellHouse(address string, buyer string) string {
	buf := bytes.Buffer{}
	buf.WriteString(h.viewHouse(address, buyer) + "\n")
	buf.WriteString(h.preBargain(address, buyer) + "\n")
	buf.WriteString(h.houseSeller.SellHouse(address, buyer))
	return buf.String()
}

// viewHouse 看房介绍基本情况
func (h *houseProxy) viewHouse(address string, buyer string) string {
	return fmt.Sprintf("带买家%s看位于%s的房屋，并介绍基本情况", buyer, address)
}

// preBargain 初步沟通价格
func (h *houseProxy) preBargain(address string, buyer string) string {
	return fmt.Sprintf("讨价还价后，初步达成购买意向")
}

// houseOwner 房东
type houseOwner struct{}

// SellHouse 房东卖房，商讨价格，签署购房协议
func (h *houseOwner) SellHouse(address string, buyer string) string {
	return fmt.Sprintf("最终商讨价格后，与%s签署购买地址为%s的购房协议。", buyer, address)
}

func main() {
	proxy := NewHouseProxy(&houseOwner{})
	fmt.Println(proxy.SellHouse("北京市海淀区中关村大街，2号院1号楼4单元502室", "李四"))
}

package main

import "fmt"

// 外观是一种结构型设计模式，能为复杂系统、程序库或框架提供一个简单 （但有限） 的接口。
// 尽管外观模式降低了程序的整体复杂度，但它同时也有助于将不需要的依赖移动到同一个位置。

// 用户在淘宝电商系统买商品时，只需要选好商品在结算页点击提交即可完成下单；
// 在客户端系统仅需要一个创建订单的方法，但是整个订单的生成需要很多步骤，比如查询用户配送地址，查询商品价格，使用优惠券，扣减商品库存，支付相应价钱等。

type (
	// TaobaoFacade 淘宝网站门面，在淘宝网站下单涉及到多个系统配合调用，包括用户系统，商品系统，优惠券系统，库存系统，支付系统，最终生成订单
	TaobaoFacade struct {
		userService    *UserService
		productService *ProductService
		couponService  *CouponService
		stockService   *StockService
		paymentService *PaymentService
	}
)

// NewTaobaoFacade 创建淘宝网站
func NewTaobaoFacade() *TaobaoFacade {
	return &TaobaoFacade{
		userService: &UserService{},
		productService: &ProductService{
			products: map[string]float64{"笔记本电脑": 6666.66},
		},
		couponService:  &CouponService{},
		stockService:   &StockService{},
		paymentService: &PaymentService{},
	}
}

// CreateOrder 根据用户名，商品名，商品数量生成购买订单
func (t *TaobaoFacade) CreateOrder(userName string, productName string, count int) string {
	// 使用优惠券
	couponInfo := t.couponService.useCoupon()
	// 扣减库存
	stockInfo := t.stockService.decreaseFor(productName, count)
	// 计算商品总价
	sumPrice := t.productService.getProductPrice(productName) * float64(count)
	// 支付价格
	payInfo := t.paymentService.pay(sumPrice)
	return fmt.Sprintf("用户%s,购买了%d件%s商品,%s,%s,%s,送货到%s", userName, count, productName, couponInfo,
		stockInfo, payInfo, t.userService.getUserAddress(userName))
}

// UserService 用户系统
type UserService struct{}

func (u *UserService) getUserAddress(userName string) string {
	return fmt.Sprintf("%s地址是:北京市海淀区中关村大街，1号院2号楼3单元402", userName)
}

// ProductService 商品系统
type ProductService struct {
	products map[string]float64
}

func (p *ProductService) getProductPrice(productName string) float64 {
	return p.products[productName]
}

// CouponService 优惠券系统
type CouponService struct{}

func (c *CouponService) useCoupon() string {
	return "使用满100减20优惠券"
}

// StockService 库存系统
type StockService struct{}

func (s *StockService) decreaseFor(productName string, count int) string {
	return fmt.Sprintf("扣减%d件%s商品库存", count, productName)
}

// PaymentService 支付系统
type PaymentService struct{}

func (p *PaymentService) pay(amount float64) string {
	return fmt.Sprintf("支付金额%.2f", amount)
}

func main() {
	// 通过门面模式，隐藏下单过程中，后端多个系统的复杂交互
	taobao := NewTaobaoFacade()
	fmt.Println(taobao.CreateOrder("张三", "笔记本电脑", 1))
}

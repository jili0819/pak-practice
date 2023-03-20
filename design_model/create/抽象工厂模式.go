package main

import "fmt"

// 抽象工厂是一种创建型设计模式，它能创建一系列相关的对象，而无需指定其具体类。
// 抽象工厂定义了用于创建不同产品的接口，但将实际的创建工作留给了具体工厂类。每个工厂类型都对应一个特定的产品变体。
// 在创建产品时，客户端代码调用的是工厂对象的构建方法，而不是直接调用构造函数 （ new操作符）。
// 由于一个工厂对应一种产品变体，因此它创建的所有产品都可相互兼容。
// 客户端代码仅通过其抽象接口与工厂和产品进行交互。该接口允许同一客户端代码与不同产品进行交互。你只需创建一个具体工厂类并将其传递给客户端代码即可。

// 厨师准备一餐时，会分别做吃的和喝的，根据早、中、晚三餐饮食习惯，会分别制作不同的饮食，厨师就相当于抽象工厂，
// 制作三餐的不同烹饪方式就好比不同抽象工厂的实现。

// Cook 厨师接口，抽象工厂
type Cook interface {
	// MakeFood 制作主食
	MakeFood() Food
	// MakeDrink 制作饮品
	MakeDrink() Drink
}

// Food 主食接口
type Food interface {
	// Eaten 被吃
	Eaten() string
}

// Drink 饮品接口
type Drink interface {
	// Drunk 被喝
	Drunk() string
}

// breakfastCook 早餐厨师
type breakfastCook struct{}

func NewBreakfastCook() *breakfastCook {
	return &breakfastCook{}
}

func (b *breakfastCook) MakeFood() Food {
	return &cakeFood{"切片面包"}
}

func (b *breakfastCook) MakeDrink() Drink {
	return &gruelDrink{"小米粥"}
}

// lunchCook 午餐厨师
type lunchCook struct{}

func NewLunchCook() *lunchCook {
	return &lunchCook{}
}

func (l *lunchCook) MakeFood() Food {
	return &dishFood{"烤全羊"}
}

func (l *lunchCook) MakeDrink() Drink {
	return &sodaDrink{"冰镇可口可乐"}
}

// dinnerCook 晚餐厨师
type dinnerCook struct{}

func NewDinnerCook() *dinnerCook {
	return &dinnerCook{}
}

func (d *dinnerCook) MakeFood() Food {
	return &noodleFood{"大盘鸡拌面"}
}

func (d *dinnerCook) MakeDrink() Drink {
	return &soupDrink{"西红柿鸡蛋汤"}
}

// cakeFood 蛋糕
type cakeFood struct {
	cakeName string
}

func (c *cakeFood) Eaten() string {
	return fmt.Sprintf("%v被吃", c.cakeName)
}

// dishFood 菜肴
type dishFood struct {
	dishName string
}

func (d *dishFood) Eaten() string {
	return fmt.Sprintf("%v被吃", d.dishName)
}

// noodleFood 面条
type noodleFood struct {
	noodleName string
}

func (n *noodleFood) Eaten() string {
	return fmt.Sprintf("%v被吃", n.noodleName)
}

// gruelDrink 粥
type gruelDrink struct {
	gruelName string
}

func (g *gruelDrink) Drunk() string {
	return fmt.Sprintf("%v被喝", g.gruelName)
}

// sodaDrink 汽水
type sodaDrink struct {
	sodaName string
}

func (s *sodaDrink) Drunk() string {
	return fmt.Sprintf("%v被喝", s.sodaName)
}

// soupDrink 汤
type soupDrink struct {
	soupName string
}

func (s *soupDrink) Drunk() string {
	return fmt.Sprintf("%v被喝", s.soupName)
}

func main() {
	fmt.Printf("breakfast: %v\n", HaveMeal(NewBreakfastCook()))
	fmt.Printf("lunch: %v\n", HaveMeal(NewLunchCook()))
	fmt.Printf("dinner: %v\n", HaveMeal(NewDinnerCook()))
}

// HaveMeal 吃饭
func HaveMeal(cook Cook) string {
	return fmt.Sprintf("%s %s", cook.MakeFood().Eaten(), cook.MakeDrink().Drunk())
}

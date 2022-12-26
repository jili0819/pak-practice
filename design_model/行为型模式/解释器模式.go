package main

import (
	"fmt"
	"strings"
)

// 解释器模式用于描述如何使用面向对象语言构成一个简单的语言解释器。
// 在某些情况下，为了更好地描述某一些特定类型的问题，我们可以创建一种新的语言，这种语言拥有自己的表达式和结构，即文法规则，
// 这些问题的实例将对应为该语言中的句子。此时，可以使用解释器模式来设计这种新的语言。
// 对解释器模式的学习能够加深我们对面向对象思想的理解，并且掌握编程语言中文法规则的解释过程。

// 定义一个解析特征值的语句解释器，提供是否包含特征值的终结表达式，并提供或表达式与且表达式，
// 同时，生成南极洲特征判断表达式，及美国人特征判断表达式，最后测试程序根据对象特征值描述，通过表达式判断是否为真

// Expression 表达式接口，包含一个解释方法
type Expression interface {
	Interpret(context string) bool
}

// terminalExpression 终结符表达式，判断表达式中是否包含匹配数据
type terminalExpression struct {
	matchData string
}

func NewTerminalExpression(matchData string) *terminalExpression {
	return &terminalExpression{matchData: matchData}
}

// Interpret 判断是否包含匹配字符
func (t *terminalExpression) Interpret(context string) bool {
	if strings.Contains(context, t.matchData) {
		return true
	}
	return false
}

// orExpression 或表达式
type orExpression struct {
	expresses []Expression
}

func NewOrExpression(expresses ...Expression) *orExpression {
	return &orExpression{
		expresses: expresses,
	}
}

func (o *orExpression) Interpret(context string) bool {
	for _, v := range o.expresses {
		if v.Interpret(context) {
			return true
		}
	}
	return false
}

// andExpression 与表达式
type andExpression struct {
	expresses []Expression
}

func NewAndExpression(expresses ...Expression) *andExpression {
	return &andExpression{
		expresses: expresses,
	}
}

func (o *andExpression) Interpret(context string) bool {
	for _, v := range o.expresses {
		if !v.Interpret(context) {
			return false
		}
	}
	return true
}

func main() {
	isAntarcticaExpression := generateCheckAntarcticaExpression()
	// 大洲描述1
	continentDescription1 := "此大洲生活着大量企鹅，全年低温，并且伴随着有暴风雪"
	fmt.Printf("%s，是否是南极洲？%t\n", continentDescription1, isAntarcticaExpression.Interpret(continentDescription1))
	// 大洲描述2
	continentDescription2 := "此大洲生活着狮子，全年高温多雨"
	fmt.Printf("%s，是否是南极洲？%t\n", continentDescription2, isAntarcticaExpression.Interpret(continentDescription2))

	isAmericanExpression := generateCheckAmericanExpression()
	peopleDescription1 := "此人生活在北美洲的黑人，说着英语，持有美国绿卡"
	fmt.Printf("%s，是否是美国人？%t\n", peopleDescription1, isAmericanExpression.Interpret(peopleDescription1))

	peopleDescription2 := "此人生活在欧洲，说着英语，是欧洲议会议员"
	fmt.Printf("%s，是否是南极洲？%t\n", peopleDescription2, isAmericanExpression.Interpret(peopleDescription2))

}

// generateCheckAntarcticaExpression 生成校验是否是南极洲表达式
func generateCheckAntarcticaExpression() Expression {
	// 判断南极洲的动物，或关系
	animalExpression := NewOrExpression(NewTerminalExpression("企鹅"),
		NewTerminalExpression("蓝鲸"))
	// 判断南极洲的天气，与关系
	weatherExpression := NewAndExpression(NewTerminalExpression("低温"),
		NewTerminalExpression("暴风雪"))
	// 最终返回动物与天气的与关系
	return NewAndExpression(animalExpression, weatherExpression)
}

// generateCheckAmericanExpression 生成检查美国人表达式
func generateCheckAmericanExpression() Expression {
	// 人种判断，或关系
	raceExpression := NewOrExpression(NewTerminalExpression("白人"),
		NewTerminalExpression("黑人"))
	// 生活方式，与关系
	lifeStyleExpression := NewAndExpression(NewTerminalExpression("英语"),
		NewTerminalExpression("北美洲"))
	// 身份，与关系
	identityExpression := NewAndExpression(lifeStyleExpression, NewTerminalExpression("美国绿卡"))
	return NewAndExpression(raceExpression, identityExpression)
}

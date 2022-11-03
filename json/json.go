package main

import (
	"fmt"
)

func main() {

	fmt.Println(6 % 10)
	//fmt.Println(a.Mul(decimal.NewFromFloat(0.8)).Floor().IntPart())
	return
	// json 数组带float64
	//var as []A
	/*as = append(as, A{
		Day:     1,
		Percent: 1,
	}, A{
		Day:     2,
		Percent: 0.8,
	})
	bytes, _ := json.Marshal(as)
	fmt.Println(string(bytes))*/
	jsonStr := "提示：\n1、最多可同时置顶%d个市、%d个省\n2、为保证效果，置顶成功至使用结束前，该招工信息将无法修改\n3、招工信息在移动端置顶展示，不同城市/工种实际曝光量，随竞争情况有所差异\n4、置顶价格根据置顶城市、工种等因素产生波动，最终价格以购买前平台展示为准"
	fmt.Println(fmt.Sprintf(jsonStr, 3, 2))
}

type A struct {
	Day     int64   `json:"day"`
	Percent float64 `json:"percent"`
}

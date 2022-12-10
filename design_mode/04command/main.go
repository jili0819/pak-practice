package main

import "fmt"

func main() {
	// 创建电饭煲，命令接受者
	electricCooker := new(ElectricCooker)
	// 创建电饭煲指令触发器
	electricCookerInvoker := new(ElectricCookerInvoker)
	// 停止
	shutdownCommand := NewShutdownCommand(electricCooker)

	// 蒸饭
	steamRiceCommand := NewSteamRiceCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(steamRiceCommand)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())
	electricCookerInvoker.SetCookCommand(shutdownCommand)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())

	// 煮粥
	cookCongeeCommand := NewCookCongeeCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(cookCongeeCommand)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())
	electricCookerInvoker.SetCookCommand(shutdownCommand)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())
}

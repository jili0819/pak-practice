package main

import "fmt"

// 命令模式
// 命令模式是一种行为设计模式，它可将请求转换为一个包含与请求相关的所有信息的独立对象。
// 该转换让你能根据不同的请求将方法参数化、延迟请求执行或将其放入队列中，且能实现可撤销操作。
// （二）示例
// 控制电饭煲做饭是一个典型的命令模式的场景，电饭煲的控制面板会提供设置煮粥、蒸饭模式，及开始和停止按钮.
// 电饭煲控制系统会根据模式的不同设置相应的火力，压强及时间等参数；
// 煮粥，蒸饭就相当于不同的命令，开始按钮就相当命令触发器，设置好做饭模式，点击开始按钮电饭煲就开始运行，同时还支持停止命令；

type (
	// ElectricCooker 电饭煲
	ElectricCooker struct {
		fire     string // 火力
		pressure string // 压力
	}

	// CookCommand 指令接口
	CookCommand interface {
		Execute() string // 指令执行方法
	}
	// steamRiceCommand 蒸饭指令
	steamRiceCommand struct {
		electricCooker *ElectricCooker // 电饭煲
	}
	// cookCongeeCommand 煮粥指令
	cookCongeeCommand struct {
		electricCooker *ElectricCooker
	}
	// shutdownCommand 停止指令
	shutdownCommand struct {
		electricCooker *ElectricCooker
	}
	// ElectricCookerInvoker 电饭煲指令触发器
	ElectricCookerInvoker struct {
		cookCommand CookCommand
	}
)

// NewSteamRiceCommand 蒸饭
func NewSteamRiceCommand(electricCooker *ElectricCooker) *steamRiceCommand {
	return &steamRiceCommand{
		electricCooker: electricCooker,
	}
}

func (s *steamRiceCommand) Execute() string {
	return "蒸饭:" + s.electricCooker.SetFire("中").SetPressure("正常").Run("30分钟")
}

// NewCookCongeeCommand 煮粥
func NewCookCongeeCommand(electricCooker *ElectricCooker) *cookCongeeCommand {
	return &cookCongeeCommand{
		electricCooker: electricCooker,
	}
}

func (c *cookCongeeCommand) Execute() string {
	return "煮粥:" + c.electricCooker.SetFire("大").SetPressure("强").Run("45分钟")
}

// NewShutdownCommand 停止
func NewShutdownCommand(electricCooker *ElectricCooker) *shutdownCommand {
	return &shutdownCommand{
		electricCooker: electricCooker,
	}
}

func (s *shutdownCommand) Execute() string {
	return s.electricCooker.Shutdown()
}

// SetCookCommand 设置指令
func (e *ElectricCookerInvoker) SetCookCommand(cookCommand CookCommand) {
	e.cookCommand = cookCommand
}

// ExecuteCookCommand 执行指令
func (e *ElectricCookerInvoker) ExecuteCookCommand() string {
	return e.cookCommand.Execute()
}

// SetFire 设置火力
func (e *ElectricCooker) SetFire(fire string) *ElectricCooker {
	e.fire = fire
	return e
}

// SetPressure 设置压力
func (e *ElectricCooker) SetPressure(pressure string) *ElectricCooker {
	e.pressure = pressure
	return e
}

// Run 持续运行指定时间
func (e *ElectricCooker) Run(duration string) string {
	return fmt.Sprintf("电饭煲设置火力为%s,压力为%s,持续运行%s;", e.fire, e.pressure, duration)
}

// Shutdown 停止
func (e *ElectricCooker) Shutdown() string {
	return "电饭煲停止运行。"
}

func main() {
	// 创建电饭煲，命令接受者
	electricCooker := new(ElectricCooker)
	// 创建电饭煲指令触发器
	electricCookerInvoker := new(ElectricCookerInvoker)
	// 停止
	shutdownCmd := NewShutdownCommand(electricCooker)

	// 蒸饭
	steamRiceCmd := NewSteamRiceCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(steamRiceCmd)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())
	electricCookerInvoker.SetCookCommand(shutdownCmd)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())

	// 煮粥
	cookCongeeCmd := NewCookCongeeCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(cookCongeeCmd)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())
	electricCookerInvoker.SetCookCommand(shutdownCmd)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())
}

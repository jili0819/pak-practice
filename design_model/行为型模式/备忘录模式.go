package main

import "fmt"

// 备忘录模式是一种行为设计模式， 允许在不暴露对象实现细节的情况下保存和恢复对象之前的状态。
// 备忘录不会影响它所处理的对象的内部结构， 也不会影响快照中保存的数据。

// 一般情况由原发对象保存生成的备忘录对象的状态不能被除原发对象之外的对象访问，所以通过内部类定义具体的备忘录对象是比较安全的，
// 但是go语言不支持内部类定义的方式，因此go语言实现备忘录对象时，首先将备忘录保存的状态设为非导出字段，避免外部对象访问，
// 其次将原发对象的引用保存到备忘录对象中，当通过备忘录对象恢复时，直接操作备忘录的恢复方法，将备份数据状态设置到原发对象中，完成恢复。

// 例子：大家平时玩的角色扮演闯关游戏的存档机制就可以通过备忘录模式实现，每到一个关键关卡，玩家经常会先保存游戏存档，用于闯关失败后重置，
// 存档会把角色状态及场景状态保存到备忘录中，同时将需要恢复游戏的引用存入备忘录，用于关卡重置；
// 每个阶段都会引用最新的数据，当需要reload时，加载map中tag的信息到引用，即可达到修改引用信息，即reload

type (
	// Originator 备忘录模式原发器接口
	Originator interface {
		Save(tag string) Memento // 当前状态保存备忘录
	}

	// RolesPlayGame 支持存档的RPG游戏
	RolesPlayGame struct {
		name          string   // 游戏名称
		rolesState    []string // 游戏角色状态
		scenarioState string   // 游戏场景状态
	}
)

// NewRolesPlayGame 根据游戏名称和角色名，创建RPG游戏
func NewRolesPlayGame(name string, roleName string) *RolesPlayGame {
	return &RolesPlayGame{
		name:          name,
		rolesState:    []string{roleName, "血量100"}, // 默认满血
		scenarioState: "开始通过第一关",                   // 默认第一关开始
	}
}

// Save 保存RPG游戏角色状态及场景状态到指定标签归档
func (r *RolesPlayGame) Save(tag string) Memento {
	return newRPGArchive(tag, r.rolesState, r.scenarioState, r)
}

func (r *RolesPlayGame) SetRolesState(rolesState []string) {
	r.rolesState = rolesState
}

func (r *RolesPlayGame) SetScenarioState(scenarioState string) {
	r.scenarioState = scenarioState
}

// String 输出RPG游戏简要信息
func (r *RolesPlayGame) String() string {
	return fmt.Sprintf("在%s游戏中，玩家使用%s,%s,%s;", r.name, r.rolesState[0], r.rolesState[1], r.scenarioState)
}

type (
	// Memento 备忘录接口
	Memento interface {
		Tag() string // 备忘录标签
		Restore()    // 根据备忘录存储数据状态恢复原对象
	}

	// rpgArchive rpg游戏存档
	rpgArchive struct {
		tag           string         // 存档标签
		rolesState    []string       // 存档的角色状态
		scenarioState string         // 存档游戏场景状态
		rpg           *RolesPlayGame // rpg游戏引用
	}
)

// newRPGArchive 根据标签，角色状态，场景状态，rpg游戏引用，创建游戏归档备忘录
func newRPGArchive(tag string, rolesState []string, scenarioState string, rpg *RolesPlayGame) *rpgArchive {
	return &rpgArchive{
		tag:           tag,
		rolesState:    rolesState,
		scenarioState: scenarioState,
		rpg:           rpg,
	}
}

func (r *rpgArchive) Tag() string {
	return r.tag
}

// Restore 根据归档数据恢复游戏状态
func (r *rpgArchive) Restore() {
	r.rpg.SetRolesState(r.rolesState)
	r.rpg.SetScenarioState(r.scenarioState)
}

// RPGArchiveManager RPG游戏归档管理器
type RPGArchiveManager struct {
	archives map[string]Memento // 存储归档标签对应归档
}

func NewRPGArchiveManager() *RPGArchiveManager {
	return &RPGArchiveManager{
		archives: make(map[string]Memento),
	}
}

// Reload 根据标签重新加载归档数据
func (r *RPGArchiveManager) Reload(tag string) {
	if archive, ok := r.archives[tag]; ok {
		fmt.Printf("重新加载%s;\n", tag)
		archive.Restore()
	}
}

// Put 保存归档数据
func (r *RPGArchiveManager) Put(memento Memento) {
	r.archives[memento.Tag()] = memento
}

func main() {
	// 创建RPG游戏存档管理器
	rpgManager := NewRPGArchiveManager()
	// 创建RPG游戏
	rpg := NewRolesPlayGame("暗黑破坏神2", "野蛮人战士")
	fmt.Println(rpg)                  // 输出游戏当前状态
	rpgManager.Put(rpg.Save("第一关存档")) // 游戏存档

	// 第一关闯关失败
	rpg.SetRolesState([]string{"野蛮人战士", "成功"})
	rpg.SetScenarioState("第一关闯关成功")
	rpgManager.Put(rpg.Save("第二关存档")) // 游戏存档
	fmt.Println(rpg)

	// 恢复存档，重新闯关
	rpgManager.Reload("第二关存档")
	fmt.Println(rpg)
}

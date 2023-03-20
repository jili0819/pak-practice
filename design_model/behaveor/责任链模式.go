package main

import "fmt"

// (责任链模式)责任链模式是一种行为设计模式， 允许你将请求沿着处理者链进行发送。收到请求后，每个处理者均可对请求进行处理，或将其传递给链上的下个处理者。
//该模式允许多个对象来对请求进行处理，而无需让发送者类与具体接收者类相耦合。链可在运行时由遵循标准处理者接口的任意处理者动态生成。
//一般意义上的责任链模式是说，请求在链上流转时任何一个满足条件的节点处理完请求后就会停止流转并返回，不过还可以根据不同的业务情况做一些改进：
//请求可以流经处理链的所有节点，不同节点会对请求做不同职责的处理；
//可以通过上下文参数保存请求对象及上游节点的处理结果，供下游节点依赖，并进一步处理；
//处理链可支持节点的异步处理，通过实现特定接口判断，是否需要异步处理；
//责任链对于请求处理节点可以设置停止标志位，不是异常，是一种满足业务流转的中断；
//责任链的拼接方式存在两种，一种是节点遍历，一个节点一个节点顺序执行；另一种是节点嵌套，内层节点嵌入在外层节点执行逻辑中，类似递归，或者“回”行结构；
//责任链的节点嵌套拼接方式多被称为拦截器链或者过滤器链，更易于实现业务流程的切面，比如监控业务执行时长，日志输出，权限校验等；

//本示例模拟实现机场登机过程，第一步办理登机牌，第二步如果有行李，就办理托运，第三步核实身份，第四步安全检查，第五步完成登机；
//其中行李托运是可选的，其他步骤必选，必选步骤有任何不满足就终止登机；
//旅客对象作为请求参数上下文，每个步骤会根据旅客对象状态判断是否处理或流转下一个节点；

type (
	// BoardingProcessor 登机过程中，各节点统一处理接口
	BoardingProcessor interface {
		SetNextProcessor(processor BoardingProcessor) BoardingProcessor
		ProcessFor(passenger *Passenger) error
	}

	// Passenger 旅客
	Passenger struct {
		name                  string // 姓名
		hasBoardingPass       bool   // 是否办理登机牌
		hasLuggage            bool   // 是否有行李需要托运
		isPassIdentityCheck   bool   // 是否通过身份校验
		isPassSecurityCheck   bool   // 是否通过安检
		isPassQualityCheck    bool   // 是否通过质检
		isCompleteForBoarding bool   // 是否完成登机
	}

	// baseBoardingProcessor 登机流程处理器基类
	baseBoardingProcessor struct {
		// nextProcessor 下一个登机处理流程
		nextProcessor BoardingProcessor
	}
)

// SetNextProcessor 基类中统一实现设置下一个处理器方法
func (b *baseBoardingProcessor) SetNextProcessor(processor BoardingProcessor) BoardingProcessor {
	b.nextProcessor = processor
	return processor
}

// ProcessFor 基类中统一实现下一个处理器流转
func (b *baseBoardingProcessor) ProcessFor(passenger *Passenger) (err error) {
	if b.nextProcessor != nil {
		if err = b.nextProcessor.ProcessFor(passenger); err != nil {
			return
		}
	}
	return
}

// boardingPassProcessor 办理登机牌处理器
type boardingPassProcessor struct {
	baseBoardingProcessor // 引用基类
}

func (b *boardingPassProcessor) ProcessFor(passenger *Passenger) (err error) {
	fmt.Printf("登记开始：\n")
	if !passenger.hasBoardingPass {
		text := fmt.Sprintf("旅客%s未办理登机牌，开始办理登机牌;\n", passenger.name)
		fmt.Printf(text)
		passenger.hasBoardingPass = true
	}
	// 成功办理登机牌后，进入下一个流程处理
	if err = b.baseBoardingProcessor.ProcessFor(passenger); err != nil {
		return
	}
	return
}

// luggageCheckInProcessor 托运行李处理器
type luggageCheckInProcessor struct {
	baseBoardingProcessor
}

func (l *luggageCheckInProcessor) ProcessFor(passenger *Passenger) (err error) {
	fmt.Printf("行李托运部门：\n")
	if passenger.hasLuggage {
		fmt.Printf("旅客%s有行李，办理行李托运;\n", passenger.name)
	} else {
		fmt.Printf("旅客%s无行李\n", passenger.name)
	}
	if err = l.baseBoardingProcessor.ProcessFor(passenger); err != nil {
		return
	}
	return
}

// identityCheckProcessor 校验身份处理器
type identityCheckProcessor struct {
	baseBoardingProcessor
}

func (i *identityCheckProcessor) ProcessFor(passenger *Passenger) (err error) {
	fmt.Printf("核实身份部门：\n")
	if !passenger.isPassIdentityCheck {
		fmt.Printf("旅客%s未核实身份信息，进行核实;\n", passenger.name)
		passenger.isPassIdentityCheck = true
	}
	if err = i.baseBoardingProcessor.ProcessFor(passenger); err != nil {
		return
	}
	return
}

// securityCheckProcessor 安检处理器
type securityCheckProcessor struct {
	baseBoardingProcessor
}

func (s *securityCheckProcessor) ProcessFor(passenger *Passenger) (err error) {
	fmt.Printf("安检部门：\n")
	if !passenger.isPassSecurityCheck {
		fmt.Printf("旅客%s未安检，进行安检;\n", passenger.name)
		passenger.isPassSecurityCheck = true
	}
	if err = s.baseBoardingProcessor.ProcessFor(passenger); err != nil {
		return
	}
	return
}

// qualityCheckProcessor 质检处理器
type qualityCheckProcessor struct {
	baseBoardingProcessor
}

func (s *qualityCheckProcessor) ProcessFor(passenger *Passenger) (err error) {
	fmt.Printf("质检部门：\n")
	if !passenger.isPassQualityCheck {
		fmt.Printf("旅客%s未质检，进行质检;\n", passenger.name)
		passenger.isPassQualityCheck = true
	}
	if err = s.baseBoardingProcessor.ProcessFor(passenger); err != nil {
		return
	}
	return
}

// completeBoardingProcessor 完成登机处理器
type completeBoardingProcessor struct {
	baseBoardingProcessor
}

func (c *completeBoardingProcessor) ProcessFor(passenger *Passenger) (err error) {
	passenger.isCompleteForBoarding = true
	fmt.Printf("旅客%s成功登机;\n", passenger.name)
	return
}

func BuildPassenger() BoardingProcessor {
	// 登记牌
	boardingPassNode := &boardingPassProcessor{}
	// process start // 新一托运=>身份校验=>安检=>质检=>登机完成
	boardingPassNode.
		SetNextProcessor(&luggageCheckInProcessor{}).
		SetNextProcessor(&identityCheckProcessor{}).
		SetNextProcessor(&securityCheckProcessor{}).
		SetNextProcessor(&qualityCheckProcessor{}).
		SetNextProcessor(&completeBoardingProcessor{})
	return boardingPassNode
}

func main() {
	passenger := BuildPassenger()
	if err := passenger.ProcessFor(&Passenger{
		name:       "张三",
		hasLuggage: true,
	}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}

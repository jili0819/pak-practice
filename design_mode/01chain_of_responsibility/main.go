package main

import "fmt"

func main() {
	passenger := BuildPassenger()
	if err := passenger.ProcessFor(&Passenger{
		name:                  "张三",
		hasBoardingPass:       false,
		hasLuggage:            true,
		isPassIdentityCheck:   false,
		isPassSecurityCheck:   false,
		isCompleteForBoarding: false,
	}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}

package main

import "fmt"

type ApiInterface interface {
	// Add 方法一
	Add(a int64) ApiInterface
	GetSum() int64
}

type addApiClient struct {
	Sum int64 `json:"sum"`
	X   int64 `json:"x"`
	Y   int64 `json:"y"`
}

func NewAddApi() ApiInterface {
	return &addApiClient{}
}

func (api *addApiClient) Add(x int64) ApiInterface {
	api.Sum += x
	return api
}

func (api *addApiClient) GetSum() int64 {
	return api.Sum
}

func main() {
	fmt.Println(NewAddApi().Add(1).Add(3).Add(5).Add(1).GetSum())
}

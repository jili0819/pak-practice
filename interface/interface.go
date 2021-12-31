package _interface

type ApiInterface interface {
	// Add 方法一
	Add(a, b int64) (int64, error)
	// Sums 方法二
	Sums(a, b int64) (int64, error)
}

type abbApiClient struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

func NewApi() ApiInterface {
	return &abbApiClient{}
}

func (api *abbApiClient) Add(a, b int64) (int64, error) {
	return a + b, nil
}

func (api *abbApiClient) Sums(a, b int64) (int64, error) {
	return a + b, nil
}

package main

//
//var PoolLimit chan struct{}
//
//var XC *xc
//
//type xc struct{}
//
///*
//InitXcPond main函数时需初始化
//num 协程限制数量
//*/
//func InitXcPond(num int) {
//	PoolLimit = make(chan struct{}, num)
//	XC = &xc{}
//}
//
//// Go 开始协程  使用上下文来传递函数需要处理的变量
//func (x *xc) Go(f func(c context.Context), ctx context.Context) {
//	fmt.Println(fmt.Sprintf("当前：%s 正在使用协程,总协程数为：%d", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), runtime.NumGoroutine()))
//	f(ctx)
//	<-PoolLimit
//	return
//}
//
//func init() {
//	InitXcPond(25)
//}
//
//func main() {
//	for i := 0; i <= 100; i++ {
//		c := context.Background()
//		c = context.WithValue(c, "myKey", i)
//		PoolLimit <- struct{}{}
//		go XC.Go(myTest, c)
//	}
//}
//
//func myTest(c context.Context) {
//	fmt.Println("key", c.Value("myKey"))
//}

func main() {
	a := A{}
	NewPool(10).AddFunc(a.PoolExecute)
}

type A struct {
	Name string
}

func (a *A) PoolExecute(x interface{}) {

}

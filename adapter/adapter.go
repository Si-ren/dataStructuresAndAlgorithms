package adapter

import "fmt"

// 目标接口
type Target interface {
	Request()
}

// 适配者接口
type Adaptee interface {
	SpecificRequest()
}

// 适配者实现
type ConcreteAdaptee struct{}

func (a *ConcreteAdaptee) SpecificRequest() {
	fmt.Println("适配者的特殊请求")
}

// 适配器实现
// 代理模式和适配器模式在实现上可能会有相似之处，例如两者都使用了接口和组合等概念。
type Adapter struct {
	// 这里可以塞任何实现SpecificRequest()方法的结构体
	Adapt Adaptee
}

func (a *Adapter) Request() {
	a.Adapt.SpecificRequest()
}

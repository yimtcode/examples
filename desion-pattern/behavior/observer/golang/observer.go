package main

import "fmt"

// Observer 观察者接口
type Observer interface {
	Update(state string)
}

// ConcreteObServer 具体的观察者实现
type ConcreteObServer struct {
}

// Update 接收通知
func (c *ConcreteObServer) Update(message string) {
	fmt.Printf("通知: %s\n", message)
}

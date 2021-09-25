package main

func main() {
	// 主题
	var subject Subject = &ConcreteSubject{Message: "有内鬼终止交易"}
	// 观察者
	var observer1 Observer = &ConcreteObServer{}
	var observer2 Observer = &ConcreteObServer{}
	// 添加观察者
	subject.Attach(observer1)
	subject.Attach(observer2)
	// 通知观察者
	subject.Notify()
}

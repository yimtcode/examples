package main

// Subject 主题（通知者）接口
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
}

// ConcreteSubject 具体的主题（通知者）接口
type ConcreteSubject struct {
	Message   string
	observers []Observer
}

// Attach 添加观察者
func (c *ConcreteSubject) Attach(observer Observer) {
	c.observers = append(c.observers, observer)
}

func (c *ConcreteSubject) getIndex(observer Observer) int {
	for i, v := range c.observers {
		if v == observer {
			return i
		}
	}

	return -1
}

// Detach 删除观察者
func (c *ConcreteSubject) Detach(observer Observer) {
	index := c.getIndex(observer)
	if index == -1 {
		return
	}

	c.observers = append(c.observers[:index], c.observers[index+1:]...)
}

// Notify 通知所有观察者
func (c *ConcreteSubject) Notify() {
	for _, v := range c.observers {
		v.Update(c.Message)
	}
}

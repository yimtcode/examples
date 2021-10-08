package array

import "errors"

var (
	ErrIndexOut = errors.New("Index out error ")
)

const (
	// 默认长度
	defaultLength = 8
	// 扩容，当空间不够时扩容比率
	rate = 2
)

// Array go实现数组
type Array struct {
	data   []interface{}
	length int
}

func NewArray(objs ...interface{}) *Array {
	arr := new(Array)
	for _, obj := range objs {
		arr.Append(obj)
	}
	return arr
}

// 懒初始化
func (a *Array) lazy() {
	if a.data == nil {
		a.data = make([]interface{}, defaultLength, defaultLength)
		a.length = 0
	}
}

// 扩容
func (a *Array) capacityExpansion() {
	capacity := cap(a.data) * rate
	newData := make([]interface{}, capacity, capacity)
	for i, v := range a.data {
		newData[i] = v
	}
	a.data = newData
}

// Append 增
func (a *Array) Append(obj interface{}) {
	a.lazy()
	if a.length == cap(a.data) {
		// 扩容
		a.capacityExpansion()
	}
	a.data[a.length] = obj
	a.length++
}

// Delete 删除
func (a *Array) Delete(index int) {
	if a.length <= index {
		panic(ErrIndexOut)
	}
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
}

// Set 改
func (a *Array) Set(index int, obj interface{}) {
	if a.length <= index {
		panic(ErrIndexOut)
	}

	a.data[index] = obj
}

// Get 查
func (a *Array) Get(index int) interface{} {
	if a.length <= index {
		panic(ErrIndexOut)
	}

	return a.data[index]
}

func (a *Array) Len() int {
	return a.length
}

func (a *Array) Cap() int {
	return cap(a.data)
}

func (a *Array) ForEach(cb func(index int, obj interface{}) bool) {
	for i := 0; i < a.length; i++ {
		if !cb(i, a.data[i]) {
			break
		}
	}
}

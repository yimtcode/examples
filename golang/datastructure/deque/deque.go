package deque

type Deque struct {
	m     map[int64]interface{}
	first int64
	last  int64
}

func New() *Deque {
	d := &Deque{}
	d.lazy()
	return d
}

func (d *Deque) PushLeft(obj interface{}) {
	d.lazy()

	d.first++
	d.m[d.first] = obj
}

func (d *Deque) PushRight(obj interface{}) {
	d.lazy()

	d.last++
	d.m[d.last] = obj
}

func (d *Deque) PopLeft() interface{} {
	d.lazy()

	if d.Len() == 0 {
		panic("Deque is empty")
	}

	obj := d.m[d.first]
	d.first--
	return obj
}

func (d *Deque) PopRight() interface{} {
	d.lazy()

	if d.Len() == 0 {
		panic("Deque is empty")
	}

	obj := d.m[d.last]
	d.last--
	return obj
}

func (d *Deque) Len() int {
	d.lazy()

	return int(d.last - d.first + 1)
}

func (d *Deque) lazy() {
	if d.m == nil {
		d.m = make(map[int64]interface{})
		d.last = -1
		d.first = 0
	}
}

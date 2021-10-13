package deque

import (
	"testing"
)

func TestDeque_Len(t *testing.T) {
	d := New()
	d.PushRight(1)
	d.PushRight(2)

	length := 2
	if d.Len() != length {
		t.Errorf("Len() == %d, want %d", d.Len(), length)
	}
}

func TestDeque_PopLeft(t *testing.T) {
	one := 1
	two := 2
	d := New()
	d.PushRight(one)
	d.PushRight(two)
	obj := d.PopLeft()

	if obj != one {
		t.Errorf("PopLeft() == %d, want %d", obj, one)
	}
}

func TestDeque_PopRight(t *testing.T) {
	one := 1
	two := 2
	d := New()
	d.PushRight(one)
	d.PushRight(two)
	obj := d.PopRight()

	if obj != two {
		t.Errorf("PopLeft() == %d, want %d", obj, two)
	}
}

func TestDeque_PushLeft(t *testing.T) {
	one := 1
	two := 2
	d := New()
	d.PushLeft(one)
	d.PushLeft(two)

	obj := d.PopLeft()
	if obj != two {
		t.Errorf("PopLeft() == %d, want %d", obj, two)
	}
}

func TestDeque_PushRight(t *testing.T) {
	one := 1
	two := 2
	d := New()
	d.PushRight(one)
	d.PushRight(two)

	obj := d.PopRight()
	if obj != two {
		t.Errorf("PopRight() == %d, want %d", obj, two)
	}
}
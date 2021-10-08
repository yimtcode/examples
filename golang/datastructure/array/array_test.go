package array

import (
	"reflect"
	"testing"
)

func TestArray_Append(t *testing.T) {
	type args struct {
		obj []interface{}
	}
	tests := []struct {
		name string
		arr  *Array
		args args
		want int
	}{
		{
			name: "Test append",
			arr:  NewArray(),
			args: args{
				obj: []interface{}{1, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.args.obj {
				tt.arr.Append(v)
			}

			if length := tt.arr.Len(); length != tt.want {
				t.Errorf("Len() => %d, want %d", length, tt.want)
			}
		})
	}
}

func TestArray_Cap(t *testing.T) {
	tests := []struct {
		name string
		arr  *Array
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Cap test",
			arr:  NewArray(1, 2, 3),
			want: defaultLength,
		}, {
			name: "Cap test 2",
			arr:  NewArray(0, 1, 2, 3, 4, 5, 6, 7, 8),
			want: defaultLength * rate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.arr
			if got := a.Cap(); got != tt.want {
				t.Errorf("Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_Delete(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		arr  *Array
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Delete test",
			arr:  NewArray(1, 2, 3),
			args: args{index: 0},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arr.Delete(tt.args.index)
			if length := tt.arr.Len(); length != tt.want {
				t.Errorf("Len() => %d, want %d", length, tt.want)
			}
		})
	}
}

func TestArray_Get(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		arr  *Array
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{
			name: "Get test",
			arr:  NewArray(1, 2, 3),
			args: args{index: 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arr.Get(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_Len(t *testing.T) {
	tests := []struct {
		name string
		arr  *Array
		want int
	}{
		// TODO: Add test cases.
		{
			"Len test",
			NewArray(1, 2, 3, 4),
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.arr.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_Set(t *testing.T) {
	type args struct {
		index int
		obj   interface{}
	}
	tests := []struct {
		name string
		arr  *Array
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Set test",
			arr:  NewArray(1, 2, 3),
			args: args{
				index: 0,
				obj:   99,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arr.Set(tt.args.index, tt.args.obj)

			if obj := tt.arr.Get(tt.args.index); !reflect.DeepEqual(obj, tt.args.obj) {
				t.Errorf("Get() => %v, want %v", obj, tt.args.obj)
			}
		})
	}
}
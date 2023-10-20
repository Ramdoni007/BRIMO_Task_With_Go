package brimotaskwithgolang

import (
	"reflect"
	"testing"
)

func TestPendulum(t *testing.T) {
	type args struct {
		values []int
	}

	testCase := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"1", args{[]int{4, 10, 9}}, []int{10, 4, 9}},
		{"2", args{[]int{-5, -10, -1, -10, -8}}, []int{-1, -8, -10, -10, -5}},
		{"3", args{[]int{6, 6, 8, 5, 10}}, []int{10, 6, 5, 6, 8}},
		{"4", args{[]int{7, -10, -1, -10, -8}}, []int{7, -8, -10, -10, -1}},
		{"5", args{[]int{18, 31, 17, 20, 29, 27, 28, 16, 21, 19}}, []int{29, 27, 20, 18, 16, 17, 19, 21, 28, 31}},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Pendulum(tt.args.values); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Pendulum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

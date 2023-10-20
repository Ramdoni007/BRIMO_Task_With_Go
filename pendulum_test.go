package brimotaskwithgolang

import (
	"reflect"
	"testing"
)

// Dan Disinilah pembuktian test code dari pendulum di lakukan membuat func TestPendulum(t)
func TestPendulum(t *testing.T) {
	// Dengan struct args yang mempunyai field values []int
	type args struct {
		values []int
	}

	// Testcase dengan konsep table testCase dilakukan dengan slice of struct yang mempunyai field name type string, args type struct args, Dan
	// wantRes dengan type slice of int
	testCase := []struct {
		name    string
		args    args
		wantRes []int
	}{
		// Dan disinilah semua contoh pembandingan Konsep Pendulum di lakukan.. "1" artinya >> test ke 1. dengan
		// args(input) bilangan args{[]int {4,10,9}} di masukan... dan []int{10,4,9} adalah hasil ekspektasi keluaran
		// benar dari output Pendulum yang sesuai testing... begitu seterusnya test Ke "2", "3","4" dan "5"
		{"1", args{[]int{4, 10, 9}}, []int{10, 4, 9}},
		{"2", args{[]int{-5, -10, -1, -10, -8}}, []int{-1, -8, -10, -10, -5}},
		{"3", args{[]int{6, 6, 8, 5, 10}}, []int{10, 6, 5, 6, 8}},
		{"4", args{[]int{7, -10, -1, -10, -8}}, []int{7, -8, -10, -10, -1}},
		{"5", args{[]int{18, 31, 17, 20, 29, 27, 28, 16, 21, 19}}, []int{29, 27, 20, 18, 16, 17, 19, 21, 28, 31}},
	}
	// Lakukan iterasi testCase di atas dengan for range dan method t.Run
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			// panggil pengkodisian dan function Pendulum yang sebelumnya berhasil di buat
			// Dan saya disini Ramdoni :) menggunakan reflect package reflect.DeepEqual function untuk melakukan
			// Pengecekan dan mengbandingkan hasil antara gotRes dan tt.WantRes ketika tidak sama maka akan di kembalikan
			// error nya :)
			if gotRes := Pendulum(tt.args.values); !reflect.DeepEqual(gotRes, tt.wantRes) {
				// Errorf merupakan method untuk menampilkan error jika hasil Pendulum tidak sesuai
				t.Errorf("Pendulum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

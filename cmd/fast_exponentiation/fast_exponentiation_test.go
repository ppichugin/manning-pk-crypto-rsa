package main

import "testing"

func Test_fastExp(t *testing.T) {
	type args struct {
		num int
		pow int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"fastExp(8,6)", args{8, 6}, 262144},
		{"fastExp(7,10)", args{7, 10}, 282475249},
		{"fastExp(9,13)", args{9, 13}, 2541865828329},
		{"fastExp(213,5)", args{213, 5}, 438427732293},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fastExp(tt.args.num, tt.args.pow); got != tt.want {
				t.Errorf("fastExp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fastExpMod(t *testing.T) {
	type args struct {
		num int
		pow int
		mod int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"fastExpMod(8, 6, 10)", args{8, 6, 10}, 4},
		{"fastExpMod(7, 10, 101)", args{7, 10, 101}, 65},
		{"fastExpMod(9, 13, 283)", args{9, 13, 283}, 179},
		{"fastExpMod(213, 5, 1000)", args{213, 5, 1000}, 293},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fastExpMod(tt.args.num, tt.args.pow, tt.args.mod); got != tt.want {
				t.Errorf("fastExpMod() = %v, want %v", got, tt.want)
			}
		})
	}
}

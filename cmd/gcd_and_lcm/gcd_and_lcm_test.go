package main

import "testing"

func Test_gcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"gcd(270,192)", args{270, 192}, 6},
		{"gcd(270,-192)", args{270, -192}, 6},
		{"gcd(-270,192)", args{-270, 192}, 6},
		{"gcd(-270,-192)", args{-270, -192}, 6},
		{"gcd(7469,2464)", args{7469, 2464}, 77},
		{"gcd(55290,115430)", args{55290, 115430}, 970},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lcm(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"lcm(270,192)", args{270, 192}, 8640},
		{"lcm(-270,192)", args{-270, 192}, 8640},
		{"lcm(270,-192)", args{270, -192}, 8640},
		{"lcm(-270,-192)", args{-270, -192}, 8640},
		{"lcm(7469,2464)", args{7469, 2464}, 239008},
		{"lcm(55290,115430)", args{55290, 115430}, 6579510},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcm(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("lcm() = %v, want %v", got, tt.want)
			}
		})
	}
}

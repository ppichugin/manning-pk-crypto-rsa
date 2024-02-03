package main

import (
	"reflect"
	"testing"
)

var (
	rawSliceOfTen = []bool{
		true, true, true, true, false, true, false, true, false, false, false, // 0-10(+1)
	}
	readySliceOfTenPrimes = []int{2, 3, 5, 7}
)

func Test_sieveOfEratosthenes(t *testing.T) {
	type args struct {
		max int
	}

	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "10",
			args: args{max: 10},
			want: rawSliceOfTen,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sieveOfEratosthenes(tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sieveOfEratosthenes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sieveToPrimes(t *testing.T) {
	type args struct {
		sieve []bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "10",
			args: args{sieve: rawSliceOfTen},
			want: readySliceOfTenPrimes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sieveToPrimes(tt.args.sieve); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sieveToPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

package quiz

import (
	"reflect"
	"testing"
)

func Test_generatePrimesV1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case1",
			args: args{50},
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePrimesV1(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generatePrimesV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generatePrimesV2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case1",
			args: args{50},
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePrimesV2(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generatePrimesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generatePrimesV3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case1",
			args: args{50},
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePrimesV3(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generatePrimesV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPrimeV1(t *testing.T) {
	t.Run("test is prime ok", func(t *testing.T) {
		for _, n := range []int{2, 3, 5, 7, 11, 13, 17, 19} {
			if isPrimeV1(n) == false {
				t.Errorf("want true, but want false")
			}
		}
	})

	t.Run("test is prime no", func(t *testing.T) {
		for _, n := range []int{1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20} {
			if isPrimeV1(n) == true {
				t.Errorf("want false, but want true")
			}
		}
	})

	t.Run("test is prime no, invalid number", func(t *testing.T) {
		for _, n := range []int{-1} {
			if isPrimeV1(n) == true {
				t.Errorf("want false, but want true")
			}
		}
	})
}

func Test_isPrimeV2(t *testing.T) {
	t.Run("test is prime ok", func(t *testing.T) {
		for _, n := range []int{2, 3, 5, 7, 11, 13, 17, 19} {
			if isPrimeV2(n) == false {
				t.Errorf("isPrimeV2(%v), want true, but want false", n)
			}
		}
	})

	t.Run("test is prime no", func(t *testing.T) {
		for _, n := range []int{1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20} {
			if isPrimeV2(n) == true {
				t.Errorf("isPrimeV2(%v), want false, but want true", n)
			}
		}
	})

	t.Run("test is prime no, invalid number", func(t *testing.T) {
		for _, n := range []int{-1} {
			if isPrimeV2(n) == true {
				t.Errorf("isPrimeV2(%v), want false, but want true", n)
			}
		}
	})
}

func Test_isPrimeV3(t *testing.T) {
	t.Run("test is prime ok", func(t *testing.T) {
		for _, n := range []int{2, 3, 5, 7, 11, 13, 17, 19} {
			if isPrimeV3(n) == false {
				t.Errorf("isPrimeV3(%v), want true, but want false", n)
			}
		}
	})

	t.Run("test is prime no", func(t *testing.T) {
		for _, n := range []int{1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20} {
			if isPrimeV3(n) == true {
				t.Errorf("isPrimeV3(%v), want false, but want true", n)
			}
		}
	})

	t.Run("test is prime no, invalid number", func(t *testing.T) {
		for _, n := range []int{-1} {
			if isPrimeV3(n) == true {
				t.Errorf("isPrimeV3(%v), want false, but want true", n)
			}
		}
	})
}

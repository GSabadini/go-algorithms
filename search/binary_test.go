package search

import "testing"

func TestBinary(t *testing.T) {
	type args struct {
		array  []int
		target int
		low    int
		high   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				array:  []int{4, 2, 1, 5, 6, 25, 23, 12, 15},
				target: 4,
				low:    0,
				high:   9 - 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Binary(tt.args.array, tt.args.target, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("Binary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryNotRegression(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				array:  []int{4, 2, 1, 5, 6, 25, 23, 12, 15},
				target: 25,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryNotRegression(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("BinaryNotRegression() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryNotRegression_2(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				array:  []int{4, 2, 1, 5, 6, 25, 23, 12, 15},
				target: 25,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryNotRegression_2(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("BinaryNotRegression_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBinary(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // this is important because the generateBenchmarkTestCase() is expensive
	for i := 0; i < b.N; i++ {
		_ = Binary(testCase, i, 0, len(testCase)-1)
	}
}

func BenchmarkBinaryNotRegression(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // this is important because the generateBenchmarkTestCase() is expensive
	for i := 0; i < b.N; i++ {
		_ = BinaryNotRegression(testCase, i)
	}
}

func BenchmarkBinaryNotRegression_2(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	b.ResetTimer() // this is important because the generateBenchmarkTestCase() is expensive
	for i := 0; i < b.N; i++ {
		_ = BinaryNotRegression_2(testCase, i)
	}
}

// This function generate consistent testcase for benchmark test.
func generateBenchmarkTestCase() []int {
	var testCase []int
	for i := 0; i < 1000; i++ {
		testCase = append(testCase, i)
	}
	return testCase
}

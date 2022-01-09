package others

import "testing"

func TestDuplicatedInt(t *testing.T) {
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
				array:  []int{1, 2, 3, 23},
				target: 23,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinearSearchDuplicatedInt(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("LinearSearchDuplicatedInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchDuplicatedInt(t *testing.T) {
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
				array:  []int{1, 2, 3, 23},
				target: 23,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchDuplicatedInt(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("BinarySearchDuplicatedInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

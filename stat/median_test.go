package stat

import "testing"

func TestMedian(t *testing.T) {
	type args struct {
		numbers []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test for an odd length of slice",
			args: args{numbers: []float64{12, 43, 54, 29, 95, 84, 74}},
			want: 54,
		},
		{
			name: "Test for an odd length of slice",
			args: args{numbers: []float64{14, 48, 33, 13, 75, 84, 56}},
			want: 48,
		},
		{
			name: "Test for an even length of slice",
			args: args{numbers: []float64{12, 43, 56, 37, 29, 95, 84, 74}},
			want: 49.5,
		},
		{
			name: "Test for an even length of slice",
			args: args{numbers: []float64{14.5, 22.75, 38.67, 23, 35.27, 24.78, 22.56, 28.74}},
			want: 23.89,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Median(tt.args.numbers); got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}

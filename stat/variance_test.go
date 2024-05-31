package stat

import "testing"

func TestVariance(t *testing.T) {
	type args struct {
		numbers []float64
		mean    float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test for an odd length of slice",
			args: args{
				numbers: []float64{12, 43, 54, 29, 95, 84, 74},
				mean:    Average([]float64{12, 43, 54, 29, 95, 84, 74}),
			},
			want: 780.9795918367346,
		},
		{
			name: "Test for an odd length of slice",
			args: args{
				numbers: []float64{14, 48, 33, 13, 75, 84, 56},
				mean:    Average([]float64{14, 48, 33, 13, 75, 84, 56}),
			},
			want: 667.265306122449,
		},
		{
			name: "Test for an even length of slice",
			args: args{
				numbers: []float64{12, 43, 56, 37, 29, 95, 84, 74},
				mean:    Average([]float64{12, 43, 56, 37, 29, 95, 84, 74}),
			},
			want: 722.9375,
		},
		{
			name: "Test for an even length of slice",
			args: args{
				numbers: []float64{14.5, 22.75, 38.67, 23, 35.27, 24.78, 22.56, 28.74},
				mean:    Average([]float64{14.5, 22.75, 38.67, 23, 35.27, 24.78, 22.56, 28.74}),
			},
			want: 52.30747343750001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Variance(tt.args.numbers, tt.args.mean); got != tt.want {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}

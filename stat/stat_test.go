package stat

import (
	"testing"
)

func TestAverage(t *testing.T) {
	type args struct {
		numbers []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test for integers only:",
			args: args{numbers: []float64{12, 43, 55, 29, 95, 84}},
			want: 53,
		},
		{
			name: "Test for integers only:",
			args: args{numbers: []float64{10, 25, 38, 23, 35, 23, 21}},
			want: 25,
		},
		{
			name: "Test for integers only:",
			args: args{numbers: []float64{12, 43, 56, 37, 29, 95, 84, 74}},
			want: 53.75,
		},
		{
			name: "Test for floating point numbers:",
			args: args{numbers: []float64{14.5, 22.75, 38.67, 23, 35.27, 24.78, 22.56}},
			want: 25.93285714285714,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.numbers); got != tt.want {
				t.Errorf("mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			got := Median(tt.args.numbers)
			if got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestStandDev(t *testing.T) {
	type args struct {
		variance float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test for an odd length of slice",
			args: args{
				variance: Variance([]float64{12, 43, 54, 29, 95, 84, 74}, Average([]float64{12, 43, 54, 29, 95, 84, 74})),
			},
			want: 27.946012091830465,
		},
		{
			name: "Test for an odd length of slice",
			args: args{
				variance: Variance([]float64{14, 48, 33, 13, 75, 84, 56}, Average([]float64{14, 48, 33, 13, 75, 84, 56})),
			},
			want: 25.8314789766759,
		},
		{
			name: "Test for an even length of slice",
			args: args{
				variance: Variance([]float64{12, 43, 56, 37, 29, 95, 84, 74}, Average([]float64{12, 43, 56, 37, 29, 95, 84, 74})),
			},
			want: 26.887497094374552,
		},
		{
			name: "Test for an even length of slice",
			args: args{
				variance: Variance([]float64{14.5, 22.75, 38.67, 23, 35.27, 24.78, 22.56, 28.74}, Average([]float64{14.5, 22.75, 38.67, 23, 35.27, 24.78, 22.56, 28.74})),
			},
			want: 7.232390575563519,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandDev(tt.args.variance); got != tt.want {
				t.Errorf("StandDev() = %v, want %v", got, tt.want)
			}
		})
	}
}

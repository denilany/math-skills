package stat

import "testing"

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

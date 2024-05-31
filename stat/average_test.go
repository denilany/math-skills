package stat

import "testing"

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
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

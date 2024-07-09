package asm

import (
	"slices"
	"testing"
)

func TestSum32(t *testing.T) {
	tests := []struct {
		name string
		args []float32
		want float32
	}{
		{
			name: "1",
			args: []float32{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			name: "1",
			args: []float32{1.6, 2.4, 3.9, 4.1},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum32(tt.args); got != tt.want {
				t.Errorf("Sum32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum64(t *testing.T) {

	tests := []struct {
		name string
		args []float64
		want float64
	}{
		{
			name: "1",
			args: []float64{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			name: "2",
			args: []float64{1.6, 2.4, 3.9, 4.1},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum64(tt.args); got != tt.want {
				t.Errorf("Sum64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMul32(t *testing.T) {
	type args struct {
		x []float32
		y []float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "1",
			args: args{
				x: []float32{2, 3, 4},
				y: []float32{2, 2, 2},
			},
			want: []float32{4, 6, 8},
		},
		{
			name: "2",
			args: args{
				x: []float32{0.5, 3.2, 4.1},
				y: []float32{2, 2, 2},
			},
			want: []float32{1, 6.4, 8.2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mul32(tt.args.x, tt.args.y)
			if !slices.Equal(tt.args.x, tt.want) {
				t.Errorf("Mul32() = %v, want %v", tt.args.x, tt.want)
			}
		})
	}
}

func TestMul64(t *testing.T) {
	type args struct {
		x []float64
		y []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "1",
			args: args{
				x: []float64{2, 3, 4},
				y: []float64{2, 2, 2},
			},
			want: []float64{4, 6, 8},
		},
		{
			name: "2",
			args: args{
				x: []float64{0.5, 3.2, 4.1},
				y: []float64{2, 2, 2},
			},
			want: []float64{1, 6.4, 8.2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mul64(tt.args.x, tt.args.y)
			if !slices.Equal(tt.args.x, tt.want) {
				t.Errorf("Mul64() = %v, want %v", tt.args.x, tt.want)
			}
		})
	}
}

func TestMulNum32(t *testing.T) {
	type args struct {
		x []float32
		a float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "1",
			args: args{
				x: []float32{2, 3, 4},
				a: 3,
			},
			want: []float32{6, 9, 12},
		},
		{
			name: "2",
			args: args{
				x: []float32{2, 3, 4},
				a: 0.5,
			},
			want: []float32{1, 1.5, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MulNum32(tt.args.x, tt.args.a)
			if !slices.Equal(tt.args.x, tt.want) {
				t.Errorf("MulNum32() = %v, want %v", tt.args.x, tt.want)
			}
		})
	}
}

func TestMulNum64(t *testing.T) {
	type args struct {
		x []float64
		a float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "1",
			args: args{
				x: []float64{2, 3, 4},
				a: 3,
			},
			want: []float64{6, 9, 12},
		},
		{
			name: "2",
			args: args{
				x: []float64{2, 3, 4},
				a: 0.5,
			},
			want: []float64{1, 1.5, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MulNum64(tt.args.x, tt.args.a)
			if !slices.Equal(tt.args.x, tt.want) {
				t.Errorf("MulNum64() = %v, want %v", tt.args.x, tt.want)
			}
		})
	}
}

func TestAdd32(t *testing.T) {
	type args struct {
		x []float32
		y []float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "1",
			args: args{
				x: []float32{2, 3, 4},
				y: []float32{2, 3, 4},
			},
			want: []float32{4, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Add32(tt.args.x, tt.args.y)
			if !slices.Equal(tt.args.x, tt.want) {
				t.Errorf("Add32() = %v, want %v", tt.args.x, tt.want)
			}
		})
	}
}

func TestAdd64(t *testing.T) {
	type args struct {
		x []float64
		y []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "1",
			args: args{
				x: []float64{2, 3, 4},
				y: []float64{2, 3, 4},
			},
			want: []float64{4, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Add64(tt.args.x, tt.args.y)
			if !slices.Equal(tt.args.x, tt.want) {
				t.Errorf("Add64() = %v, want %v", tt.args.x, tt.want)
			}
		})
	}
}

func TestAddNum32(t *testing.T) {
	type args struct {
		x []float32
		a float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "1",
			args: args{
				x: []float32{2, 3, 4},
				a: 2,
			},
			want: []float32{4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddNum32(tt.args.x, tt.args.a)
			if !slices.Equal(tt.args.x, tt.want) {
				t.Errorf("AddNum32() = %v, want %v", tt.args.x, tt.want)
			}
		})
	}
}

func TestAddNum64(t *testing.T) {
	type args struct {
		x []float64
		a float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "1",
			args: args{
				x: []float64{2, 3, 4},
				a: 2,
			},
			want: []float64{4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddNum64(tt.args.x, tt.args.a)
			if !slices.Equal(tt.args.x, tt.want) {
				t.Errorf("AddNum64() = %v, want %v", tt.args.x, tt.want)
			}
		})
	}
}

func TestMaximumNum32(t *testing.T) {
	type args struct {
		x []float32
		a float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "1",
			args: args{
				x: []float32{2, -3, 4, -10},
				a: 0,
			},
			want: []float32{2, 0, 4, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MaximumNum32(tt.args.x, tt.args.a)
			if !slices.Equal(tt.args.x, tt.want) {
				t.Errorf("MaximumNum32() = %v, want %v", tt.args.x, tt.want)
			}
		})
	}
}

func TestMaximumNum64(t *testing.T) {
	type args struct {
		x []float64
		a float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "1",
			args: args{
				x: []float64{2, -3, 4, -10},
				a: 0,
			},
			want: []float64{2, 0, 4, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MaximumNum64(tt.args.x, tt.args.a)
			if !slices.Equal(tt.args.x, tt.want) {
				t.Errorf("MaximumNum64() = %v, want %v", tt.args.x, tt.want)
			}
		})
	}
}

func TestMul64Simd(t *testing.T) {
	type args struct {
		out []float64
		x   []float64
		y   []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "1",
			args: args{
				out: []float64{0, 0, 0, 0, 0},
				x:   []float64{2, 3, 4, 1},
				y:   []float64{2, 2, 2, 3},
			},
			want: []float64{4, 6, 8, 3, 0},
		},

		{
			name: "2",
			args: args{
				out: []float64{0, 0, 0, 0},
				x:   []float64{5, 3, 3, 10},
				y:   []float64{2, 2, 7, 3},
			},
			want: []float64{10, 6, 21, 30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mul64Simd(tt.args.out, tt.args.x, tt.args.y)
			if !slices.Equal(tt.args.out, tt.want) {
				t.Errorf("Mul64Simd() = %v, want %v", tt.args.out, tt.want)
			}
		})
	}
}

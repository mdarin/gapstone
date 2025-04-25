package gapstone

import (
	"bytes"
	"testing"
)

func TestCompareNormalized(t *testing.T) {
	type args struct {
		a any
		b any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "string newline string",
			args: args{a: "hello\nworld", b: "hello world"},
			want: true,
		},
		{
			name: "byte string",
			args: args{a: []byte("go\tlang"), b: "golang"},
			want: true,
		},
		{
			name: "whitespace string string",
			args: args{a: "  data  ", b: "data"},
			want: true,
		},
		{
			name: "string string not equal",
			args: args{a: "one", b: "two"},
			want: false,
		},
		{
			name: "bytes buffer string",
			args: args{a: bytes.NewBufferString(" data "), b: "data"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareNormalized(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("CompareNormalized() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrettyPrintHex(t *testing.T) {
	type args struct {
		num any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0xc",
			args: args{num: 0xc},
			want: "0xc",
		},
		{
			name: "0x6",
			args: args{num: 0x6},
			want: "0x6",
		},
		// * Cannot use 0xfffffffffffffdd0 (untyped int constant 18446744073709551056) as int value in struct literal (overflows)compilerNumericOverflow
		// * without uint64(v) type convertion
		{
			name: "0xfffffffffffffdd0",
			args: args{num: uint64(0xfffffffffffffdd0)},
			want: "0xfffffffffffffdd0",
		},
		{
			name: "0xff",
			args: args{num: 255},
			want: "0xff",
		},
		{
			name: "0xffffffffffffffff (полное представление)",
			args: args{num: int64(-1)},
			want: "0xffffffffffffffff",
		},
		{
			name: "'test' (неподдерживаемый тип)",
			args: args{num: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrettyPrintHex(tt.args.num); got != tt.want {
				t.Errorf("PrettyPrintHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

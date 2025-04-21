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

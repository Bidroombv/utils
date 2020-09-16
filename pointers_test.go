package utils

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestStringPtr(t *testing.T) {
	var newString string
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "empty string should return nil",
			args: args{
				s: "",
			},
			want: nil,
		},
		{
			name: "declared and not initialized string should return nil",
			args: args{
				s: newString,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringPtr(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func NullPointerIsEmptyString(t *testing.T) {
	var s *string = nil
	assert.Equal(t, PtrString(s), "")
}

func StringPointerIsString(t *testing.T) {
	foo := "Foo"
	assert.Equal(t, PtrString(&foo), "Foo")
}

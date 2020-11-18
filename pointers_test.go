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

func TestIntPtr(t *testing.T) {
	intPtr := 2
	zeroPtr := 0
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		{
			name: "should return pointer to 2 for 2 value",
			args: args{
				i: 2,
			},
			want: &intPtr,
		},
		{
			name: "should return pointer to 0 for 0 value",
			args: args{
				i: 0,
			},
			want: &zeroPtr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntPtr(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPtrInt(t *testing.T) {
	intPtr := 2
	zeroPtr := 0
	type args struct {
		i *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 2 value for pointer to 2",
			args: args{
				i: &intPtr,
			},
			want: 2,
		},
		{
			name: "should return 0 value for pointer to 0",
			args: args{
				i: &zeroPtr,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PtrInt(tt.args.i)
			if err != nil {
				t.Errorf("err should be nil, err = %v, want %v", err, tt.want)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PtrInt() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("should return an error for nil", func(t *testing.T) {
		got, err := PtrInt(nil)
		if !reflect.DeepEqual(err, errIntCannotBeNil) {
			t.Errorf("PtrInt() = %v, want the err: %v", err, errIntCannotBeNil)
		}
		if !reflect.DeepEqual(got, 0) {
			t.Errorf("PtrInt() = %v, want %v", got, 0)
		}
	})
}

func TestNullPointerIsEmptyString(t *testing.T) {
	var s *string = nil
	assert.Equal(t, PtrString(s), "")
}

func TestStringPointerIsString(t *testing.T) {
	foo := "Foo"
	assert.Equal(t, PtrString(&foo), "Foo")
}

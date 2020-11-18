package utils

import "errors"

func StringPtr(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}

func PtrString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func IntPtr(i int) *int {
	return &i
}

var errIntCannotBeNil = errors.New("int cannot be nil")

func PtrInt(i *int) (int, error) {
	if i == nil {
		return 0, errIntCannotBeNil
	}

	return *i, nil
}

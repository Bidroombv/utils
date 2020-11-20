package utils

import "errors"

// Deprecated, will be deleted according to the plan https://bidroom.atlassian.net/wiki/spaces/CP/pages/1386971179/Golang+PtrType+TypePtr+conversions
func StringPtr(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}

// Deprecated, will be deleted according to the plan https://bidroom.atlassian.net/wiki/spaces/CP/pages/1386971179/Golang+PtrType+TypePtr+conversions
func PtrString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// Deprecated, will be deleted according to the plan https://bidroom.atlassian.net/wiki/spaces/CP/pages/1386971179/Golang+PtrType+TypePtr+conversions
func IntPtr(i int) *int {
	return &i
}

var errIntCannotBeNil = errors.New("int cannot be nil")

// Deprecated, will be deleted according to the plan https://bidroom.atlassian.net/wiki/spaces/CP/pages/1386971179/Golang+PtrType+TypePtr+conversions
func PtrInt(i *int) (int, error) {
	if i == nil {
		return 0, errIntCannotBeNil
	}

	return *i, nil
}

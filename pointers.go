package utils

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

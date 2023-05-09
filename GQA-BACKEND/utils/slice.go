package utils

func RemoveDuplicateElementFromSlice(ss []string) []string {
	result := make([]string, 0, len(ss))
	temp := map[string]struct{}{}
	for _, item := range ss {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func CompareStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (b == nil) || (a == nil) {
		return false
	}
	for key, value := range a {
		if value != b[key] {
			return false
		}
	}
	return true
}

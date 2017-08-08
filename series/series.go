package series

const testVersion = 2

func All(n int, s string) []string {
	if len(s) < n {
		return nil
	}

	out := []string{}

	for i := 0; i <= len(s)-n; i++ {
		out = append(out, s[i:i+n])
	}
	return out
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		ok = false
	} else {
		first = s[0:n]
		ok = true
	}
	return
}

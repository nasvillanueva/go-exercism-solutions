package accumulate

const testVersion = 1

func Accumulate(items []string, method func(string) string) []string {
	accItems := make([]string, len(items))
	for index, value := range items {
		accItems[index] = method(value)
	}
	return accItems
}

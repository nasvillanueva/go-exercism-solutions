package letter

const testVersion = 1

func ConcurrentFrequency(in []string) FreqMap {
	freqChan := make(chan FreqMap, len(in))

	for _, val := range in {
		go func(str string) {
			freqChan <- Frequency(str)
		}(val)
	}

	output := FreqMap{}

	for i := 0; i < len(in); i++ {
		for key, val := range <-freqChan {
			output[key] += val
		}
	}

	return output
}

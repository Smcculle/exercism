package letter

func mergeMaps(a, b map[rune]int) map[rune]int {
	for k, v := range b {
		a[k] += v
	}

	return a
}

func ConcurrentFrequency(s []string) FreqMap {
	result := FreqMap{}
	mapChan := make(chan FreqMap)
	for _, text := range s {
		go func(t string) {
			mapChan <- Frequency(t)
		}(text)
	}

	for range s {
		result = mergeMaps(result, <-mapChan)
	}

	return result
}

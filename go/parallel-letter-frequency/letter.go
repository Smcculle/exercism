package letter

import "sync"

var lock sync.RWMutex

func mergeMaps(a, b map[rune]int) map[rune]int {
	lock.Lock()
	defer lock.Unlock()
	for k, v := range b {
		a[k] += v
	}

	return a
}

func ConcurrentFrequency(s []string) FreqMap {
	result := FreqMap{}
	mapChan := make(chan FreqMap, 3)
	for _, text := range s {
		go func(t string) {
			mapChan <- Frequency(t)
		}(text)
	}

	for range s {
		go mergeMaps(result, <-mapChan)
	}

	return result
}

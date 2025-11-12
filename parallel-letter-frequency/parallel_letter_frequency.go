package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	var result FreqMap = FreqMap{}
	ch := make(chan FreqMap)

	for _, text := range texts {
		go func(t string) {
			ch <- Frequency(t)
		}(text)
	}

	for range texts {
		freqMap := <-ch
		for r, count := range freqMap {
			result[r] += count
		}
	}
	return result
}

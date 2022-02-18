package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in a given text and
// returns this data as a FreqMap.
func ConcurrentFrequency(input []string) FreqMap {
	m := FreqMap{}

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i, s := range input {
		wg.Add(1)
		go func(s string, i int) {
			newMap := Frequency(s)
			mutex.Lock()
			for k, v := range newMap {
				m[k] += v
			}
			mutex.Unlock()
			wg.Done()
		}(s, i)
	}
	wg.Wait()

	return m
}

/*
benchmark
	goos: darwin
	goarch: amd64
	pkg: letter
	cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
	BenchmarkSequentialFrequency-8   	   61051	     19291 ns/op
	BenchmarkConcurrentFrequency-8   	   57841	     20922 ns/op
*/

package letter

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
	m, ch := FreqMap{}, make(chan FreqMap)

	for _, s := range input {
		go func(s string) {
			ch <- Frequency(s)
		}(s)
	}

	for range input {
		for k, v := range <-ch {
			m[k] += v
		}
	}

	return m
}

/*
benchmark
	goos: darwin
	goarch: amd64
	pkg: letter
	cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
	BenchmarkSequentialFrequency-8   	   61509	     19434 ns/op
	BenchmarkConcurrentFrequency-8   	   60910	     19245 ns/op
*/

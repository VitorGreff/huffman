package frequency

func BuildFrequencyMap(content string) map[string]int {
	var m = make(map[string]int)
	for _, byte := range content {
		m[string(byte)]++
	}
	return m
}

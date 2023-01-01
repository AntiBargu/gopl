package same // import "github.com/AntiBargu/gopl/ch3/same"

func Same(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m1, m2 := make(map[rune]int), make(map[rune]int)

	for _, c := range s1 {
		m1[c]++
	}

	for _, c := range s2 {
		m2[c]++
	}

	if len(m1) != len(m2) {
		return false
	}

	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}

	return true
}

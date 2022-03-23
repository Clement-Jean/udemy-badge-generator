package main

func calculateTextLength110(s string) float64 {
	if len(s) == 0 {
		return 0.0
	}

	total := 0.0
	for i := 0; i < len(s); i++ {
		total += characterLengths[string(s[i])].(float64)
	}

	for i := 1; i < len(s); i++ {
		pair := s[i-1 : i+1]

		value, containsKey := kerningPairs[pair].(float64)
		if containsKey {
			total -= value
		}
	}
	return float64(total)
}

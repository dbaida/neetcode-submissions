func romanToInt(s string) int {
	substractable := map[byte]byte{
		'V': 'I',
		'X': 'I',
		'L': 'X',
		'C': 'X',
		'D': 'C',
		'M': 'C',
	}
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	prevRoman := s[0]
	prevInt := values[prevRoman]
	result := prevInt
	
	for i := 1; i < len(s); i++ {
		substractableBy, isSubstractable := substractable[s[i]]
		if !isSubstractable || prevRoman != substractableBy {
			result += values[s[i]]
		} else {
			result += values[s[i]] - prevInt * 2
		}
		prevRoman = s[i]
		prevInt = values[prevRoman]
	}
	
	return result
}

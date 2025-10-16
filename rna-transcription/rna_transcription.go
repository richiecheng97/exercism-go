package strand

var mp = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	var res []rune
	for _, v := range dna {
		res = append(res, mp[v])
	}
	return string(res)
}

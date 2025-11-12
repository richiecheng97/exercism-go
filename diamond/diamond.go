package diamond

import "errors"

func Gen(char byte) (string, error) {
	num := int(char - 'A')
	if num < 0 || num > 25 {
		return "", errors.New("")
	}

	size := num*2 + 1
	rows := make([][]rune, size)
	for i := range rows {
		rows[i] = make([]rune, size)
		for j := range rows[i] {
			rows[i][j] = ' '
		}
	}

	for i := 0; i <= num; i++ {
		rows[i][num-i] = rune('A' + i)
		rows[i][num+i] = rune('A' + i)
		rows[size-1-i][num-i] = rune('A' + i)
		rows[size-1-i][num+i] = rune('A' + i)
	}

	result := ""
	for i, row := range rows {
		result += string(row)
		if i < size-1 {
			result += "\n"
		}
	}

	return result, nil

}

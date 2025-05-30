package main

func reverseWords(s string) string {
	b := []byte(s)

	for i, j := 0, 0; i < len(b); i = j + 1 {
		for i < len(b) && b[i] == ' ' {
			i++
		}
		j = i + 1
		for j < len(b) && b[j] != ' ' {
			j++
		}
		for k := j - 1; i < k; i, k = i+1, k-1 {
			b[i], b[k] = b[k], b[i]
		}
	}

	return string(b)
}

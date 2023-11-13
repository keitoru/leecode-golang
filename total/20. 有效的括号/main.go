package main

func isValid(s string) bool {
	var stack []byte

	m := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	if len(s) > 0 {
		if _, ok := m[[]byte(s)[0]]; !ok {
			return false
		}
	}

	for _, c := range []byte(s) {
		if _, ok := m[c]; ok {
			stack = append(stack, c)
			continue
		}

		if len(stack) > 0 {
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if right, _ := m[left]; c != right {
				return false
			}
		} else {
			return false
		}
	}

	return len(stack) == 0
}

package stack

/*
Input {'key1': 'value1', 'key2': [1, 2, 3], 'key3': (1, 2, 3)}  Output True
Input [{'key1': 'value1', 'key2': [1, 2, 3], 'key3': (1, 2, 3)} Output False
*/

func inValue(m map[string]string, t string) bool {
	for _, v := range m {
		if v == t {
			return true
		}
	}
	return false
}

func ValidateFormat(chars string) bool {
	lookup := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	s := NewStack[string]()

	for _, c := range chars {
		if l, ok := lookup[string(c)]; ok {
			s.Push(l)
			continue
		}

		if !inValue(lookup, string(c)) {
			continue
		}

		p := s.Peek()
		if p.Valid && p.Val == string(c) {
			s.Pop()
			continue
		} else {
			return false
		}
	}

	if s.Len() != 0 {
		return false
	}

	return true
}

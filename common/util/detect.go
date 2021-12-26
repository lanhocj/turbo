package util

func empty(o interface{}) bool {
	switch s := o.(type) {
	case string:
		return len(s) == 0
	default:
		return false
	}
}

func Empty(o interface{}) bool {
	return empty(o)
}

func MustNotEmpty(obj ...interface{}) bool {
	for _, o := range obj {
		if empty(o) {
			return false
		}
	}

	return true
}

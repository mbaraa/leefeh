package db

// checkConds reports whether the provided conditions are valid or not
func checkConds(conds ...any) bool {
	return len(conds) > 1 && checkCondsMeaning(conds...)
}

func checkCondsMeaning(conds ...any) bool {
	ok := false

	switch conds[0].(type) {
	case string:
		ok = true
	default:
		return false
	}

	for _, cond := range conds[1:] {
		switch cond.(type) {
		case bool,
			int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64,
			float32, float64,
			complex64, complex128,
			string:
			ok = true
		default:
			return false
		}
	}

	return ok
}

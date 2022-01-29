package sqlbuild

// ArgsFromInt64 convert []int64 to []interface{}
func ArgsFromInt64(in []int64) (out []interface{}) {
	if len(in) == 0 {
		return nil
	}
	out = make([]interface{}, len(in))

	var i int
	for i = 0; i < len(in); i++ {
		out[i] = in[i]
	}

	return out
}

// ArgsFromString convert []string to []interface{}
func ArgsFromString(in []string) (out []interface{}) {
	if len(in) == 0 {
		return nil
	}
	out = make([]interface{}, len(in))

	var i int
	for i = 0; i < len(in); i++ {
		out[i] = in[i]
	}

	return out
}

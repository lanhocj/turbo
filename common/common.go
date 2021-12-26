package common

// Must not error
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

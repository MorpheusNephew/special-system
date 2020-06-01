package utils

// PanicIfError terminates the program if an error has occurred
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

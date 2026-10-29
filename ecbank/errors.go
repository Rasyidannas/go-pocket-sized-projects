package ecbank

// ecbankError defines a sentinel error.
type ecbankError string

// ecbankError implments the error interface.
func (e ecbankError) Error() string {
	return string(e)
}

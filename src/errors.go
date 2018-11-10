package main

// respErr holds the custom error message
type respErr struct {
	Status int
	Text   string
}

// Error returns the error text
func (e *respErr) Error() string {
	return e.Text
}

// Code returns the error code
func (e *respErr) Code() int {
	return e.Status
}

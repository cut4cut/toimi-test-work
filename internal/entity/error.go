package entity

import "errors"

var (
	ErrorNegativePrice             error = errors.New("price is negative")
	ErrorTooManyUrls               error = errors.New("too many urls")
	ErrorTooManyNameSymbols        error = errors.New("too many symbols in name")
	ErrorTooManyDescriptionSymbols error = errors.New("too many symbols in description")
)

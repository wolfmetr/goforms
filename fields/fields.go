// Package fields implements form fields for validating and cleaning
// data from http requests.
package fields

type Field interface {
	Clean(string) (interface{}, ValidationError)
}

type ValidationError interface {
	Error() string
}

type BaseField struct {
	Required bool
}

type Defaults map[string]interface{}

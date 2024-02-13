package response

type Json struct {
	Data    any
	Message string
}

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

package utils

type ValidationErrorBag struct {
	Errors map[string][]string `json:"errors"`
}

func (eb *ValidationErrorBag) AddError(field, message string) {
	if eb.Errors == nil {
		eb.Errors = make(map[string][]string)
	}

	eb.Errors[field] = append(eb.Errors[field], message)
}

func (eb *ValidationErrorBag) ContainsErrors() bool {
	return len(eb.Errors) > 0
}

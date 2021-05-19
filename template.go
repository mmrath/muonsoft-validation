package validation

import "strings"

// TemplateParameter is injected into the message while rendering the template.
type TemplateParameter struct {
	// Key is the marker in the string that will be replaced by value.
	// In general, it is recommended to use double curly braces around the key name.
	// Example: {{ keyName }}
	Key string

	// Value is set by constraint when building violation.
	Value string
}

type TemplateParameterList []TemplateParameter

func (params TemplateParameterList) With(parameters ...TemplateParameter) TemplateParameterList {
	return append(params, parameters...)
}

func (params TemplateParameterList) Prepend(parameters ...TemplateParameter) TemplateParameterList {
	return append(parameters, params...)
}

func renderMessage(template string, parameters []TemplateParameter) string {
	message := template

	for _, p := range parameters {
		message = strings.ReplaceAll(message, p.Key, p.Value)
	}

	return message
}

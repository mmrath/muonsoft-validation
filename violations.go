package validation

import (
	"errors"
	"strings"

	"golang.org/x/text/language"
)

// Violation is the abstraction for validator errors. You can use your own implementations on the application
// side to use it for your needs. In order for the validator to generate application violations,
// it is necessary to implement the ViolationFactory interface and inject it into the validator.
// You can do this by using the SetViolationFactory option in the NewValidator constructor.
type Violation interface {
	error

	// GetCode returns violation code - unique, short, and semantic string that can be used to programmatically
	// test for specific violation. All "code" values are defined in the "github.com/muonsoft/validation/code" package
	// and are protected by backward compatibility rules.
	GetCode() string

	// GetMessage returns translated message with injected values from constraint. It can be used to show
	// a description of a violation to the end-user. Possible values for build-in constraints
	// are defined in the "github.com/muonsoft/validation/message" package and can be changed at any time,
	// even in patch versions.
	GetMessage() string

	// GetMessageTemplate returns template for rendering message. Alongside parameters it can be used to
	// render the message on the client-side of the library.
	GetMessageTemplate() string

	// GetParameters returns the map of the template variables and their values provided by the specific constraint.
	GetParameters() map[string]string

	// GetPropertyPath returns a property path that points to the violated property.
	// See PropertyPath description for more info.
	GetPropertyPath() PropertyPath
}

// ViolationFactory is the abstraction that can be used to create custom violations on the application side.
// Use the SetViolationFactory option on the NewValidator constructor to inject your own factory into the validator.
type ViolationFactory interface {
	BuildViolation(code, message string) *ViolationBuilder
	CreateViolation(
		code,
		messageTemplate string,
		pluralCount int,
		parameters map[string]string,
		propertyPath PropertyPath,
		lang language.Tag,
	) Violation
}

// ViolationList is a slice of violations. It is the usual type of error that is returned from a validator.
type ViolationList []Violation

type NewViolationFunc func(
	code,
	messageTemplate string,
	pluralCount int,
	parameters map[string]string,
	propertyPath PropertyPath,
	lang language.Tag,
) Violation

func (f NewViolationFunc) BuildViolation(code, message string) *ViolationBuilder {
	return newViolationBuilder(f, code, message)
}

func (f NewViolationFunc) CreateViolation(
	code,
	messageTemplate string,
	pluralCount int,
	parameters map[string]string,
	propertyPath PropertyPath,
	lang language.Tag,
) Violation {
	return f(code, messageTemplate, pluralCount, parameters, propertyPath, lang)
}

// Error returns a formatted list of errors as a string.
func (violations ViolationList) Error() string {
	if len(violations) == 0 {
		return "the list of violations is empty, it looks like you forgot to use the AsError method somewhere"
	}

	var s strings.Builder
	s.Grow(32 * len(violations))

	for i, v := range violations {
		if i > 0 {
			s.WriteString("; ")
		}
		if iv, ok := v.(*internalViolation); ok {
			iv.writeToBuilder(&s)
		} else {
			s.WriteString(v.Error())
		}
	}

	return s.String()
}

// AppendFromError appends a single violation or a slice of violations into the end of a given slice.
// If an error does not implement the Violation or ViolationList interface, it will return an error itself.
// Otherwise nil will be returned.
//
// Example
//  violations := make(ViolationList, 0)
//  err := violations.AppendFromError(previousError)
//  if err != nil {
//      // this error is not a violation, processing must fail
//      return err
//  }
//  // violations contain appended violations from the previousError and can be processed further
func (violations *ViolationList) AppendFromError(err error) error {
	if violation, ok := UnwrapViolation(err); ok {
		*violations = append(*violations, violation)
	} else if violationList, ok := UnwrapViolationList(err); ok {
		*violations = append(*violations, violationList...)
	} else if err != nil {
		return err
	}

	return nil
}

// AsError converts the list of violations to an error. This method correctly handles cases where
// the list of violations is empty. It returns nil on an empty list, indicating that the validation was successful.
func (violations ViolationList) AsError() error {
	if len(violations) == 0 {
		return nil
	}

	return violations
}

// IsViolation can be used to verify that the error implements the Violation interface.
func IsViolation(err error) bool {
	var violation Violation

	return errors.As(err, &violation)
}

// IsViolationList can be used to verify that the error implements the ViolationList.
func IsViolationList(err error) bool {
	var violations ViolationList

	return errors.As(err, &violations)
}

// UnwrapViolation is a short function to unwrap Violation from the error.
func UnwrapViolation(err error) (Violation, bool) {
	var violation Violation

	as := errors.As(err, &violation)

	return violation, as
}

// UnwrapViolationList is a short function to unwrap ViolationList from the error.
func UnwrapViolationList(err error) (ViolationList, bool) {
	var violations ViolationList

	as := errors.As(err, &violations)

	return violations, as
}

type internalViolation struct {
	Code            string            `json:"code"`
	Message         string            `json:"message"`
	MessageTemplate string            `json:"-"`
	Parameters      map[string]string `json:"-"`
	PropertyPath    PropertyPath      `json:"propertyPath,omitempty"`
}

func (v internalViolation) Error() string {
	var s strings.Builder
	s.Grow(32)
	v.writeToBuilder(&s)

	return s.String()
}

func (v internalViolation) writeToBuilder(s *strings.Builder) {
	s.WriteString("violation")
	if len(v.PropertyPath) > 0 {
		s.WriteString(" at '" + v.PropertyPath.String() + "'")
	}
	s.WriteString(": " + v.Message)
}

func (v internalViolation) GetCode() string {
	return v.Code
}

func (v internalViolation) GetMessage() string {
	return v.Message
}

func (v internalViolation) GetMessageTemplate() string {
	return v.MessageTemplate
}

func (v internalViolation) GetParameters() map[string]string {
	return v.Parameters
}

func (v internalViolation) GetPropertyPath() PropertyPath {
	return v.PropertyPath
}

type internalViolationFactory struct {
	translator *Translator
}

func newViolationFactory(translator *Translator) *internalViolationFactory {
	return &internalViolationFactory{translator: translator}
}

func (factory *internalViolationFactory) CreateViolation(
	code,
	messageTemplate string,
	pluralCount int,
	parameters map[string]string,
	propertyPath PropertyPath,
	lang language.Tag,
) Violation {
	message := factory.translator.translate(lang, messageTemplate, pluralCount)

	return &internalViolation{
		Code:            code,
		Message:         renderMessage(message, parameters),
		MessageTemplate: messageTemplate,
		Parameters:      parameters,
		PropertyPath:    propertyPath,
	}
}

func (factory *internalViolationFactory) BuildViolation(code, message string) *ViolationBuilder {
	return newViolationBuilder(factory, code, message)
}

// ViolationBuilder used to build an internal implementation of Violation interface.
type ViolationBuilder struct {
	code            string
	messageTemplate string
	pluralCount     int
	parameters      map[string]string
	propertyPath    PropertyPath
	language        language.Tag

	violationFactory ViolationFactory
}

func newViolationBuilder(factory ViolationFactory, code, message string) *ViolationBuilder {
	return &ViolationBuilder{
		code:             code,
		messageTemplate:  message,
		violationFactory: factory,
	}
}

// SetParameters sets parameters that can be injected into the violation message.
func (b *ViolationBuilder) SetParameters(parameters map[string]string) *ViolationBuilder {
	b.parameters = parameters

	return b
}

// SetParameter sets one parameter into a parameters map. If the map is nil it creates a new map.
func (b *ViolationBuilder) SetParameter(name, value string) *ViolationBuilder {
	if b.parameters == nil {
		b.parameters = make(map[string]string)
	}
	b.parameters[name] = value

	return b
}

// SetPropertyPath sets a property path of violated attribute.
func (b *ViolationBuilder) SetPropertyPath(path PropertyPath) *ViolationBuilder {
	b.propertyPath = path

	return b
}

// SetPluralCount sets a plural number that will be used for message pluralization during translations.
func (b *ViolationBuilder) SetPluralCount(pluralCount int) *ViolationBuilder {
	b.pluralCount = pluralCount

	return b
}

// SetLanguage sets language that will be used to translate the violation message.
func (b *ViolationBuilder) SetLanguage(tag language.Tag) *ViolationBuilder {
	b.language = tag

	return b
}

// GetViolation creates a new violation with given parameters and returns it.
// Violation is created by calling the CreateViolation method of the ViolationFactory.
func (b *ViolationBuilder) GetViolation() Violation {
	return b.violationFactory.CreateViolation(
		b.code,
		b.messageTemplate,
		b.pluralCount,
		b.parameters,
		b.propertyPath,
		b.language,
	)
}

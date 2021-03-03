package test

import (
	"testing"

	"github.com/muonsoft/validation"
	"github.com/muonsoft/validation/code"
	"github.com/muonsoft/validation/it"
	"github.com/muonsoft/validation/message"
	"github.com/muonsoft/validation/validationtest"
	"github.com/muonsoft/validation/validator"
	"github.com/stretchr/testify/assert"
)

type Product struct {
	Name       string
	Tags       []string
	Components []Component
}

func (p Product) Validate(scope validation.Scope) error {
	return validator.InScope(scope).Validate(
		validation.String(
			&p.Name,
			validation.PropertyName("name"),
			it.IsNotBlank(),
		),
		validation.Iterable(
			p.Tags,
			validation.PropertyName("tags"),
			it.HasMinCount(1),
		),
		validation.Iterable(
			p.Components,
			validation.PropertyName("components"),
			it.HasMinCount(1),
		),
	)
}

type Component struct {
	ID   int
	Name string
	Tags []string
}

func (c Component) Validate(scope validation.Scope) error {
	return validator.InScope(scope).Validate(
		validation.String(
			&c.Name,
			validation.PropertyName("name"),
			it.IsNotBlank(),
		),
		validation.Iterable(
			c.Tags,
			validation.PropertyName("tags"),
			it.HasMinCount(1),
		),
	)
}

func TestValidateValue_WhenStructWithComplexRules_ExpectViolations(t *testing.T) {
	p := Product{
		Name: "",
		Components: []Component{
			{
				ID:   1,
				Name: "",
			},
		},
	}

	err := newValidator(t).ValidateValue(p)

	validationtest.AssertIsViolationList(t, err, func(t *testing.T, violations validation.ViolationList) bool {
		t.Helper()
		if assert.Len(t, violations, 4) {
			assert.Equal(t, code.NotBlank, violations[0].GetCode())
			assert.Equal(t, "name", violations[0].GetPropertyPath().Format())
			assert.Equal(t, code.CountTooFew, violations[1].GetCode())
			assert.Equal(t, "tags", violations[1].GetPropertyPath().Format())
			assert.Equal(t, code.NotBlank, violations[2].GetCode())
			assert.Equal(t, "components[0].name", violations[2].GetPropertyPath().Format())
			assert.Equal(t, code.CountTooFew, violations[3].GetCode())
			assert.Equal(t, "components[0].tags", violations[3].GetPropertyPath().Format())
		}
		return true
	})
}

func TestValidateValue_WhenValidatableString_ExpectValidationExecutedWithPassedOptionsWithoutConstraints(t *testing.T) {
	validatable := mockValidatableString{value: ""}

	err := newValidator(t).ValidateValue(
		validatable,
		validation.PropertyName("top"),
		it.IsNotBlank().Message("ignored"),
	)

	assertHasOneViolation(code.NotBlank, message.NotBlank, "top.value")(t, err)
}

func TestValidateValidatable_WhenValidatableString_ExpectValidationExecutedWithPassedOptionsWithoutConstraints(t *testing.T) {
	validatable := mockValidatableString{value: ""}

	err := newValidator(t).ValidateValidatable(
		validatable,
		validation.PropertyName("top"),
		it.IsNotBlank().Message("ignored"),
	)

	assertHasOneViolation(code.NotBlank, message.NotBlank, "top.value")(t, err)
}

func TestValidateValue_WhenValidatableStruct_ExpectValidationExecutedWithPassedOptionsWithoutConstraints(t *testing.T) {
	validatable := mockValidatableStruct{}

	err := newValidator(t).ValidateValue(
		validatable,
		validation.PropertyName("top"),
		it.IsNotBlank().Message("ignored"),
	)

	validationtest.AssertIsViolationList(t, err, func(t *testing.T, violations validation.ViolationList) bool {
		t.Helper()
		if assert.Len(t, violations, 4) {
			assert.Equal(t, "top.intValue", violations[0].GetPropertyPath().Format())
			assert.Equal(t, "top.floatValue", violations[1].GetPropertyPath().Format())
			assert.Equal(t, "top.stringValue", violations[2].GetPropertyPath().Format())
			assert.Equal(t, "top.structValue.value", violations[3].GetPropertyPath().Format())
		}
		return true
	})
}

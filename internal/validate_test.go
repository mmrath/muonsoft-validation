package internal

import (
	"fmt"
	"testing"

	"github.com/muonsoft/validation"
	"github.com/muonsoft/validation/code"
	"github.com/muonsoft/validation/it"
	"github.com/muonsoft/validation/message"
	"github.com/muonsoft/validation/validationtest"
	"github.com/stretchr/testify/assert"
)

func TestValidate_GivenValueOfType_ValueValidated(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
	}{
		{"bool", false},
		{"int8", int8(0)},
		{"uint8", uint8(0)},
		{"float32", float32(0)},
		{"string", ""},
		{"bool pointer", boolValue(false)},
		{"int64 pointer", intValue(0)},
		{"uint64 pointer", uintValue(0)},
		{"float64 pointer", floatValue(0)},
		{"string pointer", stringValue("")},
		{"bool nil", nilBool},
		{"int64 nil", nilInt},
		{"uint64 nil", nilUint},
		{"float64 nil", nilFloat},
		{"string nil", nilString},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validation.Validate(test.value, validation.PropertyName("property"), it.IsNotBlank())

			validationtest.AssertIsViolationList(t, err, func(t *testing.T, violations validation.ViolationList) bool {
				if assert.Len(t, violations, 1) {
					assertHasOneViolation(code.NotBlank, message.NotBlank, "property")(t, err)
				}
				return true
			})
		})
	}
}

func TestValidate_ValidatableString_ValidationExecutedWithPassedOptionsWithoutConstraints(t *testing.T) {
	validatable := mockValidatableString{value: ""}

	err := validation.Validate(
		validatable,
		validation.PropertyName("top"),
		it.IsNotBlank().Message("ignored"),
	)

	assertHasOneViolation(code.NotBlank, message.NotBlank, "top.value")(t, err)
}

func TestValidate_ValidatableStruct_ValidationExecutedWithPassedOptionsWithoutConstraints(t *testing.T) {
	validatable := mockValidatableStruct{}

	err := validation.Validate(
		validatable,
		validation.PropertyName("top"),
		it.IsNotBlank().Message("ignored"),
	)

	validationtest.AssertIsViolationList(t, err, func(t *testing.T, violations validation.ViolationList) bool {
		if assert.Len(t, violations, 4) {
			assert.Equal(t, "top.intValue", violations[0].GetPropertyPath().Format())
			assert.Equal(t, "top.floatValue", violations[1].GetPropertyPath().Format())
			assert.Equal(t, "top.stringValue", violations[2].GetPropertyPath().Format())
			assert.Equal(t, "top.structValue.value", violations[3].GetPropertyPath().Format())
		}
		return true
	})
}

func TestFilter_NoViolations_Nil(t *testing.T) {
	err := validation.Filter(nil, nil)

	assert.NoError(t, err)
}

func TestFilter_SingleViolation_ViolationInList(t *testing.T) {
	violation := validation.NewViolation("code", "message", nil, nil)
	wrapped := fmt.Errorf("error: %w", violation)

	err := validation.Filter(nil, wrapped)

	validationtest.AssertIsViolationList(t, err, func(t *testing.T, violations validation.ViolationList) bool {
		return assert.Len(t, violations, 1) && assert.Equal(t, violation, violations[0])
	})
}

func TestFilter_ViolationList_ViolationsInList(t *testing.T) {
	violation := validation.NewViolation("code", "message", nil, nil)
	violations := validation.ViolationList{violation}
	wrapped := fmt.Errorf("error: %w", violations)

	err := validation.Filter(nil, wrapped)

	validationtest.AssertIsViolationList(t, err, func(t *testing.T, violations validation.ViolationList) bool {
		return assert.Len(t, violations, 1) && assert.Equal(t, violation, violations[0])
	})
}

func TestFilter_UnexpectedError_Error(t *testing.T) {
	unexpectedError := fmt.Errorf("error")

	err := validation.Filter(unexpectedError)

	assert.Equal(t, unexpectedError, err)
}

package test

import (
	"testing"

	"github.com/muonsoft/validation"
	"github.com/muonsoft/validation/code"
	"github.com/muonsoft/validation/it"
	"github.com/muonsoft/validation/validationtest"
	"github.com/stretchr/testify/assert"
)

func TestValidateEach_WhenSliceOfStrings_ExpectViolationOnEachElement(t *testing.T) {
	strings := []string{"", ""}

	err := newValidator(t).ValidateEach(strings, it.IsNotBlank())

	validationtest.AssertIsViolationList(t, err, func(t *testing.T, violations validation.ViolationList) bool {
		t.Helper()
		if assert.Len(t, violations, 2) {
			assert.Equal(t, code.NotBlank, violations[0].GetCode())
			assert.Equal(t, "[0]", violations[0].GetPropertyPath().Format())
			assert.Equal(t, code.NotBlank, violations[1].GetCode())
			assert.Equal(t, "[1]", violations[1].GetPropertyPath().Format())
		}
		return true
	})
}

func TestValidateEach_WhenMapOfStrings_ExpectViolationOnEachElement(t *testing.T) {
	strings := map[string]string{"key1": "", "key2": ""}

	err := newValidator(t).ValidateEach(strings, it.IsNotBlank())

	validationtest.AssertIsViolationList(t, err, func(t *testing.T, violations validation.ViolationList) bool {
		t.Helper()
		if assert.Len(t, violations, 2) {
			assert.Equal(t, code.NotBlank, violations[0].GetCode())
			assert.Contains(t, violations[0].GetPropertyPath().Format(), "key")
			assert.Equal(t, code.NotBlank, violations[1].GetCode())
			assert.Contains(t, violations[1].GetPropertyPath().Format(), "key")
		}
		return true
	})
}

func TestValidateEachString_WhenSliceOfStrings_ExpectViolationOnEachElement(t *testing.T) {
	strings := []string{"", ""}

	err := newValidator(t).ValidateEachString(strings, it.IsNotBlank())

	validationtest.AssertIsViolationList(t, err, func(t *testing.T, violations validation.ViolationList) bool {
		t.Helper()
		if assert.Len(t, violations, 2) {
			assert.Equal(t, code.NotBlank, violations[0].GetCode())
			assert.Equal(t, "[0]", violations[0].GetPropertyPath().Format())
			assert.Equal(t, code.NotBlank, violations[1].GetCode())
			assert.Equal(t, "[1]", violations[1].GetPropertyPath().Format())
		}
		return true
	})
}

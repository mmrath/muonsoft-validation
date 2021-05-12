package test

import (
	"github.com/muonsoft/validation"
	"github.com/muonsoft/validation/code"
	"github.com/muonsoft/validation/it"
	"github.com/muonsoft/validation/message"
)

var urlConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsURL passes on nil",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsURL()},
		assert:          assertNoError,
	},
	{
		name:            "IsURL passes on empty value",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsURL()},
		stringValue:     stringValue(""),
		assert:          assertNoError,
	},
	{
		name:            "IsURL passes on valid URL",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsURL()},
		stringValue:     stringValue("http://example.com"),
		assert:          assertNoError,
	},
	{
		name:            "IsURL violation on invalid URL",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsURL()},
		stringValue:     stringValue("example.com"),
		assert:          assertHasOneViolation(code.InvalidURL, message.InvalidURL, ""),
	},
	{
		name:            "IsURL error on empty schemas",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsURL().WithSchemas()},
		stringValue:     stringValue(""),
		assert:          assertError(`failed to set up constraint "URLConstraint": empty list of schemas`),
	},
	{
		name:            "IsURL passes on valid URL with custom schema",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsURL().WithSchemas("ftp")},
		stringValue:     stringValue("ftp://example.com"),
		assert:          assertNoError,
	},
	{
		name:            "IsURL with relative schema passes on valid relative URL",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsURL().WithRelativeSchema()},
		stringValue:     stringValue("//example.com"),
		assert:          assertNoError,
	},
	{
		name:            "IsURL with relative schema passes on valid absolute URL",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsURL().WithRelativeSchema()},
		stringValue:     stringValue("https://example.com"),
		assert:          assertNoError,
	},
	{
		name:            "IsURL violation on invalid URL with custom message",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsURL().Message(`Unexpected URL "{{ value }}"`)},
		stringValue:     stringValue("example.com"),
		assert:          assertHasOneViolation(code.InvalidURL, `Unexpected URL "example.com"`, ""),
	},
	{
		name:            "IsURL passes when condition is false",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsURL().When(false)},
		stringValue:     stringValue("example.com"),
		assert:          assertNoError,
	},
	{
		name:            "IsURL violation when condition is true",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsURL().When(true)},
		stringValue:     stringValue("example.com"),
		assert:          assertHasOneViolation(code.InvalidURL, message.InvalidURL, ""),
	},
}

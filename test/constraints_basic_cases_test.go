package test

import (
	"github.com/muonsoft/validation"
	"github.com/muonsoft/validation/code"
	"github.com/muonsoft/validation/it"
	"github.com/muonsoft/validation/message"

	"time"
)

var isNotBlankConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsNotBlank violation on nil",
		isApplicableFor: anyValueType,
		options:         []validation.Option{it.IsNotBlank()},
		assert:          assertHasOneViolation(code.NotBlank, message.NotBlank, ""),
	},
	{
		name:            "IsNotBlank violation on empty value",
		isApplicableFor: anyValueType,
		boolValue:       boolValue(false),
		intValue:        intValue(0),
		floatValue:      floatValue(0),
		stringValue:     stringValue(""),
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		options:         []validation.Option{it.IsNotBlank()},
		assert:          assertHasOneViolation(code.NotBlank, message.NotBlank, ""),
	},
	{
		name:            "IsNotBlank violation on empty value when condition is true",
		isApplicableFor: anyValueType,
		boolValue:       boolValue(false),
		intValue:        intValue(0),
		floatValue:      floatValue(0),
		stringValue:     stringValue(""),
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		options:         []validation.Option{it.IsNotBlank().When(true)},
		assert:          assertHasOneViolation(code.NotBlank, message.NotBlank, ""),
	},
	{
		name:            "IsNotBlank violation on nil with custom path",
		isApplicableFor: anyValueType,
		options: []validation.Option{
			validation.PropertyName("properties"),
			validation.ArrayIndex(0),
			validation.PropertyName("value"),
			it.IsNotBlank(),
		},
		assert: assertHasOneViolation(code.NotBlank, message.NotBlank, customPath),
	},
	{
		name:            "IsNotBlank violation on nil with custom message",
		isApplicableFor: anyValueType,
		options:         []validation.Option{it.IsNotBlank().Message(customMessage)},
		assert:          assertHasOneViolation(code.NotBlank, customMessage, ""),
	},
	{
		name:            "IsNotBlank passes on value",
		isApplicableFor: anyValueType,
		boolValue:       boolValue(true),
		intValue:        intValue(1),
		floatValue:      floatValue(0.1),
		stringValue:     stringValue("a"),
		timeValue:       timeValue(time.Now()),
		sliceValue:      []string{"a"},
		mapValue:        map[string]string{"a": "a"},
		options:         []validation.Option{it.IsNotBlank()},
		assert:          assertNoError,
	},
	{
		name:            "IsNotBlank passes on nil when allowed",
		isApplicableFor: exceptValueTypes("countable"),
		options:         []validation.Option{it.IsNotBlank().AllowNil()},
		assert:          assertNoError,
	},
	{
		name:            "IsNotBlank passes on nil when condition is false",
		isApplicableFor: exceptValueTypes("countable"),
		options:         []validation.Option{it.IsNotBlank().When(false)},
		assert:          assertNoError,
	},
}

var isBlankConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsBlank violation on value",
		isApplicableFor: anyValueType,
		boolValue:       boolValue(true),
		intValue:        intValue(1),
		floatValue:      floatValue(0.1),
		stringValue:     stringValue("a"),
		timeValue:       timeValue(time.Now()),
		sliceValue:      []string{"a"},
		mapValue:        map[string]string{"a": "a"},
		options:         []validation.Option{it.IsBlank()},
		assert:          assertHasOneViolation(code.Blank, message.Blank, ""),
	},
	{
		name:            "IsBlank violation on value when condition is true",
		isApplicableFor: anyValueType,
		boolValue:       boolValue(true),
		intValue:        intValue(1),
		floatValue:      floatValue(0.1),
		stringValue:     stringValue("a"),
		timeValue:       timeValue(time.Now()),
		sliceValue:      []string{"a"},
		mapValue:        map[string]string{"a": "a"},
		options:         []validation.Option{it.IsBlank().When(true)},
		assert:          assertHasOneViolation(code.Blank, message.Blank, ""),
	},
	{
		name:            "IsBlank violation on value with custom path",
		isApplicableFor: anyValueType,
		boolValue:       boolValue(true),
		intValue:        intValue(1),
		floatValue:      floatValue(0.1),
		stringValue:     stringValue("a"),
		timeValue:       timeValue(time.Now()),
		sliceValue:      []string{"a"},
		mapValue:        map[string]string{"a": "a"},
		options: []validation.Option{
			validation.PropertyName("properties"),
			validation.ArrayIndex(0),
			validation.PropertyName("value"),
			it.IsBlank(),
		},
		assert: assertHasOneViolation(code.Blank, message.Blank, customPath),
	},
	{
		name:            "IsBlank violation on value with custom message",
		isApplicableFor: anyValueType,
		boolValue:       boolValue(true),
		intValue:        intValue(1),
		floatValue:      floatValue(0.1),
		stringValue:     stringValue("a"),
		timeValue:       timeValue(time.Now()),
		sliceValue:      []string{"a"},
		mapValue:        map[string]string{"a": "a"},
		options:         []validation.Option{it.IsBlank().Message(customMessage)},
		assert:          assertHasOneViolation(code.Blank, customMessage, ""),
	},
	{
		name:            "IsBlank passes on nil",
		isApplicableFor: anyValueType,
		options:         []validation.Option{it.IsBlank()},
		assert:          assertNoError,
	},
	{
		name:            "IsBlank passes on empty value",
		isApplicableFor: anyValueType,
		boolValue:       boolValue(false),
		intValue:        intValue(0),
		floatValue:      floatValue(0.0),
		stringValue:     stringValue(""),
		timeValue:       timeValue(time.Time{}),
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		options:         []validation.Option{it.IsBlank()},
		assert:          assertNoError,
	},
	{
		name:            "IsBlank passes on value when condition is false",
		isApplicableFor: anyValueType,
		boolValue:       boolValue(true),
		intValue:        intValue(1),
		floatValue:      floatValue(0.1),
		stringValue:     stringValue("a"),
		timeValue:       timeValue(time.Now()),
		sliceValue:      []string{"a"},
		mapValue:        map[string]string{"a": "a"},
		options:         []validation.Option{it.IsBlank().When(false)},
		assert:          assertNoError,
	},
}

var isNotNilConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "isNotNil violation on nil",
		isApplicableFor: specificValueTypes(intType, floatType, stringType, timeType, iterableType),
		options:         []validation.Option{it.IsNotNil()},
		assert:          assertHasOneViolation(code.NotNil, message.NotNil, ""),
	},
	{
		name:            "isNotNil violation on empty value",
		isApplicableFor: specificValueTypes(intType, floatType, stringType, timeType, iterableType),
		intValue:        intValue(0),
		floatValue:      floatValue(0),
		stringValue:     stringValue(""),
		timeValue:       &time.Time{},
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		options:         []validation.Option{it.IsNotNil()},
		assert:          assertNoError,
	},
	{
		name:            "isNotNil violation on empty value when condition is true",
		isApplicableFor: specificValueTypes(intType, floatType, stringType, timeType, iterableType),
		intValue:        intValue(0),
		floatValue:      floatValue(0),
		stringValue:     stringValue(""),
		timeValue:       &time.Time{},
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		options:         []validation.Option{it.IsNotNil().When(true)},
		assert:          assertNoError,
	},
	{
		name:            "isNotNil violation on nil with custom path",
		isApplicableFor: specificValueTypes(intType, floatType, stringType, timeType, iterableType),
		options: []validation.Option{
			validation.PropertyName("properties"),
			validation.ArrayIndex(0),
			validation.PropertyName("value"),
			it.IsNotNil(),
		},
		assert: assertHasOneViolation(code.NotNil, message.NotNil, customPath),
	},
	{
		name:            "isNotNil violation on nil with custom message",
		isApplicableFor: specificValueTypes(intType, floatType, stringType, timeType, iterableType),
		options:         []validation.Option{it.IsNotNil().Message(customMessage)},
		assert:          assertHasOneViolation(code.NotNil, customMessage, ""),
	},
	{
		name:            "isNotNil passes on value",
		isApplicableFor: specificValueTypes(intType, floatType, stringType, timeType, iterableType),
		intValue:        intValue(1),
		floatValue:      floatValue(0.1),
		stringValue:     stringValue("a"),
		timeValue:       timeValue(time.Now()),
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		options:         []validation.Option{it.IsNotNil()},
		assert:          assertNoError,
	},
	{
		name:            "isNotNil passes on nil when condition is false",
		isApplicableFor: specificValueTypes(intType, floatType, stringType, timeType, iterableType),
		options:         []validation.Option{it.IsNotNil().When(false)},
		assert:          assertNoError,
	},
}

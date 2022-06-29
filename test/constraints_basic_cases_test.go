package test

import (
	"time"

	"github.com/muonsoft/validation"
	"github.com/muonsoft/validation/it"
	"github.com/muonsoft/validation/message"
)

var isNotBlankConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsNotBlank violation on nil",
		isApplicableFor: specificValueTypes(boolType, stringType, countableType, timeType),
		constraint:      it.IsNotBlank(),
		assert:          assertHasOneViolation(validation.ErrIsBlank, message.IsBlank),
	},
	{
		name:            "IsNotBlank violation on empty value",
		isApplicableFor: specificValueTypes(boolType, stringType, countableType, timeType),
		boolValue:       boolValue(false),
		intValue:        intValue(0),
		floatValue:      floatValue(0),
		stringValue:     stringValue(""),
		stringsValue:    []string{},
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		constraint:      it.IsNotBlank(),
		assert:          assertHasOneViolation(validation.ErrIsBlank, message.IsBlank),
	},
	{
		name:            "IsNotBlank violation on empty value when condition is true",
		isApplicableFor: specificValueTypes(boolType, stringType, countableType, timeType),
		boolValue:       boolValue(false),
		intValue:        intValue(0),
		floatValue:      floatValue(0),
		stringValue:     stringValue(""),
		stringsValue:    []string{},
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		constraint:      it.IsNotBlank().When(true),
		assert:          assertHasOneViolation(validation.ErrIsBlank, message.IsBlank),
	},
	{
		name:            "IsNotBlank violation on nil with custom message",
		isApplicableFor: specificValueTypes(boolType, stringType, countableType, timeType),
		constraint: it.IsNotBlank().
			WithError(ErrCustom).
			WithMessage(
				customMessage,
				validation.TemplateParameter{Key: "{{ custom }}", Value: "parameter"},
			),
		assert: assertHasOneViolation(ErrCustom, renderedCustomMessage),
	},
	{
		name:            "IsNotBlank passes on value",
		isApplicableFor: specificValueTypes(boolType, stringType, countableType, timeType),
		boolValue:       boolValue(true),
		intValue:        intValue(1),
		floatValue:      floatValue(0.1),
		stringValue:     stringValue("a"),
		stringsValue:    []string{""},
		timeValue:       timeValue(time.Now()),
		sliceValue:      []string{"a"},
		mapValue:        map[string]string{"a": "a"},
		constraint:      it.IsNotBlank(),
		assert:          assertNoError,
	},
	{
		name:            "IsNotBlank passes on nil when allowed",
		isApplicableFor: specificValueTypes(boolType, stringType, timeType),
		constraint:      it.IsNotBlank().AllowNil(),
		assert:          assertNoError,
	},
	{
		name:            "IsNotBlank passes on nil when condition is false",
		isApplicableFor: specificValueTypes(boolType, stringType, timeType),
		constraint:      it.IsNotBlank().When(false),
		assert:          assertNoError,
	},
	{
		name:            "IsNotBlank passes on nil when groups not match",
		isApplicableFor: specificValueTypes(boolType, stringType, timeType),
		constraint:      it.IsNotBlank().WhenGroups(testGroup),
		assert:          assertNoError,
	},
}

var isNotBlankNumberConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsNotBlankNumber violation on nil",
		isApplicableFor: specificValueTypes(intType),
		constraint:      it.IsNotBlankNumber[int](),
		assert:          assertHasOneViolation(validation.ErrIsBlank, message.IsBlank),
	},
	{
		name:            "IsNotBlankNumber violation on empty int value",
		isApplicableFor: specificValueTypes(intType),
		intValue:        intValue(0),
		constraint:      it.IsNotBlankNumber[int](),
		assert:          assertHasOneViolation(validation.ErrIsBlank, message.IsBlank),
	},
	{
		name:            "IsNotBlankNumber violation on empty float value",
		isApplicableFor: specificValueTypes(floatType),
		floatValue:      floatValue(0),
		constraint:      it.IsNotBlankNumber[float64](),
		assert:          assertHasOneViolation(validation.ErrIsBlank, message.IsBlank),
	},
	{
		name:            "IsNotBlankNumber violation on empty value when condition is true",
		isApplicableFor: specificValueTypes(intType),
		intValue:        intValue(0),
		constraint:      it.IsNotBlankNumber[int]().When(true),
		assert:          assertHasOneViolation(validation.ErrIsBlank, message.IsBlank),
	},
	{
		name:            "IsNotBlankNumber passes on value",
		isApplicableFor: specificValueTypes(intType),
		intValue:        intValue(1),
		constraint:      it.IsNotBlankNumber[int](),
		assert:          assertNoError,
	},
	{
		name:            "IsNotBlankNumber passes on nil when allowed",
		isApplicableFor: specificValueTypes(intType),
		constraint:      it.IsNotBlankNumber[int]().AllowNil(),
		assert:          assertNoError,
	},
	{
		name:            "IsNotBlankNumber passes on nil when condition is false",
		isApplicableFor: specificValueTypes(intType),
		constraint:      it.IsNotBlankNumber[int]().When(false),
		assert:          assertNoError,
	},
	{
		name:            "IsNotBlankNumber passes on nil when groups not match",
		isApplicableFor: specificValueTypes(intType),
		constraint:      it.IsNotBlankNumber[int]().WhenGroups(testGroup),
		assert:          assertNoError,
	},
}

var isNotBlankComparableConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsNotBlankComparable violation on nil",
		isApplicableFor: specificValueTypes(comparableType),
		constraint:      it.IsNotBlankComparable[string](),
		assert:          assertHasOneViolation(validation.ErrIsBlank, message.IsBlank),
	},
	{
		name:            "IsNotBlankComparable violation on empty value",
		isApplicableFor: specificValueTypes(comparableType),
		stringValue:     stringValue(""),
		constraint:      it.IsNotBlankComparable[string](),
		assert:          assertHasOneViolation(validation.ErrIsBlank, message.IsBlank),
	},
	{
		name:            "IsNotBlankComparable violation on empty value when condition is true",
		isApplicableFor: specificValueTypes(comparableType),
		stringValue:     stringValue(""),
		constraint:      it.IsNotBlankComparable[string]().When(true),
		assert:          assertHasOneViolation(validation.ErrIsBlank, message.IsBlank),
	},
	{
		name:            "IsNotBlankComparable passes on value",
		isApplicableFor: specificValueTypes(comparableType),
		stringValue:     stringValue("a"),
		constraint:      it.IsNotBlankComparable[string](),
		assert:          assertNoError,
	},
	{
		name:            "IsNotBlankComparable passes on nil when allowed",
		isApplicableFor: specificValueTypes(comparableType),
		constraint:      it.IsNotBlankComparable[string]().AllowNil(),
		assert:          assertNoError,
	},
	{
		name:            "IsNotBlankComparable passes on nil when condition is false",
		isApplicableFor: specificValueTypes(comparableType),
		constraint:      it.IsNotBlankComparable[string]().When(false),
		assert:          assertNoError,
	},
	{
		name:            "IsNotBlankComparable passes on nil when groups not match",
		isApplicableFor: specificValueTypes(comparableType),
		constraint:      it.IsNotBlankComparable[string]().WhenGroups(testGroup),
		assert:          assertNoError,
	},
}

var isBlankConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsBlank violation on value",
		isApplicableFor: specificValueTypes(boolType, stringType, countableType, timeType),
		boolValue:       boolValue(true),
		intValue:        intValue(1),
		floatValue:      floatValue(0.1),
		stringValue:     stringValue("a"),
		stringsValue:    []string{""},
		timeValue:       timeValue(time.Now()),
		sliceValue:      []string{"a"},
		mapValue:        map[string]string{"a": "a"},
		constraint:      it.IsBlank(),
		assert:          assertHasOneViolation(validation.ErrNotBlank, message.NotBlank),
	},
	{
		name:            "IsBlank violation on value when condition is true",
		isApplicableFor: specificValueTypes(boolType, stringType, countableType, timeType),
		boolValue:       boolValue(true),
		intValue:        intValue(1),
		floatValue:      floatValue(0.1),
		stringValue:     stringValue("a"),
		stringsValue:    []string{""},
		timeValue:       timeValue(time.Now()),
		sliceValue:      []string{"a"},
		mapValue:        map[string]string{"a": "a"},
		constraint:      it.IsBlank().When(true),
		assert:          assertHasOneViolation(validation.ErrNotBlank, message.NotBlank),
	},
	{
		name:            "IsBlank violation on value with custom message",
		isApplicableFor: specificValueTypes(boolType, stringType, countableType, timeType),
		boolValue:       boolValue(true),
		intValue:        intValue(1),
		floatValue:      floatValue(0.1),
		stringValue:     stringValue("a"),
		stringsValue:    []string{""},
		timeValue:       timeValue(time.Now()),
		sliceValue:      []string{"a"},
		mapValue:        map[string]string{"a": "a"},
		constraint: it.IsBlank().
			WithError(ErrCustom).
			WithMessage(
				customMessage,
				validation.TemplateParameter{Key: "{{ custom }}", Value: "parameter"},
			),
		assert: assertHasOneViolation(ErrCustom, renderedCustomMessage),
	},
	{
		name:            "IsBlank passes on nil",
		isApplicableFor: specificValueTypes(boolType, stringType, countableType, timeType),
		constraint:      it.IsBlank(),
		assert:          assertNoError,
	},
	{
		name:            "IsBlank passes on empty value",
		isApplicableFor: specificValueTypes(boolType, stringType, countableType, timeType),
		boolValue:       boolValue(false),
		intValue:        intValue(0),
		floatValue:      floatValue(0.0),
		stringValue:     stringValue(""),
		timeValue:       timeValue(time.Time{}),
		stringsValue:    []string{},
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		constraint:      it.IsBlank(),
		assert:          assertNoError,
	},
	{
		name:            "IsBlank passes on value when condition is false",
		isApplicableFor: specificValueTypes(boolType, stringType, countableType, timeType),
		boolValue:       boolValue(true),
		intValue:        intValue(1),
		floatValue:      floatValue(0.1),
		stringValue:     stringValue("a"),
		timeValue:       timeValue(time.Now()),
		stringsValue:    []string{""},
		sliceValue:      []string{"a"},
		mapValue:        map[string]string{"a": "a"},
		constraint:      it.IsBlank().When(false),
		assert:          assertNoError,
	},
	{
		name:            "IsBlank passes on value when groups not match",
		isApplicableFor: specificValueTypes(boolType, stringType, countableType, timeType),
		boolValue:       boolValue(true),
		intValue:        intValue(1),
		floatValue:      floatValue(0.1),
		stringValue:     stringValue("a"),
		timeValue:       timeValue(time.Now()),
		stringsValue:    []string{""},
		sliceValue:      []string{"a"},
		mapValue:        map[string]string{"a": "a"},
		constraint:      it.IsBlank().WhenGroups(testGroup),
		assert:          assertNoError,
	},
}

var isBlankNumberConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsBlankNumber violation on value",
		isApplicableFor: specificValueTypes(intType),
		intValue:        intValue(1),
		constraint:      it.IsBlankNumber[int](),
		assert:          assertHasOneViolation(validation.ErrNotBlank, message.NotBlank),
	},
	{
		name:            "IsBlankNumber violation on value when condition is true",
		isApplicableFor: specificValueTypes(intType),
		intValue:        intValue(1),
		constraint:      it.IsBlankNumber[int]().When(true),
		assert:          assertHasOneViolation(validation.ErrNotBlank, message.NotBlank),
	},
	{
		name:            "IsBlankNumber passes on nil",
		isApplicableFor: specificValueTypes(intType),
		constraint:      it.IsBlankNumber[int](),
		assert:          assertNoError,
	},
	{
		name:            "IsBlankNumber passes on empty value",
		isApplicableFor: specificValueTypes(intType),
		intValue:        intValue(0),
		constraint:      it.IsBlankNumber[int](),
		assert:          assertNoError,
	},
	{
		name:            "IsBlankNumber passes on value when condition is false",
		isApplicableFor: specificValueTypes(intType),
		intValue:        intValue(1),
		constraint:      it.IsBlankNumber[int]().When(false),
		assert:          assertNoError,
	},
	{
		name:            "IsBlankNumber passes on value when groups not match",
		isApplicableFor: specificValueTypes(intType),
		intValue:        intValue(1),
		constraint:      it.IsBlankNumber[int]().WhenGroups(testGroup),
		assert:          assertNoError,
	},
}

var isBlankComparableConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsBlankComparable violation on value",
		isApplicableFor: specificValueTypes(comparableType),
		stringValue:     stringValue("a"),
		constraint:      it.IsBlankComparable[string](),
		assert:          assertHasOneViolation(validation.ErrNotBlank, message.NotBlank),
	},
	{
		name:            "IsBlankComparable violation on value when condition is true",
		isApplicableFor: specificValueTypes(comparableType),
		stringValue:     stringValue("a"),
		constraint:      it.IsBlankComparable[string]().When(true),
		assert:          assertHasOneViolation(validation.ErrNotBlank, message.NotBlank),
	},
	{
		name:            "IsBlankComparable passes on nil",
		isApplicableFor: specificValueTypes(comparableType),
		constraint:      it.IsBlankComparable[string](),
		assert:          assertNoError,
	},
	{
		name:            "IsBlankComparable passes on empty value",
		isApplicableFor: specificValueTypes(comparableType),
		stringValue:     stringValue(""),
		constraint:      it.IsBlankComparable[string](),
		assert:          assertNoError,
	},
	{
		name:            "IsBlankComparable passes on value when condition is false",
		isApplicableFor: specificValueTypes(comparableType),
		stringValue:     stringValue("a"),
		constraint:      it.IsBlankComparable[string]().When(false),
		assert:          assertNoError,
	},
	{
		name:            "IsBlankComparable passes on value when groups not match",
		isApplicableFor: specificValueTypes(comparableType),
		stringValue:     stringValue("a"),
		constraint:      it.IsBlankComparable[string]().WhenGroups(testGroup),
		assert:          assertNoError,
	},
}

var isNotNilConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsNotNil violation on nil",
		isApplicableFor: specificValueTypes(nilType, boolType, stringType, timeType),
		constraint:      it.IsNotNil(),
		assert:          assertHasOneViolation(validation.ErrIsNil, message.IsNil),
	},
	{
		name:            "IsNotNil passes on empty value",
		isApplicableFor: specificValueTypes(nilType, boolType, stringType, timeType),
		boolValue:       boolValue(false),
		intValue:        intValue(0),
		floatValue:      floatValue(0),
		stringValue:     stringValue(""),
		stringsValue:    []string{},
		timeValue:       &time.Time{},
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		constraint:      it.IsNotNil(),
		assert:          assertNoError,
	},
	{
		name:            "IsNotNil passes on empty value when condition is true",
		isApplicableFor: specificValueTypes(nilType, boolType, stringType, timeType),
		boolValue:       boolValue(false),
		intValue:        intValue(0),
		floatValue:      floatValue(0),
		stringValue:     stringValue(""),
		stringsValue:    []string{},
		timeValue:       &time.Time{},
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		constraint:      it.IsNotNil().When(true),
		assert:          assertNoError,
	},
	{
		name:            "IsNotNil violation on nil with custom message",
		isApplicableFor: specificValueTypes(nilType, boolType, stringType, timeType),
		constraint: it.IsNotNil().
			WithError(ErrCustom).
			WithMessage(
				customMessage,
				validation.TemplateParameter{Key: "{{ custom }}", Value: "parameter"},
			),
		assert: assertHasOneViolation(ErrCustom, renderedCustomMessage),
	},
	{
		name:            "IsNotNil passes on value",
		isApplicableFor: specificValueTypes(nilType, boolType, stringType, timeType),
		boolValue:       boolValue(true),
		intValue:        intValue(1),
		floatValue:      floatValue(0.1),
		stringValue:     stringValue("a"),
		stringsValue:    []string{},
		timeValue:       timeValue(time.Now()),
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		constraint:      it.IsNotNil(),
		assert:          assertNoError,
	},
	{
		name:            "IsNotNil passes on nil when condition is false",
		isApplicableFor: specificValueTypes(nilType, boolType, stringType, timeType),
		constraint:      it.IsNotNil().When(false),
		assert:          assertNoError,
	},
	{
		name:            "IsNotNil passes on nil when groups not match",
		isApplicableFor: specificValueTypes(nilType, boolType, stringType, timeType),
		constraint:      it.IsNotNil().WhenGroups(testGroup),
		assert:          assertNoError,
	},
}

var isNotNilNumberConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsNotNilNumber violation on nil",
		isApplicableFor: specificValueTypes(intType),
		constraint:      it.IsNotNilNumber[int](),
		assert:          assertHasOneViolation(validation.ErrIsNil, message.IsNil),
	},
	{
		name:            "IsNotNilNumber passes on empty value",
		isApplicableFor: specificValueTypes(intType),
		intValue:        intValue(0),
		constraint:      it.IsNotNilNumber[int](),
		assert:          assertNoError,
	},
	{
		name:            "IsNotNilNumber passes on empty value when condition is true",
		isApplicableFor: specificValueTypes(intType),
		intValue:        intValue(0),
		constraint:      it.IsNotNilNumber[int]().When(true),
		assert:          assertNoError,
	},
	{
		name:            "IsNotNilNumber passes on value",
		isApplicableFor: specificValueTypes(intType),
		intValue:        intValue(1),
		constraint:      it.IsNotNilNumber[int](),
		assert:          assertNoError,
	},
	{
		name:            "IsNotNilNumber passes on nil when condition is false",
		isApplicableFor: specificValueTypes(intType),
		constraint:      it.IsNotNilNumber[int]().When(false),
		assert:          assertNoError,
	},
	{
		name:            "IsNotNilNumber passes on nil when groups not match",
		isApplicableFor: specificValueTypes(intType),
		constraint:      it.IsNotNilNumber[int]().WhenGroups(testGroup),
		assert:          assertNoError,
	},
}

var isNotNilComparableConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsNotNilComparable violation on nil",
		isApplicableFor: specificValueTypes(comparableType),
		constraint:      it.IsNotNilComparable[string](),
		assert:          assertHasOneViolation(validation.ErrIsNil, message.IsNil),
	},
	{
		name:            "IsNotNilComparable passes on empty value",
		isApplicableFor: specificValueTypes(comparableType),
		stringValue:     stringValue(""),
		constraint:      it.IsNotNilComparable[string](),
		assert:          assertNoError,
	},
	{
		name:            "IsNotNilComparable passes on empty value when condition is true",
		isApplicableFor: specificValueTypes(comparableType),
		stringValue:     stringValue(""),
		constraint:      it.IsNotNilComparable[string]().When(true),
		assert:          assertNoError,
	},
	{
		name:            "IsNotNilComparable passes on value",
		isApplicableFor: specificValueTypes(comparableType),
		stringValue:     stringValue("a"),
		constraint:      it.IsNotNilComparable[string](),
		assert:          assertNoError,
	},
	{
		name:            "IsNotNilComparable passes on nil when condition is false",
		isApplicableFor: specificValueTypes(comparableType),
		constraint:      it.IsNotNilComparable[string]().When(false),
		assert:          assertNoError,
	},
	{
		name:            "IsNotNilComparable passes on nil when groups not match",
		isApplicableFor: specificValueTypes(comparableType),
		constraint:      it.IsNotNilComparable[string]().WhenGroups(testGroup),
		assert:          assertNoError,
	},
}

var isNilConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsNil passes on nil",
		isApplicableFor: specificValueTypes(nilType, boolType, stringType, timeType),
		constraint:      it.IsNil(),
		assert:          assertNoError,
	},
	{
		name:            "IsNil violation on empty value",
		isApplicableFor: specificValueTypes(nilType, boolType, stringType, timeType),
		boolValue:       boolValue(false),
		intValue:        intValue(0),
		floatValue:      floatValue(0),
		stringValue:     stringValue(""),
		stringsValue:    []string{},
		timeValue:       &time.Time{},
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		constraint:      it.IsNil(),
		assert:          assertHasOneViolation(validation.ErrNotNil, message.NotNil),
	},
	{
		name:            "IsNil passes on nil when condition is true",
		isApplicableFor: specificValueTypes(nilType, boolType, stringType, timeType),
		constraint:      it.IsNil().When(true),
		assert:          assertNoError,
	},
	{
		name:            "IsNil violation on empty value with custom message",
		isApplicableFor: specificValueTypes(nilType, boolType, stringType, timeType),
		boolValue:       boolValue(false),
		intValue:        intValue(0),
		floatValue:      floatValue(0),
		stringValue:     stringValue(""),
		stringsValue:    []string{},
		timeValue:       &time.Time{},
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		constraint: it.IsNil().
			WithError(ErrCustom).
			WithMessage(
				customMessage,
				validation.TemplateParameter{Key: "{{ custom }}", Value: "parameter"},
			),
		assert: assertHasOneViolation(ErrCustom, renderedCustomMessage),
	},
	{
		name:            "IsNil violation on value",
		isApplicableFor: specificValueTypes(nilType, boolType, stringType, timeType),
		boolValue:       boolValue(true),
		intValue:        intValue(1),
		floatValue:      floatValue(0.1),
		stringValue:     stringValue("a"),
		stringsValue:    []string{},
		timeValue:       timeValue(time.Now()),
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		constraint:      it.IsNil(),
		assert:          assertHasOneViolation(validation.ErrNotNil, message.NotNil),
	},
	{
		name:            "IsNil passes on empty value when condition is false",
		isApplicableFor: specificValueTypes(nilType, boolType, stringType, timeType),
		boolValue:       boolValue(false),
		intValue:        intValue(0),
		floatValue:      floatValue(0),
		stringValue:     stringValue(""),
		stringsValue:    []string{},
		timeValue:       &time.Time{},
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		constraint:      it.IsNil().When(false),
		assert:          assertNoError,
	},
	{
		name:            "IsNil passes on empty value when groups not match",
		isApplicableFor: specificValueTypes(nilType, boolType, stringType, timeType),
		boolValue:       boolValue(false),
		intValue:        intValue(0),
		floatValue:      floatValue(0),
		stringValue:     stringValue(""),
		stringsValue:    []string{},
		timeValue:       &time.Time{},
		sliceValue:      []string{},
		mapValue:        map[string]string{},
		constraint:      it.IsNil().WhenGroups(testGroup),
		assert:          assertNoError,
	},
}

var isNilNumberConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsNilNumber passes on nil",
		isApplicableFor: specificValueTypes(intType),
		constraint:      it.IsNilNumber[int](),
		assert:          assertNoError,
	},
	{
		name:            "IsNilNumber violation on empty value",
		isApplicableFor: specificValueTypes(intType),
		intValue:        intValue(0),
		constraint:      it.IsNilNumber[int](),
		assert:          assertHasOneViolation(validation.ErrNotNil, message.NotNil),
	},
	{
		name:            "IsNilNumber passes on nil when condition is true",
		isApplicableFor: specificValueTypes(intType),
		constraint:      it.IsNilNumber[int]().When(true),
		assert:          assertNoError,
	},
	{
		name:            "IsNilNumber violation on value",
		isApplicableFor: specificValueTypes(intType),
		intValue:        intValue(1),
		constraint:      it.IsNilNumber[int](),
		assert:          assertHasOneViolation(validation.ErrNotNil, message.NotNil),
	},
	{
		name:            "IsNilNumber passes on empty value when condition is false",
		isApplicableFor: specificValueTypes(intType),
		intValue:        intValue(0),
		constraint:      it.IsNilNumber[int]().When(false),
		assert:          assertNoError,
	},
	{
		name:            "IsNilNumber passes on empty value when groups not match",
		isApplicableFor: specificValueTypes(intType),
		intValue:        intValue(0),
		constraint:      it.IsNilNumber[int]().WhenGroups(testGroup),
		assert:          assertNoError,
	},
}

var isNilComparableConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsNilComparable passes on nil",
		isApplicableFor: specificValueTypes(comparableType),
		constraint:      it.IsNilComparable[string](),
		assert:          assertNoError,
	},
	{
		name:            "IsNilComparable violation on empty value",
		isApplicableFor: specificValueTypes(comparableType),
		stringValue:     stringValue(""),
		constraint:      it.IsNilComparable[string](),
		assert:          assertHasOneViolation(validation.ErrNotNil, message.NotNil),
	},
	{
		name:            "IsNilComparable passes on nil when condition is true",
		isApplicableFor: specificValueTypes(comparableType),
		constraint:      it.IsNilComparable[string]().When(true),
		assert:          assertNoError,
	},
	{
		name:            "IsNilComparable violation on value",
		isApplicableFor: specificValueTypes(comparableType),
		stringValue:     stringValue("a"),
		constraint:      it.IsNilComparable[string](),
		assert:          assertHasOneViolation(validation.ErrNotNil, message.NotNil),
	},
	{
		name:            "IsNilComparable passes on empty value when condition is false",
		isApplicableFor: specificValueTypes(comparableType),
		stringValue:     stringValue(""),
		constraint:      it.IsNilComparable[string]().When(false),
		assert:          assertNoError,
	},
	{
		name:            "IsNilComparable passes on empty value when groups not match",
		isApplicableFor: specificValueTypes(comparableType),
		stringValue:     stringValue(""),
		constraint:      it.IsNilComparable[string]().WhenGroups(testGroup),
		assert:          assertNoError,
	},
}

var isTrueConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsTrue passes on nil",
		isApplicableFor: specificValueTypes(boolType),
		constraint:      it.IsTrue(),
		assert:          assertNoError,
	},
	{
		name:            "IsTrue violation on empty value",
		isApplicableFor: specificValueTypes(boolType),
		boolValue:       boolValue(false),
		constraint:      it.IsTrue(),
		assert:          assertHasOneViolation(validation.ErrNotTrue, message.NotTrue),
	},
	{
		name:            "IsTrue violation on empty value when condition is true",
		isApplicableFor: specificValueTypes(boolType),
		boolValue:       boolValue(false),
		constraint:      it.IsTrue().When(true),
		assert:          assertHasOneViolation(validation.ErrNotTrue, message.NotTrue),
	},
	{
		name:            "IsTrue violation on empty value with custom message",
		isApplicableFor: specificValueTypes(boolType),
		boolValue:       boolValue(false),
		constraint: it.IsTrue().
			WithError(ErrCustom).
			WithMessage(
				customMessage,
				validation.TemplateParameter{Key: "{{ custom }}", Value: "parameter"},
			),
		assert: assertHasOneViolation(ErrCustom, renderedCustomMessage),
	},
	{
		name:            "IsTrue passes on value",
		isApplicableFor: specificValueTypes(boolType),
		boolValue:       boolValue(true),
		constraint:      it.IsTrue(),
		assert:          assertNoError,
	},
	{
		name:            "IsTrue passes on empty value when condition is false",
		isApplicableFor: specificValueTypes(boolType),
		boolValue:       boolValue(false),
		constraint:      it.IsTrue().When(false),
		assert:          assertNoError,
	},
	{
		name:            "IsTrue passes on empty value when groups not match",
		isApplicableFor: specificValueTypes(boolType),
		boolValue:       boolValue(false),
		constraint:      it.IsTrue().WhenGroups(testGroup),
		assert:          assertNoError,
	},
}

var isFalseConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsFalse passes on nil",
		isApplicableFor: specificValueTypes(boolType),
		constraint:      it.IsFalse(),
		assert:          assertNoError,
	},
	{
		name:            "IsFalse passes on empty value",
		isApplicableFor: specificValueTypes(boolType),
		boolValue:       boolValue(false),
		constraint:      it.IsFalse(),
		assert:          assertNoError,
	},
	{
		name:            "IsFalse violation on error value when condition is true",
		isApplicableFor: specificValueTypes(boolType),
		boolValue:       boolValue(true),
		constraint:      it.IsFalse().When(true),
		assert:          assertHasOneViolation(validation.ErrNotFalse, message.NotFalse),
	},
	{
		name:            "IsFalse violation on error value with custom message",
		isApplicableFor: specificValueTypes(boolType),
		boolValue:       boolValue(true),
		constraint:      it.IsFalse().WithMessage(customMessage),
		assert:          assertHasOneViolation(validation.ErrNotFalse, customMessage),
	},
	{
		name:            "IsFalse passes on value",
		isApplicableFor: specificValueTypes(boolType),
		boolValue:       boolValue(false),
		constraint:      it.IsFalse(),
		assert:          assertNoError,
	},
	{
		name:            "IsFalse passes on error value when condition is false",
		isApplicableFor: specificValueTypes(boolType),
		boolValue:       boolValue(true),
		constraint:      it.IsFalse().When(false),
		assert:          assertNoError,
	},
	{
		name:            "IsFalse passes on error value when groups not match",
		isApplicableFor: specificValueTypes(boolType),
		boolValue:       boolValue(true),
		constraint:      it.IsFalse().WhenGroups(testGroup),
		assert:          assertNoError,
	},
}

package test

import (
	"net"

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

var emailConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsEmail passes on valid email",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsEmail()},
		stringValue:     stringValue("user@example.com"),
		assert:          assertNoError,
	},
	{
		name:            "IsEmail violation on invalid email",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsEmail()},
		stringValue:     stringValue("invalid"),
		assert:          assertHasOneViolation(code.InvalidEmail, message.InvalidEmail, ""),
	},
	{
		name:            "IsHTML5Email passes on valid email",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsHTML5Email()},
		stringValue:     stringValue("user@example.com"),
		assert:          assertNoError,
	},
	{
		name:            "IsHTML5Email violation on invalid email",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsHTML5Email()},
		stringValue:     stringValue("invalid"),
		assert:          assertHasOneViolation(code.InvalidEmail, message.InvalidEmail, ""),
	},
}

var ipConstraintTestCases = []ConstraintValidationTestCase{
	{
		name:            "IsIP passes on nil",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsIP()},
		assert:          assertNoError,
	},
	{
		name:            "IsIP passes on empty value",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsIP()},
		stringValue:     stringValue(""),
		assert:          assertNoError,
	},
	{
		name:            "IsIP passes on valid IP v4",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsIP()},
		stringValue:     stringValue("123.123.123.123"),
		assert:          assertNoError,
	},
	{
		name:            "IsIP violation on invalid IP v4",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsIP()},
		stringValue:     stringValue("123.123.123.321"),
		assert:          assertHasOneViolation(code.InvalidIP, message.InvalidIP, ""),
	},
	{
		name:            "IsIPv4 passes on valid IP v4",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsIPv4()},
		stringValue:     stringValue("123.123.123.123"),
		assert:          assertNoError,
	},
	{
		name:            "IsIPv4 violation on IP v6",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsIPv4()},
		stringValue:     stringValue("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
		assert:          assertHasOneViolation(code.InvalidIP, message.InvalidIP, ""),
	},
	{
		name:            "IsIPv6 passes on valid IP v6",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsIPv6()},
		stringValue:     stringValue("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
		assert:          assertNoError,
	},
	{
		name:            "IsIPv6 violation on IP v4",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsIPv6()},
		stringValue:     stringValue("123.123.123.123"),
		assert:          assertHasOneViolation(code.InvalidIP, message.InvalidIP, ""),
	},
	{
		name:            "IsIP violation on private IP",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsIP().DenyPrivateIP()},
		stringValue:     stringValue("192.168.1.0"),
		assert:          assertHasOneViolation(code.ProhibitedIP, message.ProhibitedIP, ""),
	},
	{
		name:            "IsIP violation on custom IP",
		isApplicableFor: specificValueTypes(stringType),
		options: []validation.Option{it.IsIP().DenyIP(func(ip net.IP) bool {
			return ip.IsLoopback()
		})},
		stringValue: stringValue("127.0.0.1"),
		assert:      assertHasOneViolation(code.ProhibitedIP, message.ProhibitedIP, ""),
	},
	{
		name:            "IsIP violation with custom message",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsIP().InvalidMessage(`Unexpected IP "{{ value }}"`)},
		stringValue:     stringValue("123.123.123.321"),
		assert:          assertHasOneViolation(code.InvalidIP, `Unexpected IP "123.123.123.321"`, ""),
	},
	{
		name:            "IsIP violation with custom restricted message",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsIP().DenyPrivateIP().ProhibitedMessage(`Unexpected IP "{{ value }}"`)},
		stringValue:     stringValue("192.168.1.0"),
		assert:          assertHasOneViolation(code.ProhibitedIP, `Unexpected IP "192.168.1.0"`, ""),
	},
	{
		name:            "IsIP passes when condition is false",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsIP().When(false)},
		stringValue:     stringValue("123.123.123.321"),
		assert:          assertNoError,
	},
	{
		name:            "IsIP passes when condition is false",
		isApplicableFor: specificValueTypes(stringType),
		options:         []validation.Option{it.IsIP().When(true)},
		stringValue:     stringValue("123.123.123.321"),
		assert:          assertHasOneViolation(code.InvalidIP, message.InvalidIP, ""),
	},
}
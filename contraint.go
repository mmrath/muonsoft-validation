package validation

import (
	"github.com/muonsoft/validation/code"
	"github.com/muonsoft/validation/generic"
	"github.com/muonsoft/validation/message"

	"time"
)

// Constraint is the base interface to build validation constraints.
type Constraint interface {
	Option
	// Name is a constraint name that can be used in internal errors.
	Name() string
}

// NilConstraint is used for constraints that need to check value for nil. In common case
// you do not need to implement it in your constraints because nil values should be ignored.
type NilConstraint interface {
	Constraint
	ValidateNil(scope Scope) error
}

// BoolConstraint is used to build constraints for boolean values validation.
type BoolConstraint interface {
	Constraint
	ValidateBool(value *bool, scope Scope) error
}

// NumberConstraint is used to build constraints for numeric values validation.
//
// At this moment working with numbers is based on reflection.
// Be aware. This constraint is subject to be changed after generics implementation in Go.
type NumberConstraint interface {
	Constraint
	ValidateNumber(value generic.Number, scope Scope) error
}

// StringConstraint is used to build constraints for string values validation.
type StringConstraint interface {
	Constraint
	ValidateString(value *string, scope Scope) error
}

// IterableConstraint is used to build constraints for validation of iterables (arrays, slices, or maps).
//
// At this moment working with numbers is based on reflection.
// Be aware. This constraint is subject to be changed after generics implementation in Go.
type IterableConstraint interface {
	Constraint
	ValidateIterable(value generic.Iterable, scope Scope) error
}

// CountableConstraint is used to build constraints for simpler validation of iterable elements count.
type CountableConstraint interface {
	Constraint
	ValidateCountable(count int, scope Scope) error
}

// TimeConstraint is used to build constraints for date/time validation.
type TimeConstraint interface {
	Constraint
	ValidateTime(value *time.Time, scope Scope) error
}

type ValidateStringFunc func(value string) bool

type CustomStringConstraint struct {
	isIgnored       bool
	isValid         ValidateStringFunc
	name            string
	code            string
	messageTemplate string
}

func NewCustomStringConstraint(isValid ValidateStringFunc, parameters ...string) CustomStringConstraint {
	constraint := CustomStringConstraint{
		isValid:         isValid,
		name:            "CustomStringConstraint",
		code:            code.NotValid,
		messageTemplate: message.NotValid,
	}

	if len(parameters) > 0 {
		constraint.name = parameters[0]
	}
	if len(parameters) > 1 {
		constraint.code = parameters[1]
	}
	if len(parameters) > 2 {
		constraint.messageTemplate = parameters[2]
	}

	return constraint
}

func (c CustomStringConstraint) SetUp() error {
	return nil
}

func (c CustomStringConstraint) Name() string {
	return c.name
}

func (c CustomStringConstraint) Message(message string) CustomStringConstraint {
	c.messageTemplate = message
	return c
}

func (c CustomStringConstraint) When(condition bool) CustomStringConstraint {
	c.isIgnored = !condition
	return c
}

func (c CustomStringConstraint) ValidateString(value *string, scope Scope) error {
	if c.isIgnored || value == nil || *value == "" || c.isValid(*value) {
		return nil
	}

	return scope.BuildViolation(c.code, c.messageTemplate).
		AddParameter("{{ value }}", *value).
		CreateViolation()
}

type notFoundConstraint struct {
	key string
}

func (c notFoundConstraint) SetUp() error {
	return ConstraintNotFoundError{Key: c.key}
}

func (c notFoundConstraint) Name() string {
	return "notFoundConstraint"
}

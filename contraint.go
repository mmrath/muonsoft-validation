package validation

import (
	"time"

	"github.com/muonsoft/validation/code"
	"github.com/muonsoft/validation/message"
)

type Numeric interface {
	~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Constraint is the base interface to build validation constraints.
// Deprecated
type Constraint interface {
	// Name is a constraint name that can be used in internal errors.
	Name() string
}

// BoolConstraint is used to build constraints for boolean values validation.
type BoolConstraint interface {
	ValidateBool(value *bool, scope Scope) error
}

// NumberConstraint is used to build constraints for numeric values validation.
type NumberConstraint[T Numeric] interface {
	ValidateNumber(value *T, scope Scope) error
}

// StringConstraint is used to build constraints for string values validation.
type StringConstraint interface {
	ValidateString(value *string, scope Scope) error
}

// ComparableConstraint is used to build constraints for generic comparable value validation.
type ComparableConstraint[T comparable] interface {
	ValidateComparable(value *T, scope Scope) error
}

// ComparablesConstraint is used to build constraints for generic comparable values validation.
type ComparablesConstraint[T comparable] interface {
	ValidateComparables(values []T, scope Scope) error
}

// todo ? or NillableConstraint?
type SliceConstraint[T any] interface {
	ValidateSlice(value []T, scope Scope) error
}

// todo ?
type MapConstraint[K comparable, V any] interface {
	ValidateMap(value map[K]V, scope Scope) error
}

// CountableConstraint is used to build constraints for simpler validation of iterable elements count.
type CountableConstraint interface {
	ValidateCountable(count int, scope Scope) error
}

// TimeConstraint is used to build constraints for date/time validation.
type TimeConstraint interface {
	ValidateTime(value *time.Time, scope Scope) error
}

type controlConstraint interface {
	validate(scope Scope, validate ValidateByConstraintFunc) (*ViolationList, error)
}

// CustomStringConstraint can be used to create custom constraints for validating string values
// based on function with signature func(string) bool.
type CustomStringConstraint struct {
	isIgnored         bool
	isValid           func(string) bool
	groups            []string
	code              string
	messageTemplate   string
	messageParameters TemplateParameterList
}

// NewCustomStringConstraint creates a new string constraint from a function with signature func(string) bool.
// Optional parameters can be used to set up violation code (first), message template (second).
// All other parameters are ignored.
func NewCustomStringConstraint(isValid func(string) bool, parameters ...string) CustomStringConstraint {
	constraint := CustomStringConstraint{
		isValid:         isValid,
		code:            code.NotValid,
		messageTemplate: message.Templates[code.NotValid],
	}

	if len(parameters) > 0 {
		constraint.code = parameters[0]
	}
	if len(parameters) > 1 {
		constraint.messageTemplate = parameters[1]
	}

	return constraint
}

// Code overrides default code for produced violation.
func (c CustomStringConstraint) Code(code string) CustomStringConstraint {
	c.code = code
	return c
}

// Message sets the violation message template. You can set custom template parameters
// for injecting its values into the final message. Also, you can use default parameters:
//
//	{{ value }} - the current (invalid) value.
func (c CustomStringConstraint) Message(template string, parameters ...TemplateParameter) CustomStringConstraint {
	c.messageTemplate = template
	c.messageParameters = parameters
	return c
}

// When enables conditional validation of this constraint. If the expression evaluates to false,
// then the constraint will be ignored.
func (c CustomStringConstraint) When(condition bool) CustomStringConstraint {
	c.isIgnored = !condition
	return c
}

// WhenGroups enables conditional validation of the constraint by using the validation groups.
func (c CustomStringConstraint) WhenGroups(groups ...string) CustomStringConstraint {
	c.groups = groups
	return c
}

func (c CustomStringConstraint) ValidateString(value *string, scope Scope) error {
	if c.isIgnored || scope.IsIgnored(c.groups...) || value == nil || *value == "" || c.isValid(*value) {
		return nil
	}

	return scope.BuildViolation(c.code, c.messageTemplate).
		SetParameters(
			c.messageParameters.Prepend(
				TemplateParameter{Key: "{{ value }}", Value: *value},
			)...,
		).
		AddParameter("{{ value }}", *value).
		CreateViolation()
}

//
// // ConditionalConstraint is used for conditional validation.
// // Use the When function to initiate a conditional check.
// // If the condition is true, then the constraints passed through the Then function will be applied.
// // Otherwise, the constraints passed through the Else function will apply.
// type ConditionalConstraint struct {
// 	condition       bool
// 	thenConstraints []Constraint
// 	elseConstraints []Constraint
// }
//
// // When function is used to initiate conditional validation.
// // If the condition is true, then the constraints passed through the Then function will be applied.
// // Otherwise, the constraints passed through the Else function will apply.
// func When(condition bool) ConditionalConstraint {
// 	return ConditionalConstraint{condition: condition}
// }
//
// // Then function is used to set a sequence of constraints to be applied if the condition is true.
// // If the list is empty error will be returned.
// func (c ConditionalConstraint) Then(constraints ...Constraint) ConditionalConstraint {
// 	c.thenConstraints = constraints
// 	return c
// }
//
// // Else function is used to set a sequence of constraints to be applied if a condition is false.
// func (c ConditionalConstraint) Else(constraints ...Constraint) ConditionalConstraint {
// 	c.elseConstraints = constraints
// 	return c
// }
//
// // SetUp will return an error if Then function did not set any constraints.
// func (c ConditionalConstraint) SetUp() error {
// 	if len(c.thenConstraints) == 0 {
// 		return errThenBranchNotSet
// 	}
//
// 	return nil
// }
//
// // Name is the constraint name.
// func (c ConditionalConstraint) Name() string {
// 	return "ConditionalConstraint"
// }
//
// func (c ConditionalConstraint) validate(
// 	scope Scope,
// 	validate ValidateByConstraintFunc,
// ) (*ViolationList, error) {
// 	violations := &ViolationList{}
// 	var constraints []Constraint
//
// 	if c.condition {
// 		constraints = c.thenConstraints
// 	} else {
// 		constraints = c.elseConstraints
// 	}
//
// 	for _, constraint := range constraints {
// 		err := violations.AppendFromError(validate(constraint, scope))
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
//
// 	return violations, nil
// }
//
// // SequentialConstraint is used to set constraints allowing to interrupt the validation once
// // the first violation is raised.
// type SequentialConstraint struct {
// 	isIgnored   bool
// 	groups      []string
// 	constraints []Constraint
// }
//
// // Sequentially function used to set of constraints that should be validated step-by-step.
// // If the list is empty error will be returned.
// func Sequentially(constraints ...Constraint) SequentialConstraint {
// 	return SequentialConstraint{constraints: constraints}
// }
//
// // SetUp will return an error if the list of constraints is empty.
// func (c SequentialConstraint) SetUp() error {
// 	if len(c.constraints) == 0 {
// 		return errSequentiallyConstraintsNotSet
// 	}
// 	return nil
// }
//
// // Name is the constraint name.
// func (c SequentialConstraint) Name() string {
// 	return "SequentialConstraint"
// }
//
// // When enables conditional validation of this constraint. If the expression evaluates to false,
// // then the constraint will be ignored.
// func (c SequentialConstraint) When(condition bool) SequentialConstraint {
// 	c.isIgnored = !condition
// 	return c
// }
//
// // WhenGroups enables conditional validation of the constraint by using the validation groups.
// func (c SequentialConstraint) WhenGroups(groups ...string) SequentialConstraint {
// 	c.groups = groups
// 	return c
// }
//
// func (c SequentialConstraint) validate(
// 	scope Scope,
// 	validate ValidateByConstraintFunc,
// ) (*ViolationList, error) {
// 	if c.isIgnored || scope.IsIgnored(c.groups...) {
// 		return nil, nil
// 	}
//
// 	violations := &ViolationList{}
//
// 	for _, constraint := range c.constraints {
// 		err := violations.AppendFromError(validate(constraint, scope))
// 		if err != nil {
// 			return nil, err
// 		} else if violations.len > 0 {
// 			return violations, nil
// 		}
// 	}
//
// 	return violations, nil
// }
//
// // AtLeastOneOfConstraint is used to set constraints allowing checks that the value satisfies
// // at least one of the given constraints.
// // The validation stops as soon as one constraint is satisfied.
// type AtLeastOneOfConstraint struct {
// 	isIgnored   bool
// 	groups      []string
// 	constraints []Constraint
// }
//
// // AtLeastOneOf function used to set of constraints that the value satisfies at least one of the given constraints.
// // If the list is empty error will be returned.
// func AtLeastOneOf(constraints ...Constraint) AtLeastOneOfConstraint {
// 	return AtLeastOneOfConstraint{constraints: constraints}
// }
//
// // SetUp will return an error if the list of constraints is empty.
// func (c AtLeastOneOfConstraint) SetUp() error {
// 	if len(c.constraints) == 0 {
// 		return errAtLeastOneOfConstraintsNotSet
// 	}
// 	return nil
// }
//
// // Name is the constraint name.
// func (c AtLeastOneOfConstraint) Name() string {
// 	return "AtLeastOneOfConstraint"
// }
//
// // When enables conditional validation of this constraint. If the expression evaluates to false,
// // then the constraint will be ignored.
// func (c AtLeastOneOfConstraint) When(condition bool) AtLeastOneOfConstraint {
// 	c.isIgnored = !condition
// 	return c
// }
//
// // WhenGroups enables conditional validation of the constraint by using the validation groups.
// func (c AtLeastOneOfConstraint) WhenGroups(groups ...string) AtLeastOneOfConstraint {
// 	c.groups = groups
// 	return c
// }
//
// func (c AtLeastOneOfConstraint) validate(
// 	scope Scope,
// 	validate ValidateByConstraintFunc,
// ) (*ViolationList, error) {
// 	if c.isIgnored || scope.IsIgnored(c.groups...) {
// 		return nil, nil
// 	}
//
// 	violations := &ViolationList{}
//
// 	for _, constraint := range c.constraints {
// 		violation := validate(constraint, scope)
// 		if violation == nil {
// 			return nil, nil
// 		}
//
// 		err := violations.AppendFromError(violation)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
//
// 	return violations, nil
// }
//
// // CompoundConstraint is used to create your own set of reusable constraints, representing rules to use consistently.
// type CompoundConstraint struct {
// 	isIgnored   bool
// 	groups      []string
// 	constraints []Constraint
// }
//
// // Compound function used to create set of reusable constraints.
// // If the list is empty error will be returned.
// func Compound(constraints ...Constraint) CompoundConstraint {
// 	return CompoundConstraint{constraints: constraints}
// }
//
// // SetUp will return an error if the list of constraints is empty.
// func (c CompoundConstraint) SetUp() error {
// 	if len(c.constraints) == 0 {
// 		return errCompoundConstraintsNotSet
// 	}
// 	return nil
// }
//
// // Name is the constraint name.
// func (c CompoundConstraint) Name() string {
// 	return "CompoundConstraint"
// }
//
// // When enables conditional validation of this constraint. If the expression evaluates to false,
// // then the constraint will be ignored.
// func (c CompoundConstraint) When(condition bool) CompoundConstraint {
// 	c.isIgnored = !condition
// 	return c
// }
//
// // WhenGroups enables conditional validation of the constraint by using the validation groups.
// func (c CompoundConstraint) WhenGroups(groups ...string) CompoundConstraint {
// 	c.groups = groups
// 	return c
// }
//
// func (c CompoundConstraint) validate(
// 	scope Scope,
// 	validate ValidateByConstraintFunc,
// ) (*ViolationList, error) {
// 	if c.isIgnored || scope.IsIgnored(c.groups...) {
// 		return nil, nil
// 	}
//
// 	violations := &ViolationList{}
//
// 	for _, constraint := range c.constraints {
// 		err := violations.AppendFromError(validate(constraint, scope))
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
//
// 	return violations, nil
// }

type notFoundConstraint struct {
	key string
}

func (c notFoundConstraint) SetUp() error {
	return ConstraintNotFoundError{Key: c.key}
}

func (c notFoundConstraint) Name() string {
	return "notFoundConstraint"
}

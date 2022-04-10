package validation

// WhenArgument is used to build conditional validation. Use the When function to initiate a conditional check.
// If the condition is true, then the arguments passed through the Then function will be processed.
// Otherwise, the arguments passed through the Else function will be processed.
type WhenArgument struct {
	isTrue        bool
	options       []Option
	thenArguments []Argument
	elseArguments []Argument
}

// When function is used to initiate conditional validation.
// If the condition is true, then the arguments passed through the Then function will be processed.
// Otherwise, the arguments passed through the Else function will be processed.
func When(isTrue bool) WhenArgument {
	return WhenArgument{isTrue: isTrue}
}

// Then function is used to set a sequence of arguments to be processed if the condition is true.
func (arg WhenArgument) Then(arguments ...Argument) WhenArgument {
	arg.thenArguments = arguments
	return arg
}

// Else function is used to set a sequence of arguments to be processed if a condition is false.
func (arg WhenArgument) Else(arguments ...Argument) WhenArgument {
	arg.elseArguments = arguments
	return arg
}

// With returns a copy of WhenArgument with appended options.
func (arg WhenArgument) With(options ...Option) WhenArgument {
	arg.options = append(arg.options, options...)
	return arg
}

func (arg WhenArgument) setUp(ctx *executionContext) {
	ctx.addValidator(arg.options, func(scope Scope) (*ViolationList, error) {
		var err error
		if arg.isTrue {
			err = scope.validate(arg.thenArguments...)
		} else {
			err = scope.validate(arg.elseArguments...)
		}

		return unwrapViolationList(err)
	})
}

// SequentialArgument can be used to interrupt validation process when the first violation is raised.
type SequentialArgument struct {
	isIgnored bool
	groups    []string
	options   []Option
	arguments []Argument
}

// Sequentially function used to run validation process step-by-step.
func Sequentially(arguments ...Argument) SequentialArgument {
	return SequentialArgument{arguments: arguments}
}

// With returns a copy of SequentialArgument with appended options.
func (arg SequentialArgument) With(options ...Option) SequentialArgument {
	arg.options = append(arg.options, options...)
	return arg
}

// When enables conditional validation of this argument. If the expression evaluates to false,
// then the argument will be ignored.
func (arg SequentialArgument) When(condition bool) SequentialArgument {
	arg.isIgnored = !condition
	return arg
}

// WhenGroups enables conditional validation of the argument by using the validation groups.
func (arg SequentialArgument) WhenGroups(groups ...string) SequentialArgument {
	arg.groups = groups
	return arg
}

func (arg SequentialArgument) setUp(ctx *executionContext) {
	ctx.addValidator(arg.options, func(scope Scope) (*ViolationList, error) {
		if arg.isIgnored || scope.IsIgnored(arg.groups...) {
			return nil, nil
		}

		violations := &ViolationList{}

		for _, argument := range arg.arguments {
			err := violations.AppendFromError(scope.validate(argument))
			if err != nil {
				return nil, err
			}
			if violations.len > 0 {
				return violations, nil
			}
		}

		return violations, nil
	})
}

// AtLeastOneOfArgument can be used to set up validation process to check that the value satisfies
// at least one of the given constraints. The validation stops as soon as one constraint is satisfied.
type AtLeastOneOfArgument struct {
	isIgnored bool
	groups    []string
	options   []Option
	arguments []Argument
}

// AtLeastOneOf can be used to set up validation process to check that the value satisfies
// at least one of the given constraints. The validation stops as soon as one constraint is satisfied.
func AtLeastOneOf(arguments ...Argument) AtLeastOneOfArgument {
	return AtLeastOneOfArgument{arguments: arguments}
}

// With returns a copy of AtLeastOneOfArgument with appended options.
func (arg AtLeastOneOfArgument) With(options ...Option) AtLeastOneOfArgument {
	arg.options = append(arg.options, options...)
	return arg
}

// When enables conditional validation of this argument. If the expression evaluates to false,
// then the argument will be ignored.
func (arg AtLeastOneOfArgument) When(condition bool) AtLeastOneOfArgument {
	arg.isIgnored = !condition
	return arg
}

// WhenGroups enables conditional validation of the argument by using the validation groups.
func (arg AtLeastOneOfArgument) WhenGroups(groups ...string) AtLeastOneOfArgument {
	arg.groups = groups
	return arg
}

func (arg AtLeastOneOfArgument) setUp(ctx *executionContext) {
	ctx.addValidator(arg.options, func(scope Scope) (*ViolationList, error) {
		if arg.isIgnored || scope.IsIgnored(arg.groups...) {
			return nil, nil
		}

		violations := &ViolationList{}

		for _, argument := range arg.arguments {
			violation := scope.validate(argument)
			if violation == nil {
				return nil, nil
			}

			err := violations.AppendFromError(violation)
			if err != nil {
				return nil, err
			}
		}

		return violations, nil
	})
}

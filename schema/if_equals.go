package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func validateEquals(builtin parse.IntrinsicFunction, ctx PropertyContext) reporting.Reports {
	if errs := validateIntrinsicFunctionBasicCriteria(parse.FnEquals, builtin, ctx); errs != nil {
		return errs
	}

	value := builtin.UnderlyingMap[string(parse.FnEquals)]
	args, ok := value.([]interface{})
	if !ok || args == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Invalid type for \"Fn::Equals\" key: %T", value)}
	}

	if len(args) != 2 {
		return reporting.Reports{reporting.NewFailure(ctx, "Incorrect number of arguments (expected: 2, actual: %d)", len(args))}
	}

	reports := make(reporting.Reports, 0, 10)

	left := args[0]
	if errs := validateEqualsItem(left, PropertyContextAdd(NewPropertyContext(ctx, Schema{Type: ValueString}), "Value-1")); errs != nil {
		reports = append(reports, errs...)
	}

	right := args[1]
	if errs := validateEqualsItem(right, PropertyContextAdd(NewPropertyContext(ctx, Schema{Type: ValueString}), "Value-2")); errs != nil {
		reports = append(reports, errs...)
	}

	return reporting.Safe(reports)
}

func validateEqualsItem(value interface{}, ctx PropertyContext) reporting.Reports {
	if value == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Value is null")}
	}

	switch t := value.(type) {
	case parse.IntrinsicFunction:
		_, errs := ValidateIntrinsicFunctions(t, ctx, SupportedFunctions{
			parse.FnAnd:       true,
			parse.FnCondition: true,
			parse.FnEquals:    true,
			parse.FnFindInMap: true,
			parse.FnIf:        true,
			parse.FnNot:       true,
			parse.FnOr:        true,
			parse.FnRef:       true,
		})
		return errs
	}

	return nil
}

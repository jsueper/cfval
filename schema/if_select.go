package schema

import (
	"strconv"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-select.html
func validateSelect(builtin parse.IntrinsicFunction, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	value, found := builtin.UnderlyingMap["Fn::Select"]
	if !found || value == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Missing \"Fn::Select\" key")}
	}

	if len(builtin.UnderlyingMap) > 1 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(builtin.UnderlyingMap, "Fn::Select"))}
	}

	switch t := value.(type) {
	case []interface{}:
		return validateSelectParameters(builtin, t, ctx)
	}

	return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Invalid \"Fn::Select\" key: %s", value)}
}

func validateSelectParameters(builtin parse.IntrinsicFunction, args []interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if len(args) != 2 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Wrong number of arguments to Fn::Select (expected 2, got %d)", len(args))}
	}

	reports := make(reporting.Reports, 0, 10)

	index := args[0]
	array := args[1]

	if _, errs := validateSelectIndex(builtin, index, array, PropertyContextAdd(ctx, "Index")); errs != nil {
		reports = append(reports, errs...)
	}

	if _, errs := validateSelectArray(builtin, array, PropertyContextAdd(ctx, "Array")); errs != nil {
		reports = append(reports, errs...)
	}

	return reporting.ValidateAbort, reporting.Safe(reports)
}

func validateSelectIndex(builtin parse.IntrinsicFunction, index interface{}, array interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if index == nil {
		return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Index cannot be null")}
	}

	switch t := index.(type) {
	case float64:
		if t < 0 {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Index cannot less than zero")}
		} else if arr, ok := array.([]interface{}); ok && int(t) >= len(arr) {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Index cannot greater than array length")}
		}

		return reporting.ValidateOK, nil
	case parse.IntrinsicFunction:
		return ValidateIntrinsicFunctions(t, ctx, SupportedFunctions{
			parse.FnRef:       true,
			parse.FnFindInMap: true,
		})
	}

	return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Invalid value for index %v", index)}
}

func validateSelectArray(builtin parse.IntrinsicFunction, array interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if array == nil {
		return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Array cannot be null")}
	}

	switch t := array.(type) {
	case []interface{}:
		reports := make(reporting.Reports, 0, 10)
		for i, item := range t {
			if item == nil {
				reports = append(reports, reporting.NewFailure(PropertyContextAdd(ctx, strconv.Itoa(i)), "Array item cannot be null"))
			}
		}
		return reporting.ValidateOK, reporting.Safe(reports)
	case parse.IntrinsicFunction:
		return ValidateIntrinsicFunctions(t, ctx, SupportedFunctions{
			parse.FnRef:       true,
			parse.FnFindInMap: true,
			parse.FnGetAtt:    true,
			parse.FnGetAZs:    true,
			parse.FnIf:        true,
		})
	}

	return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Invalid value for array %s", array)}
}
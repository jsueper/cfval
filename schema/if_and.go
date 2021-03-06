package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-conditions.html#d0e97554
func validateAnd(builtin parse.IntrinsicFunction, ctx PropertyContext) reporting.Reports {
	return validateAndOr(parse.FnAnd, builtin, ctx)
}

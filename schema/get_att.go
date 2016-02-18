package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

type GetAtt struct {
	definition []interface{}
}

func NewGetAtt(definition []interface{}) GetAtt {
	return GetAtt{definition}
}

func (ga GetAtt) Validate(ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if len(ga.definition) != 2 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(fmt.Sprintf("GetAtt has incorrect number of arguments (expected: 2, actual: %d)", len(ga.definition)), ctx.Path())}
	}

	template := ctx.Template()
	if resourceID, ok := ga.definition[0].(string); ok {
		if resource, ok := template.Resources[resourceID]; ok {
			if attributeName, ok := ga.definition[1].(string); ok {
				definition := ctx.Definitions().Lookup(resource.Type)
				// TODO: BUG this line below should be attribute, ok
				if resource, ok := definition.Attributes[attributeName]; ok {
					// TODO: make this common, so GetAtt and others can use it
					targetType := resource.Type
					switch targetType.CoercibleTo(ctx.Property().Type) {
					case CoercionNever:
						return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(fmt.Sprintf("GetAtt value of %s.%s is %s but is being assigned to a %s property", resourceID, attributeName, targetType.Describe(), ctx.Property().Type.Describe()), ctx.Path())}
					case CoercionBegrudgingly:
						return reporting.ValidateAbort, reporting.Reports{reporting.NewWarning(fmt.Sprintf("GetAtt value of %s.%s is %s but is being dangerously coerced to a %s property", resourceID, attributeName, targetType.Describe(), ctx.Property().Type.Describe()), ctx.Path())}
					}

					return reporting.ValidateAbort, nil
				}
			}

			// attribute not found on resource
			return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(fmt.Sprintf("GetAtt %s.%s is not an attribute", resourceID, ga.definition[1]), ctx.Path())}
		}

		// resource not found
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(fmt.Sprintf("GetAtt '%s' is not a resource", resourceID), ctx.Path())}
	}

	// resource not a string
	return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(fmt.Sprintf("GetAtt '%s' is not a valid resource name", ga.definition[0]), ctx.Path())}
}

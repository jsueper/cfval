package schema

import (
	"fmt"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

type NestedResource struct {
	Description string
	Properties  Properties
}

func (res NestedResource) Describe() string {
	return res.Description
}

func (NestedResource) CoercibleTo(PropertyType) Coercion {
	return CoercionNever
}

// TODO: This is all a bit hairy. We shouldn't need to be creating the
// 			 TemplateNestedResource here, ideally `self` should already refer to
//			 one and value should already be a map[string]inteface{}
func (res NestedResource) Validate(property Schema, value interface{}, self SelfRepresentation, definitions ResourceDefinitions, context []string) (reporting.ValidateResult, reporting.Reports) {
	if values, ok := value.(map[string]interface{}); ok {
		tnr := parse.NewTemplateResource(self.Template())
		tnr.Properties = values
		failures, visited := res.Properties.Validate(tnr, definitions, values, context)

		// Reject any properties we weren't expecting
		for key := range res.Properties {
			if !visited[key] {
				failures = append(failures, reporting.NewFailure(fmt.Sprintf("Unknown property '%s' for nested %s", key, res.Description), append(context, key)))
			}
		}

		if len(failures) == 0 {
			return reporting.ValidateOK, nil
		}

		return reporting.ValidateOK, failures
	}

	return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Invalid type %T for nested resource %s", value, res.Description), context)}
}

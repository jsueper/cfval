package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

type NestedResource struct {
	Description string
	Properties  ValidatableProperties
}

func (res NestedResource) AwsType() string {
	return res.Description
}

func (res NestedResource) PropertyDefault(name string) (interface{}, bool) {
	return res.Properties.PropertyDefault(name)
}

func (NestedResource) IsArray() bool {
	return false
}

func (res NestedResource) Describe() string {
	return res.Description
}

func (res NestedResource) Same(to PropertyType) bool {
	if nr, ok := to.(NestedResource); ok {
		return nr.Description == res.Description
	}

	return false
}

func (NestedResource) CoercibleTo(PropertyType) Coercion {
	return CoercionNever
}

// TODO: This is all a bit hairy. We shouldn't need to be creating the
// 			 TemplateNestedResource here, ideally `self` should already refer to
//			 one and value should already be a map[string]inteface{}
func (res NestedResource) Validate(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if values, ok := value.(map[string]interface{}); ok {
		property := ctx.Property()
		tnr := parse.NewTemplateResource(property.Type.Describe(), values)

		nestedResourceContext := NewResourceContext(ctx, ResourceWithDefinition{tnr, property.Type})
		failures := res.Properties.Validate(nestedResourceContext)

		return reporting.ValidateOK, reporting.Safe(failures)
	}

	return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Invalid type %T for nested resource %s", value, res.Description)}
}

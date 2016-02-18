package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

type Resource struct {
	AwsType      string
	Attributes   map[string]Schema
	Properties   Properties
	ReturnValue  Schema
	ValidateFunc func(ResourceContext) (reporting.ValidateResult, reporting.Reports)
}

func (rd Resource) Validate(ctx ResourceContext) (reporting.ValidateResult, reporting.Reports) {
	if rd.ValidateFunc != nil {
		return rd.ValidateFunc(ctx)
	}

	failures, visited := rd.Properties.Validate(ctx)

	// Reject any properties we weren't expecting
	for _, key := range ctx.CurrentResource().Properties() {
		if !visited[key] {
			failures = append(failures, reporting.NewFailure(fmt.Sprintf("Unknown property '%s' for %s", key, rd.AwsType), ResourceContextAdd(ctx, key)))
		}
	}

	if len(failures) == 0 {
		return reporting.ValidateOK, nil
	}

	return reporting.ValidateOK, failures
}

func (rd Resource) TargetType() PropertyType {
	return rd.ReturnValue.TargetType()
}

func (r Resource) PropertyDefault(name string) interface{} {
	return r.Properties[name].Default
}

func NewUnrecognisedResource(awsType string) Resource {
	return Resource{
		ValidateFunc: func(ctx ResourceContext) (reporting.ValidateResult, reporting.Reports) {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Unrecognised resource %s", awsType), ctx)}
		},
	}
}

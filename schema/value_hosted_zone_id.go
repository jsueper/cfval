package schema

import "github.com/jagregory/cfval/reporting"

var HostedZoneID = FuncType{
	Description: "HostedZoneID",

	Fn: func(property Schema, value interface{}, self SelfRepresentation, definitions ResourceDefinitions, context []string) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(property, value, self, definitions, context); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		// TODO: HostedZoneID validation
		return reporting.ValidateOK, nil
	},
}

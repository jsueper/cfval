package iam

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html
func InstanceProfile() Resource {
	return Resource{
		AwsType: "AWS::IAM::InstanceProfile",

		Attributes: map[string]Schema{
			"Arn": Schema{
				Type: ValueString,
			},
		},

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: map[string]Schema{
			"Path": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"Roles": Schema{
				Type:     ValueString,
				Array:    true,
				Required: constraints.Always,
			},
		},
	}
}

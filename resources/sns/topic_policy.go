package sns

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html
var TopicPolicy = Resource{
	AwsType: "AWS::SNS::TopicPolicy",

	Properties: Properties{
		"PolicyDocument": Schema{
			Type:     JSON,
			Required: constraints.Always,
		},

		"Topics": Schema{
			Type:     Multiple(ValueString),
			Required: constraints.Always,
		},
	},
}

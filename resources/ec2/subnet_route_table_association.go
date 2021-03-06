package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet-route-table-assoc.html
var SubnetRouteTableAssociation = Resource{
	AwsType: "AWS::EC2::SubnetRouteTableAssociation",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"RouteTableId": Schema{
			Type:     RouteTableID,
			Required: constraints.Always,
		},

		"SubnetId": Schema{
			Type:     SubnetID,
			Required: constraints.Always,
		},
	},
}

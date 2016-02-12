package ec2

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html
func Subnet() Resource {
	return Resource{
		AwsType: "AWS::EC2::Subnet",

		ReturnValue: Schema{
			Type: SubnetID,
		},

		Properties: Properties{
			"AvailabilityZone": Schema{
				Type: AvailabilityZone,
			},

			"CidrBlock": Schema{
				Type:     CIDR,
				Required: constraints.Always,
			},

			"MapPublicIpOnLaunch": Schema{
				Type:    ValueBool,
				Default: false,
			},

			"Tags": Schema{
				Type:  common.ResourceTag,
				Array: true,
			},

			"VpcId": Schema{
				Type:     VpcID,
				Required: constraints.Always,
			},
		},
	}
}
package rds

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html
var DBCluster = Resource{
	AwsType: "AWS::RDS::DBCluster",

	Attributes: map[string]Schema{
		"Endpoint.Address": Schema{
			Type: ValueString,
		},

		"Endpoint.Port": Schema{
			Type: ValueNumber,
		},
	},

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"AvailabilityZones": Schema{
			Type:  AvailabilityZone,
			Array: true,
		},

		"BackupRetentionPeriod": Schema{
			Type: ValueNumber,
		},

		"DatabaseName": Schema{
			Type: ValueString,
		},

		"DBClusterParameterGroupName": Schema{
			Type: ValueString,
		},

		"DBSubnetGroupName": Schema{
			Type: ValueString,
		},

		"Engine": Schema{
			Type:         ValueString,
			Required:     constraints.Always,
			ValidateFunc: SingleValueValidate("aurora"),
		},

		"EngineVersion": Schema{
			Type: ValueString,
		},

		"KmsKeyId": Schema{
			Type: ValueString,
		},

		"MasterUsername": Schema{
			Type: ValueString,
			ValidateFunc: RegexpValidate(
				`^[a-zA-Z][a-zA-Z0-9]{1,15}$`,
				"Must be 1 to 16 alphanumeric characters. First character must be a letter.",
			),
			Required:  constraints.PropertyNotExists("SnapshotIdentifier"),
			Conflicts: constraints.PropertyExists("SnapshotIdentifier"),
		},

		"MasterUserPassword": Schema{
			Type: ValueString,
			ValidateFunc: RegexpValidate(
				`^[^\/"@]{8,41}$`,
				`This password can contain any printable ASCII character except "/", """, or "@". Must contain from 8 to 41 characters.`,
			),
			Required:  constraints.PropertyNotExists("SnapshotIdentifier"),
			Conflicts: constraints.PropertyExists("SnapshotIdentifier"),
		},

		"Port": Schema{
			Type:    ValueNumber,
			Default: 3306,
		},

		"PreferredBackupWindow": Schema{
			Type: ValueString,
		},

		"PreferredMaintenanceWindow": Schema{
			Type: ValueString,
		},

		"SnapshotIdentifier": Schema{
			Type: ValueString,
		},

		"Tags": Schema{
			Type:  common.ResourceTag,
			Array: true,
		},

		"VpcSecurityGroupIds": Schema{
			Type:  SecurityGroupID,
			Array: true,
		},
	},
}

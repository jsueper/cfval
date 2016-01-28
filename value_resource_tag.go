package main

var ResourceTag = Schema{
	Type: Resource{
		AwsType: "Resource Tag",
		Properties: map[string]Schema{
			"Key": Schema{
				Type:     TypeString,
				Required: true,
			},

			"Value": Schema{
				Type:     TypeString,
				Required: true,
			},
		},
	},
}

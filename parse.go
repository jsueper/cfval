package main

import "encoding/json"

type resourceCtor func() Resource

var typeCtors = map[string]resourceCtor{
	"AWS::AutoScaling::AutoScalingGroup":    autoScalingGroup,
	"AWS::AutoScaling::LaunchConfiguration": launchConfiguration,
	"AWS::EC2::EIP":                         eip,
	"AWS::EC2::Instance":                    instance,
	"AWS::EC2::InternetGateway":             internetGateway,
	"AWS::EC2::Route":                       route,
	"AWS::EC2::RouteTable":                  routeTable,
	"AWS::EC2::SecurityGroup":               securityGroup,
	"AWS::EC2::Subnet":                      subnet,
	"AWS::EC2::SubnetRouteTableAssociation": subnetRouteTableAssociation,
	"AWS::EC2::VPCGatewayAttachment":        vpcGatewayAttachment,
	"AWS::S3::Bucket":                       bucket,
}

func parseTemplateJSON(data []byte, forgiving bool) (*Template, error) {
	var temp struct {
		Parameters map[string]Parameter
		Resources  map[string]struct {
			Type       string
			Properties map[string]interface{}
		}
	}

	err := json.Unmarshal(data, &temp)

	if err != nil {
		return nil, err
	}

	template := &Template{
		Resources: make(map[string]TemplateResource),
	}
	template.Parameters = temp.Parameters

	for key, rawResource := range temp.Resources {
		if ctor, ok := typeCtors[rawResource.Type]; ok {
			template.Resources[key] = TemplateResource{
				Definition: ctor(),
				Properties: rawResource.Properties,
			}
		} else if !forgiving {
			template.Resources[key] = NewUnrecognisedResource(rawResource.Type)
		}
	}

	return template, nil
}

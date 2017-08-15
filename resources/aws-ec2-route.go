package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::Route AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html
type AWSEC2Route struct {

	// DestinationCidrBlock AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-destinationcidrblock

	DestinationCidrBlock string `json:"DestinationCidrBlock"`

	// DestinationIpv6CidrBlock AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-destinationipv6cidrblock

	DestinationIpv6CidrBlock string `json:"DestinationIpv6CidrBlock"`

	// GatewayId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-gatewayid

	GatewayId string `json:"GatewayId"`

	// InstanceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-instanceid

	InstanceId string `json:"InstanceId"`

	// NatGatewayId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-natgatewayid

	NatGatewayId string `json:"NatGatewayId"`

	// NetworkInterfaceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-networkinterfaceid

	NetworkInterfaceId string `json:"NetworkInterfaceId"`

	// RouteTableId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-routetableid

	RouteTableId string `json:"RouteTableId"`

	// VpcPeeringConnectionId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-route.html#cfn-ec2-route-vpcpeeringconnectionid

	VpcPeeringConnectionId string `json:"VpcPeeringConnectionId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Route) AWSCloudFormationType() string {
	return "AWS::EC2::Route"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2Route) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2RouteResources retrieves all AWSEC2Route items from a CloudFormation template
func GetAllAWSEC2Route(template *Template) map[string]*AWSEC2Route {

	results := map[string]*AWSEC2Route{}
	for name, resource := range template.Resources {
		result := &AWSEC2Route{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2RouteWithName retrieves all AWSEC2Route items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2Route(name string, template *Template) (*AWSEC2Route, error) {

	result := &AWSEC2Route{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2Route{}, errors.New("resource not found")

}

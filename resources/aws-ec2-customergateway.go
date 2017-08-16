package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2CustomerGateway AWS CloudFormation Resource (AWS::EC2::CustomerGateway)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html
type AWSEC2CustomerGateway struct {

	// BgpAsn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html#cfn-ec2-customergateway-bgpasn
	BgpAsn int `json:"BgpAsn"`

	// IpAddress AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html#cfn-ec2-customergateway-ipaddress
	IpAddress string `json:"IpAddress"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html#cfn-ec2-customergateway-tags
	Tags []Tag `json:"Tags"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customer-gateway.html#cfn-ec2-customergateway-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2CustomerGateway) AWSCloudFormationType() string {
	return "AWS::EC2::CustomerGateway"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2CustomerGateway) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2CustomerGatewayResources retrieves all AWSEC2CustomerGateway items from a CloudFormation template
func (t *Template) GetAllAWSEC2CustomerGatewayResources() map[string]*AWSEC2CustomerGateway {

	results := map[string]*AWSEC2CustomerGateway{}
	for name, resource := range t.Resources {
		result := &AWSEC2CustomerGateway{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2CustomerGatewayWithName retrieves all AWSEC2CustomerGateway items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2CustomerGatewayWithName(name string) (*AWSEC2CustomerGateway, error) {

	result := &AWSEC2CustomerGateway{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2CustomerGateway{}, errors.New("resource not found")

}

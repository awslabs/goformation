package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::EgressOnlyInternetGateway AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-egressonlyinternetgateway.html
type AWSEC2EgressOnlyInternetGateway struct {

	// VpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-egressonlyinternetgateway.html#cfn-ec2-egressonlyinternetgateway-vpcid
	VpcId string `json:"VpcId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2EgressOnlyInternetGateway) AWSCloudFormationType() string {
	return "AWS::EC2::EgressOnlyInternetGateway"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2EgressOnlyInternetGateway) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2EgressOnlyInternetGatewayResources retrieves all AWSEC2EgressOnlyInternetGateway items from a CloudFormation template
func GetAllAWSEC2EgressOnlyInternetGatewayResources(template *Template) map[string]*AWSEC2EgressOnlyInternetGateway {

	results := map[string]*AWSEC2EgressOnlyInternetGateway{}
	for name, resource := range template.Resources {
		result := &AWSEC2EgressOnlyInternetGateway{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2EgressOnlyInternetGatewayWithName retrieves all AWSEC2EgressOnlyInternetGateway items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSEC2EgressOnlyInternetGatewayWithName(name string, template *Template) (*AWSEC2EgressOnlyInternetGateway, error) {

	result := &AWSEC2EgressOnlyInternetGateway{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2EgressOnlyInternetGateway{}, errors.New("resource not found")

}

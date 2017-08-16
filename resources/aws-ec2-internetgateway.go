package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2InternetGateway AWS CloudFormation Resource (AWS::EC2::InternetGateway)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-internetgateway.html
type AWSEC2InternetGateway struct {

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-internetgateway.html#cfn-ec2-internetgateway-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2InternetGateway) AWSCloudFormationType() string {
	return "AWS::EC2::InternetGateway"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2InternetGateway) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2InternetGatewayResources retrieves all AWSEC2InternetGateway items from a CloudFormation template
func (t *Template) GetAllAWSEC2InternetGatewayResources() map[string]*AWSEC2InternetGateway {

	results := map[string]*AWSEC2InternetGateway{}
	for name, resource := range t.Resources {
		result := &AWSEC2InternetGateway{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2InternetGatewayWithName retrieves all AWSEC2InternetGateway items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2InternetGatewayWithName(name string) (*AWSEC2InternetGateway, error) {

	result := &AWSEC2InternetGateway{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2InternetGateway{}, errors.New("resource not found")

}

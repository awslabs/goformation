package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::NatGateway AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-natgateway.html
type AWSEC2NatGateway struct {

	// AllocationId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-natgateway.html#cfn-ec2-natgateway-allocationid
	AllocationId string `json:"AllocationId"`

	// SubnetId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-natgateway.html#cfn-ec2-natgateway-subnetid
	SubnetId string `json:"SubnetId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2NatGateway) AWSCloudFormationType() string {
	return "AWS::EC2::NatGateway"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2NatGateway) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2NatGatewayResources retrieves all AWSEC2NatGateway items from a CloudFormation template
func GetAllAWSEC2NatGateway(template *Template) map[string]*AWSEC2NatGateway {

	results := map[string]*AWSEC2NatGateway{}
	for name, resource := range template.Resources {
		result := &AWSEC2NatGateway{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2NatGatewayWithName retrieves all AWSEC2NatGateway items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2NatGateway(name string, template *Template) (*AWSEC2NatGateway, error) {

	result := &AWSEC2NatGateway{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2NatGateway{}, errors.New("resource not found")

}

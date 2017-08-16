package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2VPCEndpoint AWS CloudFormation Resource (AWS::EC2::VPCEndpoint)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html
type AWSEC2VPCEndpoint struct {

	// PolicyDocument AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-policydocument
	PolicyDocument interface{} `json:"PolicyDocument"`

	// RouteTableIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-routetableids
	RouteTableIds []string `json:"RouteTableIds"`

	// ServiceName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-servicename
	ServiceName string `json:"ServiceName"`

	// VpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-vpcid
	VpcId string `json:"VpcId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPCEndpoint) AWSCloudFormationType() string {
	return "AWS::EC2::VPCEndpoint"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPCEndpoint) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2VPCEndpointResources retrieves all AWSEC2VPCEndpoint items from a CloudFormation template
func (t *Template) GetAllAWSEC2VPCEndpointResources() map[string]*AWSEC2VPCEndpoint {

	results := map[string]*AWSEC2VPCEndpoint{}
	for name, resource := range t.Resources {
		result := &AWSEC2VPCEndpoint{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2VPCEndpointWithName retrieves all AWSEC2VPCEndpoint items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCEndpointWithName(name string) (*AWSEC2VPCEndpoint, error) {

	result := &AWSEC2VPCEndpoint{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2VPCEndpoint{}, errors.New("resource not found")

}

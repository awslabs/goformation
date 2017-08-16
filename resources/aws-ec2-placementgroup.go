package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2PlacementGroup AWS CloudFormation Resource (AWS::EC2::PlacementGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html
type AWSEC2PlacementGroup struct {

	// Strategy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html#cfn-ec2-placementgroup-strategy
	Strategy string `json:"Strategy"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2PlacementGroup) AWSCloudFormationType() string {
	return "AWS::EC2::PlacementGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2PlacementGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2PlacementGroupResources retrieves all AWSEC2PlacementGroup items from a CloudFormation template
func (t *Template) GetAllAWSEC2PlacementGroupResources() map[string]*AWSEC2PlacementGroup {

	results := map[string]*AWSEC2PlacementGroup{}
	for name, resource := range t.Resources {
		result := &AWSEC2PlacementGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2PlacementGroupWithName retrieves all AWSEC2PlacementGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2PlacementGroupWithName(name string) (*AWSEC2PlacementGroup, error) {

	result := &AWSEC2PlacementGroup{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2PlacementGroup{}, errors.New("resource not found")

}

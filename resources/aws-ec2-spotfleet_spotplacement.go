package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::SpotFleet.SpotPlacement AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-placement.html
type AWSEC2SpotFleet_SpotPlacement struct {

	// AvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-placement.html#cfn-ec2-spotfleet-spotplacement-availabilityzone

	AvailabilityZone string `json:"AvailabilityZone"`

	// GroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-placement.html#cfn-ec2-spotfleet-spotplacement-groupname

	GroupName string `json:"GroupName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_SpotPlacement) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.SpotPlacement"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SpotFleet_SpotPlacement) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2SpotFleet_SpotPlacementResources retrieves all AWSEC2SpotFleet_SpotPlacement items from a CloudFormation template
func GetAllAWSEC2SpotFleet_SpotPlacement(template *Template) map[string]*AWSEC2SpotFleet_SpotPlacement {

	results := map[string]*AWSEC2SpotFleet_SpotPlacement{}
	for name, resource := range template.Resources {
		result := &AWSEC2SpotFleet_SpotPlacement{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2SpotFleet_SpotPlacementWithName retrieves all AWSEC2SpotFleet_SpotPlacement items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2SpotFleet_SpotPlacement(name string, template *Template) (*AWSEC2SpotFleet_SpotPlacement, error) {

	result := &AWSEC2SpotFleet_SpotPlacement{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2SpotFleet_SpotPlacement{}, errors.New("resource not found")

}

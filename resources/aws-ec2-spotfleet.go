package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEC2SpotFleet AWS CloudFormation Resource (AWS::EC2::SpotFleet)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-spotfleet.html
type AWSEC2SpotFleet struct {

	// SpotFleetRequestConfigData AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-spotfleet.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata
	SpotFleetRequestConfigData AWSEC2SpotFleet_SpotFleetRequestConfigData `json:"SpotFleetRequestConfigData"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SpotFleet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2SpotFleetResources retrieves all AWSEC2SpotFleet items from a CloudFormation template
func GetAllAWSEC2SpotFleetResources(template *Template) map[string]*AWSEC2SpotFleet {

	results := map[string]*AWSEC2SpotFleet{}
	for name, resource := range template.Resources {
		result := &AWSEC2SpotFleet{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2SpotFleetWithName retrieves all AWSEC2SpotFleet items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSEC2SpotFleetWithName(name string, template *Template) (*AWSEC2SpotFleet, error) {

	result := &AWSEC2SpotFleet{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2SpotFleet{}, errors.New("resource not found")

}

package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::SpotFleet.GroupIdentifier AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-securitygroups.html
type AWSEC2SpotFleet_GroupIdentifier struct {

	// GroupId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-securitygroups.html#cfn-ec2-spotfleet-groupidentifier-groupid

	GroupId string `json:"GroupId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_GroupIdentifier) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.GroupIdentifier"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SpotFleet_GroupIdentifier) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2SpotFleet_GroupIdentifierResources retrieves all AWSEC2SpotFleet_GroupIdentifier items from a CloudFormation template
func GetAllAWSEC2SpotFleet_GroupIdentifier(template *Template) map[string]*AWSEC2SpotFleet_GroupIdentifier {

	results := map[string]*AWSEC2SpotFleet_GroupIdentifier{}
	for name, resource := range template.Resources {
		result := &AWSEC2SpotFleet_GroupIdentifier{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2SpotFleet_GroupIdentifierWithName retrieves all AWSEC2SpotFleet_GroupIdentifier items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2SpotFleet_GroupIdentifier(name string, template *Template) (*AWSEC2SpotFleet_GroupIdentifier, error) {

	result := &AWSEC2SpotFleet_GroupIdentifier{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2SpotFleet_GroupIdentifier{}, errors.New("resource not found")

}

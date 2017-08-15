package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::SpotFleet.IamInstanceProfileSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-iaminstanceprofile.html
type AWSEC2SpotFleet_IamInstanceProfileSpecification struct {

	// Arn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-iaminstanceprofile.html#cfn-ec2-spotfleet-iaminstanceprofilespecification-arn
	Arn string `json:"Arn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_IamInstanceProfileSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.IamInstanceProfileSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SpotFleet_IamInstanceProfileSpecification) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2SpotFleet_IamInstanceProfileSpecificationResources retrieves all AWSEC2SpotFleet_IamInstanceProfileSpecification items from a CloudFormation template
func GetAllAWSEC2SpotFleet_IamInstanceProfileSpecification(template *Template) map[string]*AWSEC2SpotFleet_IamInstanceProfileSpecification {

	results := map[string]*AWSEC2SpotFleet_IamInstanceProfileSpecification{}
	for name, resource := range template.Resources {
		result := &AWSEC2SpotFleet_IamInstanceProfileSpecification{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2SpotFleet_IamInstanceProfileSpecificationWithName retrieves all AWSEC2SpotFleet_IamInstanceProfileSpecification items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2SpotFleet_IamInstanceProfileSpecification(name string, template *Template) (*AWSEC2SpotFleet_IamInstanceProfileSpecification, error) {

	result := &AWSEC2SpotFleet_IamInstanceProfileSpecification{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2SpotFleet_IamInstanceProfileSpecification{}, errors.New("resource not found")

}

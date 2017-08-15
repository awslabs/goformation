package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EC2::SpotFleet.PrivateIpAddressSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces-privateipaddresses.html
type AWSEC2SpotFleet_PrivateIpAddressSpecification struct {

	// Primary AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces-privateipaddresses.html#cfn-ec2-spotfleet-privateipaddressspecification-primary
	Primary bool `json:"Primary"`

	// PrivateIpAddress AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications-networkinterfaces-privateipaddresses.html#cfn-ec2-spotfleet-privateipaddressspecification-privateipaddress
	PrivateIpAddress string `json:"PrivateIpAddress"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_PrivateIpAddressSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.PrivateIpAddressSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SpotFleet_PrivateIpAddressSpecification) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEC2SpotFleet_PrivateIpAddressSpecificationResources retrieves all AWSEC2SpotFleet_PrivateIpAddressSpecification items from a CloudFormation template
func GetAllAWSEC2SpotFleet_PrivateIpAddressSpecification(template *Template) map[string]*AWSEC2SpotFleet_PrivateIpAddressSpecification {

	results := map[string]*AWSEC2SpotFleet_PrivateIpAddressSpecification{}
	for name, resource := range template.Resources {
		result := &AWSEC2SpotFleet_PrivateIpAddressSpecification{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEC2SpotFleet_PrivateIpAddressSpecificationWithName retrieves all AWSEC2SpotFleet_PrivateIpAddressSpecification items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEC2SpotFleet_PrivateIpAddressSpecification(name string, template *Template) (*AWSEC2SpotFleet_PrivateIpAddressSpecification, error) {

	result := &AWSEC2SpotFleet_PrivateIpAddressSpecification{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEC2SpotFleet_PrivateIpAddressSpecification{}, errors.New("resource not found")

}

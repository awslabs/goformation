package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::GameLift::Fleet.IpPermission AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ec2inboundpermission.html
type AWSGameLiftFleet_IpPermission struct {

	// FromPort AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ec2inboundpermission.html#cfn-gamelift-fleet-ec2inboundpermissions-fromport
	FromPort int64 `json:"FromPort"`

	// IpRange AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ec2inboundpermission.html#cfn-gamelift-fleet-ec2inboundpermissions-iprange
	IpRange string `json:"IpRange"`

	// Protocol AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ec2inboundpermission.html#cfn-gamelift-fleet-ec2inboundpermissions-protocol
	Protocol string `json:"Protocol"`

	// ToPort AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ec2inboundpermission.html#cfn-gamelift-fleet-ec2inboundpermissions-toport
	ToPort int64 `json:"ToPort"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGameLiftFleet_IpPermission) AWSCloudFormationType() string {
	return "AWS::GameLift::Fleet.IpPermission"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSGameLiftFleet_IpPermission) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSGameLiftFleet_IpPermissionResources retrieves all AWSGameLiftFleet_IpPermission items from a CloudFormation template
func GetAllAWSGameLiftFleet_IpPermission(template *Template) map[string]*AWSGameLiftFleet_IpPermission {

	results := map[string]*AWSGameLiftFleet_IpPermission{}
	for name, resource := range template.Resources {
		result := &AWSGameLiftFleet_IpPermission{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSGameLiftFleet_IpPermissionWithName retrieves all AWSGameLiftFleet_IpPermission items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSGameLiftFleet_IpPermission(name string, template *Template) (*AWSGameLiftFleet_IpPermission, error) {

	result := &AWSGameLiftFleet_IpPermission{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSGameLiftFleet_IpPermission{}, errors.New("resource not found")

}

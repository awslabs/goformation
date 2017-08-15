package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::GameLift::Fleet AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html
type AWSGameLiftFleet struct {

	// BuildId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-buildid
	BuildId string `json:"BuildId"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-description
	Description string `json:"Description"`

	// DesiredEC2Instances AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-desiredec2instances
	DesiredEC2Instances int64 `json:"DesiredEC2Instances"`

	// EC2InboundPermissions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-ec2inboundpermissions
	EC2InboundPermissions []AWSGameLiftFleet_IpPermission `json:"EC2InboundPermissions"`

	// EC2InstanceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-ec2instancetype
	EC2InstanceType string `json:"EC2InstanceType"`

	// LogPaths AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-logpaths
	LogPaths []string `json:"LogPaths"`

	// MaxSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-maxsize
	MaxSize int64 `json:"MaxSize"`

	// MinSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-minsize
	MinSize int64 `json:"MinSize"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-name
	Name string `json:"Name"`

	// ServerLaunchParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-serverlaunchparameters
	ServerLaunchParameters string `json:"ServerLaunchParameters"`

	// ServerLaunchPath AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-serverlaunchpath
	ServerLaunchPath string `json:"ServerLaunchPath"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGameLiftFleet) AWSCloudFormationType() string {
	return "AWS::GameLift::Fleet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSGameLiftFleet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSGameLiftFleetResources retrieves all AWSGameLiftFleet items from a CloudFormation template
func GetAllAWSGameLiftFleet(template *Template) map[string]*AWSGameLiftFleet {

	results := map[string]*AWSGameLiftFleet{}
	for name, resource := range template.Resources {
		result := &AWSGameLiftFleet{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSGameLiftFleetWithName retrieves all AWSGameLiftFleet items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSGameLiftFleet(name string, template *Template) (*AWSGameLiftFleet, error) {

	result := &AWSGameLiftFleet{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSGameLiftFleet{}, errors.New("resource not found")

}

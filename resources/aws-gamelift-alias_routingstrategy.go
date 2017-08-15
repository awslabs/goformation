package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::GameLift::Alias.RoutingStrategy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html
type AWSGameLiftAlias_RoutingStrategy struct {

	// FleetId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-fleetid
	FleetId string `json:"FleetId"`

	// Message AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-message
	Message string `json:"Message"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGameLiftAlias_RoutingStrategy) AWSCloudFormationType() string {
	return "AWS::GameLift::Alias.RoutingStrategy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSGameLiftAlias_RoutingStrategy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::GameLift::Alias AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html
type AWSGameLiftAlias struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-description

	Description string `json:"Description"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-name

	Name string `json:"Name"`

	// RoutingStrategy AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-routingstrategy

	RoutingStrategy AWSGameLiftAlias_RoutingStrategy `json:"RoutingStrategy"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGameLiftAlias) AWSCloudFormationType() string {
	return "AWS::GameLift::Alias"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSGameLiftAlias) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSGameLiftAliasResources retrieves all AWSGameLiftAlias items from a CloudFormation template
func GetAllAWSGameLiftAlias(template *Template) map[string]*AWSGameLiftAlias {

	results := map[string]*AWSGameLiftAlias{}
	for name, resource := range template.Resources {
		result := &AWSGameLiftAlias{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSGameLiftAliasWithName retrieves all AWSGameLiftAlias items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSGameLiftAlias(name string, template *Template) (*AWSGameLiftAlias, error) {

	result := &AWSGameLiftAlias{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSGameLiftAlias{}, errors.New("resource not found")

}

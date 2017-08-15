package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSGameLiftBuild AWS CloudFormation Resource (AWS::GameLift::Build)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html
type AWSGameLiftBuild struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-name
	Name string `json:"Name"`

	// StorageLocation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-storagelocation
	StorageLocation AWSGameLiftBuild_S3Location `json:"StorageLocation"`

	// Version AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-version
	Version string `json:"Version"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGameLiftBuild) AWSCloudFormationType() string {
	return "AWS::GameLift::Build"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSGameLiftBuild) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSGameLiftBuildResources retrieves all AWSGameLiftBuild items from a CloudFormation template
func GetAllAWSGameLiftBuildResources(template *Template) map[string]*AWSGameLiftBuild {

	results := map[string]*AWSGameLiftBuild{}
	for name, resource := range template.Resources {
		result := &AWSGameLiftBuild{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSGameLiftBuildWithName retrieves all AWSGameLiftBuild items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSGameLiftBuildWithName(name string, template *Template) (*AWSGameLiftBuild, error) {

	result := &AWSGameLiftBuild{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSGameLiftBuild{}, errors.New("resource not found")

}

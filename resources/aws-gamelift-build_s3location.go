package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::GameLift::Build.S3Location AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html
type AWSGameLiftBuild_S3Location struct {

	// Bucket AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storage-bucket

	Bucket string `json:"Bucket"`

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storage-key

	Key string `json:"Key"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storage-rolearn

	RoleArn string `json:"RoleArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGameLiftBuild_S3Location) AWSCloudFormationType() string {
	return "AWS::GameLift::Build.S3Location"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSGameLiftBuild_S3Location) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSGameLiftBuild_S3LocationResources retrieves all AWSGameLiftBuild_S3Location items from a CloudFormation template
func GetAllAWSGameLiftBuild_S3Location(template *Template) map[string]*AWSGameLiftBuild_S3Location {

	results := map[string]*AWSGameLiftBuild_S3Location{}
	for name, resource := range template.Resources {
		result := &AWSGameLiftBuild_S3Location{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSGameLiftBuild_S3LocationWithName retrieves all AWSGameLiftBuild_S3Location items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSGameLiftBuild_S3Location(name string, template *Template) (*AWSGameLiftBuild_S3Location, error) {

	result := &AWSGameLiftBuild_S3Location{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSGameLiftBuild_S3Location{}, errors.New("resource not found")

}

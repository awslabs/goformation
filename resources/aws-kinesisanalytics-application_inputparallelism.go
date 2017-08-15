package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::Application.InputParallelism AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputparallelism.html
type AWSKinesisAnalyticsApplication_InputParallelism struct {

	// Count AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputparallelism.html#cfn-kinesisanalytics-application-inputparallelism-count

	Count int64 `json:"Count"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_InputParallelism) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.InputParallelism"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplication_InputParallelism) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplication_InputParallelismResources retrieves all AWSKinesisAnalyticsApplication_InputParallelism items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplication_InputParallelism(template *Template) map[string]*AWSKinesisAnalyticsApplication_InputParallelism {

	results := map[string]*AWSKinesisAnalyticsApplication_InputParallelism{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplication_InputParallelism{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplication_InputParallelismWithName retrieves all AWSKinesisAnalyticsApplication_InputParallelism items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplication_InputParallelism(name string, template *Template) (*AWSKinesisAnalyticsApplication_InputParallelism, error) {

	result := &AWSKinesisAnalyticsApplication_InputParallelism{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplication_InputParallelism{}, errors.New("resource not found")

}

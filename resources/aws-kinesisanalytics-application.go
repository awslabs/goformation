package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::Application AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-application.html
type AWSKinesisAnalyticsApplication struct {

	// ApplicationCode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-application.html#cfn-kinesisanalytics-application-applicationcode
	ApplicationCode string `json:"ApplicationCode"`

	// ApplicationDescription AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-application.html#cfn-kinesisanalytics-application-applicationdescription
	ApplicationDescription string `json:"ApplicationDescription"`

	// ApplicationName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-application.html#cfn-kinesisanalytics-application-applicationname
	ApplicationName string `json:"ApplicationName"`

	// Inputs AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-application.html#cfn-kinesisanalytics-application-inputs
	Inputs []AWSKinesisAnalyticsApplication_Input `json:"Inputs"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplication) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplicationResources retrieves all AWSKinesisAnalyticsApplication items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplication(template *Template) map[string]*AWSKinesisAnalyticsApplication {

	results := map[string]*AWSKinesisAnalyticsApplication{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplication{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplicationWithName retrieves all AWSKinesisAnalyticsApplication items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplication(name string, template *Template) (*AWSKinesisAnalyticsApplication, error) {

	result := &AWSKinesisAnalyticsApplication{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplication{}, errors.New("resource not found")

}

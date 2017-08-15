package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticBeanstalk::ApplicationVersion.SourceBundle AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-sourcebundle.html
type AWSElasticBeanstalkApplicationVersion_SourceBundle struct {

	// S3Bucket AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-sourcebundle.html#cfn-beanstalk-sourcebundle-s3bucket
	S3Bucket string `json:"S3Bucket"`

	// S3Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-sourcebundle.html#cfn-beanstalk-sourcebundle-s3key
	S3Key string `json:"S3Key"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkApplicationVersion_SourceBundle) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::ApplicationVersion.SourceBundle"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticBeanstalkApplicationVersion_SourceBundle) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticBeanstalkApplicationVersion_SourceBundleResources retrieves all AWSElasticBeanstalkApplicationVersion_SourceBundle items from a CloudFormation template
func GetAllAWSElasticBeanstalkApplicationVersion_SourceBundle(template *Template) map[string]*AWSElasticBeanstalkApplicationVersion_SourceBundle {

	results := map[string]*AWSElasticBeanstalkApplicationVersion_SourceBundle{}
	for name, resource := range template.Resources {
		result := &AWSElasticBeanstalkApplicationVersion_SourceBundle{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticBeanstalkApplicationVersion_SourceBundleWithName retrieves all AWSElasticBeanstalkApplicationVersion_SourceBundle items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticBeanstalkApplicationVersion_SourceBundle(name string, template *Template) (*AWSElasticBeanstalkApplicationVersion_SourceBundle, error) {

	result := &AWSElasticBeanstalkApplicationVersion_SourceBundle{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticBeanstalkApplicationVersion_SourceBundle{}, errors.New("resource not found")

}

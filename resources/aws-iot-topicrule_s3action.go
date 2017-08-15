package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::TopicRule.S3Action AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-s3.html
type AWSIoTTopicRule_S3Action struct {

	// BucketName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-s3.html#cfn-iot-s3-bucketname
	BucketName string `json:"BucketName"`

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-s3.html#cfn-iot-s3-key
	Key string `json:"Key"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-s3.html#cfn-iot-s3-rolearn
	RoleArn string `json:"RoleArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_S3Action) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.S3Action"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule_S3Action) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTTopicRule_S3ActionResources retrieves all AWSIoTTopicRule_S3Action items from a CloudFormation template
func GetAllAWSIoTTopicRule_S3Action(template *Template) map[string]*AWSIoTTopicRule_S3Action {

	results := map[string]*AWSIoTTopicRule_S3Action{}
	for name, resource := range template.Resources {
		result := &AWSIoTTopicRule_S3Action{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTTopicRule_S3ActionWithName retrieves all AWSIoTTopicRule_S3Action items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTTopicRule_S3Action(name string, template *Template) (*AWSIoTTopicRule_S3Action, error) {

	result := &AWSIoTTopicRule_S3Action{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTTopicRule_S3Action{}, errors.New("resource not found")

}

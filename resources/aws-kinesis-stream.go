package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSKinesisStream AWS CloudFormation Resource (AWS::Kinesis::Stream)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html
type AWSKinesisStream struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html#cfn-kinesis-stream-name
	Name string `json:"Name"`

	// RetentionPeriodHours AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html#cfn-kinesis-stream-retentionperiodhours
	RetentionPeriodHours int `json:"RetentionPeriodHours"`

	// ShardCount AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html#cfn-kinesis-stream-shardcount
	ShardCount int `json:"ShardCount"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html#cfn-kinesis-stream-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisStream) AWSCloudFormationType() string {
	return "AWS::Kinesis::Stream"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisStream) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisStreamResources retrieves all AWSKinesisStream items from a CloudFormation template
func (t *Template) GetAllAWSKinesisStreamResources() map[string]*AWSKinesisStream {

	results := map[string]*AWSKinesisStream{}
	for name, resource := range t.Resources {
		result := &AWSKinesisStream{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisStreamWithName retrieves all AWSKinesisStream items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisStreamWithName(name string) (*AWSKinesisStream, error) {

	result := &AWSKinesisStream{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisStream{}, errors.New("resource not found")

}

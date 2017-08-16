package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSCloudTrailTrail AWS CloudFormation Resource (AWS::CloudTrail::Trail)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html
type AWSCloudTrailTrail struct {

	// CloudWatchLogsLogGroupArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-cloudwatchlogsloggroupaem
	CloudWatchLogsLogGroupArn string `json:"CloudWatchLogsLogGroupArn"`

	// CloudWatchLogsRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-cloudwatchlogsrolearn
	CloudWatchLogsRoleArn string `json:"CloudWatchLogsRoleArn"`

	// EnableLogFileValidation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-enablelogfilevalidation
	EnableLogFileValidation bool `json:"EnableLogFileValidation"`

	// IncludeGlobalServiceEvents AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-includeglobalserviceevents
	IncludeGlobalServiceEvents bool `json:"IncludeGlobalServiceEvents"`

	// IsLogging AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-islogging
	IsLogging bool `json:"IsLogging"`

	// IsMultiRegionTrail AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-ismultiregiontrail
	IsMultiRegionTrail bool `json:"IsMultiRegionTrail"`

	// KMSKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-kmskeyid
	KMSKeyId string `json:"KMSKeyId"`

	// S3BucketName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-s3bucketname
	S3BucketName string `json:"S3BucketName"`

	// S3KeyPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-s3keyprefix
	S3KeyPrefix string `json:"S3KeyPrefix"`

	// SnsTopicName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-snstopicname
	SnsTopicName string `json:"SnsTopicName"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#cfn-cloudtrail-trail-tags
	Tags []Tag `json:"Tags"`

	// TrailName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail.html#aws-cloudtrail-trail-trailname
	TrailName string `json:"TrailName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudTrailTrail) AWSCloudFormationType() string {
	return "AWS::CloudTrail::Trail"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudTrailTrail) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudTrailTrailResources retrieves all AWSCloudTrailTrail items from a CloudFormation template
func (t *Template) GetAllAWSCloudTrailTrailResources() map[string]*AWSCloudTrailTrail {

	results := map[string]*AWSCloudTrailTrail{}
	for name, resource := range t.Resources {
		result := &AWSCloudTrailTrail{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudTrailTrailWithName retrieves all AWSCloudTrailTrail items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudTrailTrailWithName(name string) (*AWSCloudTrailTrail, error) {

	result := &AWSCloudTrailTrail{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudTrailTrail{}, errors.New("resource not found")

}

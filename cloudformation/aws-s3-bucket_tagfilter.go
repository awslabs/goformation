package cloudformation

import (
	"encoding/json"
)

// AWSS3Bucket_TagFilter AWS CloudFormation Resource (AWS::S3::Bucket.TagFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-tagfilter.html
type AWSS3Bucket_TagFilter struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-tagfilter.html#cfn-s3-bucket-tagfilter-key
	Key *Value `json:"Key,omitempty"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-tagfilter.html#cfn-s3-bucket-tagfilter-value
	Value *Value `json:"Value,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_TagFilter) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.TagFilter"
}

func (r *AWSS3Bucket_TagFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}

package cloudformation

import (
	"encoding/json"
)

// AWSS3Bucket_ReplicationConfiguration AWS CloudFormation Resource (AWS::S3::Bucket.ReplicationConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration.html
type AWSS3Bucket_ReplicationConfiguration struct {

	// Role AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration.html#cfn-s3-bucket-replicationconfiguration-role
	Role *Value `json:"Role,omitempty"`

	// Rules AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration.html#cfn-s3-bucket-replicationconfiguration-rules
	Rules []AWSS3Bucket_ReplicationRule `json:"Rules,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_ReplicationConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.ReplicationConfiguration"
}

func (r *AWSS3Bucket_ReplicationConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}

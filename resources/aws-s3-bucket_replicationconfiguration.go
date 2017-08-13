package resources

// AWS::S3::Bucket.ReplicationConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration.html
type AWSS3BucketReplicationConfiguration struct {
    
    // Role AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration.html#cfn-s3-bucket-replicationconfiguration-role
    Role string `json:"Role"`
    
    // Rules AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration.html#cfn-s3-bucket-replicationconfiguration-rules
    Rules []AWSS3BucketReplicationConfigurationReplicationRule `json:"Rules"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3BucketReplicationConfiguration) AWSCloudFormationType() string {
    return "AWS::S3::Bucket.ReplicationConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3BucketReplicationConfiguration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}
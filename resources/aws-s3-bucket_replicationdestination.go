package resources

// AWS::S3::Bucket.ReplicationDestination AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html
type AWSS3BucketReplicationDestination struct {
    
    // Bucket AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html#cfn-s3-bucket-replicationconfiguration-rules-destination-bucket
    Bucket string `json:"Bucket"`
    
    // StorageClass AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationconfiguration-rules-destination.html#cfn-s3-bucket-replicationconfiguration-rules-destination-storageclass
    StorageClass string `json:"StorageClass"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3BucketReplicationDestination) AWSCloudFormationType() string {
    return "AWS::S3::Bucket.ReplicationDestination"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3BucketReplicationDestination) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}
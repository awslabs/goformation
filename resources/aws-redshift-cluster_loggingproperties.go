package resources

// AWS::Redshift::Cluster.LoggingProperties AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html
type AWSRedshiftClusterLoggingProperties struct {
    
    // BucketName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-bucketname
    BucketName string `json:"BucketName"`
    
    // S3KeyPrefix AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-s3keyprefix
    S3KeyPrefix string `json:"S3KeyPrefix"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRedshiftClusterLoggingProperties) AWSCloudFormationType() string {
    return "AWS::Redshift::Cluster.LoggingProperties"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRedshiftClusterLoggingProperties) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}
package resources

// AWS::EFS::FileSystem.ElasticFileSystemTag AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemtags.html
type AWSEFSFileSystemElasticFileSystemTag struct {
    
    // Key AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemtags.html#cfn-efs-filesystem-filesystemtags-key
    Key string `json:"Key"`
    
    // Value AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemtags.html#cfn-efs-filesystem-filesystemtags-value
    Value string `json:"Value"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEFSFileSystemElasticFileSystemTag) AWSCloudFormationType() string {
    return "AWS::EFS::FileSystem.ElasticFileSystemTag"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEFSFileSystemElasticFileSystemTag) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}
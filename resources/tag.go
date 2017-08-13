package resources

// Tag AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationinstance-tag.html
type Tag struct {
    
    // Key AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationinstance-tag.html#cfn-dms-replicationinstance-tag-key
    Key string `json:"Key"`
    
    // Value AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationinstance-tag.html#cfn-dms-replicationinstance-tag-value
    Value string `json:"Value"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Tag) AWSCloudFormationType() string {
    return "Tag"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *Tag) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}
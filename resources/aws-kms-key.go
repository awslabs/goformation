package resources

// AWS::KMS::Key AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html
type AWSKMSKey struct {
    
    // Description AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-description
    Description string `json:"Description"`
    
    // EnableKeyRotation AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enablekeyrotation
    EnableKeyRotation bool `json:"EnableKeyRotation"`
    
    // Enabled AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enabled
    Enabled bool `json:"Enabled"`
    
    // KeyPolicy AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keypolicy
    KeyPolicy string `json:"KeyPolicy"`
    
    // KeyUsage AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyusage
    KeyUsage string `json:"KeyUsage"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKMSKey) AWSCloudFormationType() string {
    return "AWS::KMS::Key"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKMSKey) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}
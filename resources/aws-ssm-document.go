package resources

// AWS::SSM::Document AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-document.html
type AWSSSMDocument struct {
    
    // Content AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-document.html#cfn-ssm-document-content
    Content string `json:"Content"`
    
    // DocumentType AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-document.html#cfn-ssm-document-documenttype
    DocumentType string `json:"DocumentType"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMDocument) AWSCloudFormationType() string {
    return "AWS::SSM::Document"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSSMDocument) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}
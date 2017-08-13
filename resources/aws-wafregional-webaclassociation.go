package resources

// AWS::WAFRegional::WebACLAssociation AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html
type AWSWAFRegionalWebACLAssociation struct {
    
    // ResourceArn AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html#cfn-wafregional-webaclassociation-resourcearn
    ResourceArn string `json:"ResourceArn"`
    
    // WebACLId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webaclassociation.html#cfn-wafregional-webaclassociation-webaclid
    WebACLId string `json:"WebACLId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalWebACLAssociation) AWSCloudFormationType() string {
    return "AWS::WAFRegional::WebACLAssociation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalWebACLAssociation) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}
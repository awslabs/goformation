package resources

// AWS::EMR::SecurityConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html
type AWSEMRSecurityConfiguration struct {
    
    // Name AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html#cfn-emr-securityconfiguration-name
    Name string `json:"Name"`
    
    // SecurityConfiguration AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html#cfn-emr-securityconfiguration-securityconfiguration
    SecurityConfiguration string `json:"SecurityConfiguration"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRSecurityConfiguration) AWSCloudFormationType() string {
    return "AWS::EMR::SecurityConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRSecurityConfiguration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}